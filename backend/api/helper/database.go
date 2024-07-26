package helper

import(
	"database/sql"
)

func NewDB() *sql.DB{
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/scmt")
	PanicIfError(err)

	return db
}