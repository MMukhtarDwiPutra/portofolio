package helper

import(
	"database/sql"
	"time"
)

func NewDB() *sql.DB{
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/scmt")
	PanicIfError(err)

	db.SetMaxOpenConns(5000) // Adjust based on your server capacity
    db.SetMaxIdleConns(100)
    db.SetConnMaxLifetime(time.Second * 1)

	return db
}