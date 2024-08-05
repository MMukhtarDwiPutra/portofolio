package domain

type User struct{
	ID int `bson:"id" json:"id"`
	Username string `bson:"username" json:"username"`
	Password string `bson:"password" json:"password"`
	Fullname string `bson:"fullname" json:"fullname"`
	Role string `bson:"role" json:"role"`
	Asal string `bson:"asal" json:"asal"`
	NomorHP string `bson:"nomor_hp" json:"nomor_hp"`
	JenisAkun string `bson:"jenis_akun" json:"jenis_akun"`
}