package config

import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type config struct {
	slave  *gorm.DB
	master *gorm.DB
	solr   *gorm.DB
}

type Config interface {
	Slave() *gorm.DB
	Master() *gorm.DB
}

func NewConfig() Config {
	config := config{}
	slave, master := InitDB()
	InitSolr()

	config.master = master
	config.slave = slave

	return &config
}

func (c *config) Slave() *gorm.DB {
	return c.slave
}

func (c *config) Master() *gorm.DB {
	return c.master
}

func InitDB() (*gorm.DB, *gorm.DB) {
	USER := os.Getenv("USER_DB")
	PASS := os.Getenv("PASS_DB")
	HOST := os.Getenv("HOST_DB")
	PORT := os.Getenv("PORT_DB")
	DBNAME := os.Getenv("DBNAME")

	dsnSlave := USER + ":" + PASS + "@tcp(" + HOST + ":" + PORT + ")/" + DBNAME + "?charset=utf8&parseTime=true&loc=Local"

	slaveConn, err := sql.Open("mysql", dsnSlave)
	if err != nil {
		panic("Cannot connect database")
	}

	if err := slaveConn.Ping(); err != nil {
		panic("Cannot connect database")
	}

	dbSlave, err := gorm.Open(mysql.New(mysql.Config{
		Conn: slaveConn,
	}), &gorm.Config{SkipDefaultTransaction: true})
	if err != nil {
		panic("Cannot connect database")
	}

	dsnMaster := USER + ":" + PASS + "@tcp(" + HOST + ":" + PORT + ")/" + DBNAME + "?charset=utf8&parseTime=true&loc=Local"
	masterConn, err := sql.Open("mysql", dsnMaster)
	if err != nil {
		panic("Cannot connect database")
	}

	if err := masterConn.Ping(); err != nil {
		panic("Cannot connect database")
	}

	dbMaster, err := gorm.Open(mysql.New(mysql.Config{
		Conn: masterConn,
	}), &gorm.Config{SkipDefaultTransaction: true})
	if err != nil {
		panic("Cannot connect database")
	}

	return dbSlave, dbMaster
}
