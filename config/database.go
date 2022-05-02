package config

import (
	"fmt"

	solr "github.com/rtt/Go-Solr"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	USER := "root"
	PASS := "050617Islam"
	HOST := "localhost"
	PORT := "3306"
	DBNAME := "tes-go"

	dsn := USER + ":" + PASS + "@tcp(" + HOST + ":" + PORT + ")/" + DBNAME + "?charset=utf8&parseTime=true&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{SkipDefaultTransaction: true})
	if err != nil {
		panic("Cannot connect database")
	}
	// DB.AutoMigrate()
}

func ConnectSolr() {

	// init a connection
	_, err := solr.Init("172.17.62.112", 8983, "collection-name")

	if err != nil {
		fmt.Println(err)
		return
	}

	// define a solr query string
	// q := "q=*:*"

	// // perform a query
	// res, err := s.SelectRaw(q)

	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// // grab results for ease of use later on
	// results := res.Results

	// // print a summary and loop over results, priting the "title" and "latlng" fields
	// fmt.Println(
	// 	fmt.Sprintf("Query: %#v\nHits: %d\nNum Results: %d\nQtime: %d\nStatus: %d\n\nResults\n-------\n",
	// 		q,
	// 		results.NumFound,
	// 		results.Len(),
	// 		res.QTime,
	// 		res.Status))

	// for i := 0; i < results.Len(); i++ {
	// 	fmt.Println("Some field:", results.Get(i).Field("id"))
	// 	fmt.Println("Some other field:", results.Get(i).Field("title"))

	// 	fmt.Println("")
	// }
}
