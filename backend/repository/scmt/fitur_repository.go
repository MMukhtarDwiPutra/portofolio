package repository

import(
	"database/sql"
	// "portofolio.com/domain/scmt"
	"portofolio.com/api/helper"
)

type FiturRepository interface{
	GetFitur(nama_fitur string) string
	SetFitur(status string, nama_fitur string)
}

type fiturRepository struct{
	db *sql.DB
}

func NewFiturRepository(db *sql.DB) *fiturRepository{
	return &fiturRepository{db}
}

func (r *fiturRepository) GetFitur(nama_fitur string) string{
	var status string

	rows, err := r.db.Query("SELECT status FROM fitur WHERE name = ?",nama_fitur)
	helper.PanicIfError(err)

	for rows.Next(){
		rows.Scan(&status)
	}

	return status
}

func (r *fiturRepository) SetFitur(status string, nama_fitur string){
	set, err := r.db.Query("UPDATE fitur SET status = ? WHERE name = ? ", status, nama_fitur)
	helper.PanicIfError(err)

	defer set.Close()
}