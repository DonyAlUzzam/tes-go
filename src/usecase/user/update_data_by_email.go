package users

type UpdateDataByEmailRequest struct {
	Email    string `json:"email" query:"email" form:"email"`
	Nama     string `json:"nama" form:"nama"`
	NoHp     string `json:"no_hp" form:"no_hp"`
	Alamat   string `json:"alamat" form:"alamat"`
	Password string `json:"password" form:"password"`
	Ktp      string `json:"ktp" form:"ktp"`
}
