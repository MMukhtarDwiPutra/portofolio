package repository

import(
	"portofolio.com/api/helper"
	"database/sql"
	"portofolio.com/domain/scmt"
	// "fmt"
)

type UserRepository interface{
	Register(user domain.User)
	GetUserByUsername(username string) domain.User
	GetUserById(id int) domain.User
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

func (r *userRepository) GetUserByUsername(username string) domain.User{
	rows, err := r.db.Query("SELECT * FROM user WHERE username = ?", username);
	helper.PanicIfError(err)

	var user domain.User
	if rows.Next(){
		rows.Scan(&user.ID, &user.Username, &user.Password, &user.Fullname, &user.Role, &user.Asal, &user.NomorHP, &user.JenisAkun)
	}

	return user
}

func (r *userRepository) GetUserById(id int) domain.User{
	rows, err := r.db.Query("SELECT * FROM user WHERE id = ?", id);
	helper.PanicIfError(err)

	var user domain.User
	if rows.Next(){
		rows.Scan(&user.ID, &user.Username, &user.Password, &user.Fullname, &user.Role, &user.Asal, &user.NomorHP, &user.JenisAkun)
	}

	return user
}