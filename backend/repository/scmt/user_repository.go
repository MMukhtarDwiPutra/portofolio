package repository

import(
	"portofolio.com/api/helper"
	"database/sql"
	"portofolio.com/domain/scmt"
	"fmt"
)

type UserRepository interface{
	Register(user domain.User)
	GetUser(username string) domain.User
}

type userRepository struct{
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *userRepository{
	return &userRepository{db}
}

func (r *userRepository) Register(user domain.User){
	_, err := r.db.Query("INSERT INTO `user`(`username`, `password`, `fullname`, `role`, `asal`, `nomor_hp`, `jenis_akun`) VALUES (?, ?, ?, ?, ?, ?, ?)", user.Username, user.Password, user.Fullname, user.Role, user.Asal, user.NomorHP, user.JenisAkun);
	helper.PanicIfError(err)
}

func (r *userRepository) GetUser(username string) domain.User{
	fmt.Println("SELECT * FROM user WHERE username = ?"+username)
	rows, err := r.db.Query("SELECT * FROM user WHERE username = ?", username);
	helper.PanicIfError(err)

	var user domain.User
	if rows.Next(){
		rows.Scan(&user.ID, &user.Username, &user.Password, &user.Fullname, &user.Role, &user.Asal, &user.NomorHP, &user.JenisAkun)
	}

	return user
}