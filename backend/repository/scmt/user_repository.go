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
	GetDataUserById(id int) domain.User
	ChangeDataUser(username string, id int)
	ChangePassword(password string, id int)
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

func (r *userRepository) GetDataUserById(id int) domain.User{
	rows, err := r.db.Query("SELECT `id`, `username`, `fullname`, `role`, `asal`, `nomor_hp`, `jenis_akun` FROM user WHERE id = ?", id);
	helper.PanicIfError(err)

	var user domain.User
	if rows.Next(){
		rows.Scan(&user.ID, &user.Username, &user.Fullname, &user.Role, &user.Asal, &user.NomorHP, &user.JenisAkun)
	}

	return user
}

func (r *userRepository) ChangeDataUser(fullname string, id int){
	_, err := r.db.Query("UPDATE `user` SET fullname = ? WHERE id = ?",fullname, id);
	helper.PanicIfError(err)
}

func (r *userRepository) ChangePassword(password string, id int){
	_, err := r.db.Query("UPDATE `user` SET password = ? WHERE id = ?",password, id);
	helper.PanicIfError(err)
}