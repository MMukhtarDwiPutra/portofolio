package repository

import (
	"database/sql"
	"log"
	"portofolio.com/domain/scmt"
	"portofolio.com/api/helper"
)

type DataTmpRepository interface{
	GetAllDataTmp() []domain.DataTmp
	InsertData(data domain.DataTmp)
	GetTableLastUpdate() (string, string)
	DeleteAllData()
	GetLastDataTmp() domain.DataTmp
	CountRetailPerWitel(merk string) []domain.CountResponse
	CountPremiumPerWitel(merk string) []domain.CountResponse
	CountSTBPerWitel(merk string) []domain.CountResponse
	CountAPPerWitel(merk string) []domain.CountResponse
	GetWitelsFromDataByMerk(merk string) []string
}

type dataTmpRepository struct{
	db *sql.DB
}

func NewDataTmpSCMTRepository(db *sql.DB) *dataTmpRepository{
	return &dataTmpRepository{db}
}

func (r *dataTmpRepository) GetAllDataTmp() []domain.DataTmp{
	var dataTmps []domain.DataTmp
	rows, err := r.db.Query("SELECT `id`, `region`, `lokasi_wh`, `status`, `jumlah`, `deskripsi` FROM `data_tmp`")
	helper.PanicIfError(err)

	for rows.Next(){
		var dataTmp domain.DataTmp
		rows.Scan(&dataTmp.ID, &dataTmp.Region, &dataTmp.LokasiWH, &dataTmp.Status, &dataTmp.Jumlah, &dataTmp.Deskripsi)

		dataTmps = append(dataTmps, dataTmp)
	}

	return dataTmps
}

func (r *dataTmpRepository) InsertData(data domain.DataTmp){
	insert, err := r.db.Query(`INSERT INTO data_tmp (region, lokasi_wh, status, jumlah, deskripsi) VALUES (?, ?, ?, ?, ?) `, data.Region, data.LokasiWH, data.Status, data.Jumlah, data.Deskripsi)
	if err != nil{
		log.Fatal(err)
	}
	defer insert.Close()
}

func (r *dataTmpRepository) GetTableLastUpdate() (string, string){
	var waktu_update string
	var waktu_dibuat string

	rows, err := r.db.Query("SELECT UPDATE_TIME as waktu_update, CREATE_TIME as waktu_dibuat FROM information_schema.tables WHERE TABLE_SCHEMA = 'scmt' AND TABLE_NAME = 'data_tmp'")
	helper.PanicIfError(err)

	for rows.Next(){
		rows.Scan(&waktu_update, &waktu_dibuat)
	}

	return waktu_update, waktu_dibuat
}

func (r *dataTmpRepository) DeleteAllData(){
	r.db.Query("DELETE FROM data_tmp")
}

func (r *dataTmpRepository) GetLastDataTmp() domain.DataTmp{
	var dataTmp domain.DataTmp

	rows, err := r.db.Query("SELECT * FROM data_tmp ORDER BY id DESC LIMIT 1")
	helper.PanicIfError(err)

	for rows.Next(){
		rows.Scan(&dataTmp.ID, &dataTmp.Region, &dataTmp.LokasiWH, &dataTmp.Status, &dataTmp.Jumlah, &dataTmp.Deskripsi)
	}

	return dataTmp
}

func (r *dataTmpRepository) CountRetailPerWitel(merk string) []domain.CountResponse{
	var countRetails []domain.CountResponse
	var rows *sql.Rows
	var err error

	if(merk == "zte"){
		rows, err = r.db.Query("SELECT lokasi_wh, sum(jumlah) as stock FROM data_tmp WHERE (deskripsi LIKE '%ONT_ZTE_F670L%') AND (deskripsi like '%ZTE%') AND (status LIKE '%AVAILABLE%' OR status LIKE '%INTECHNICIAN%') GROUP BY lokasi_wh")
	}else if(merk == "nokia"){
		rows, err = r.db.Query("SELECT lokasi_wh, sum(jumlah) as stock FROM `data_tmp` WHERE (deskripsi like '%G240WL%' OR deskripsi like '%2425G%') AND (status LIKE '%AVAILABLE%' OR status LIKE '%INTECHNICIAN%') GROUP BY lokasi_wh")
	}else if(merk == "fiberhome"){
		rows, err = r.db.Query("SELECT lokasi_wh, sum(jumlah) as stock FROM `data_tmp` WHERE (deskripsi LIKE '%HG6145D2%' or deskripsi LIKE '%HG6145F%') AND (deskripsi like '%Fiberhome%') AND (status LIKE '%AVAILABLE%' OR status LIKE '%INTECHNICIAN%') GROUP BY lokasi_wh")
	}else if(merk == "huawei"){
		rows, err = r.db.Query("SELECT lokasi_wh, sum(jumlah) as stock FROM `data_tmp` WHERE (deskripsi LIKE '%HG8145%') AND (deskripsi like '%Huawei%') AND (status LIKE '%AVAILABLE%' OR status LIKE '%INTECHNICIAN%') GROUP BY lokasi_wh")		
	}

	helper.PanicIfError(err)

	for rows.Next(){
		var countRetail domain.CountResponse
		rows.Scan(&countRetail.LokasiWH, &countRetail.Stock)

		countRetails = append(countRetails, countRetail)
	}

	return countRetails
}

func (r *dataTmpRepository) CountPremiumPerWitel(merk string) []domain.CountResponse{
	var countPremiums []domain.CountResponse
	var rows *sql.Rows
	var err error

	if(merk == "zte"){
		rows, err = r.db.Query("SELECT lokasi_wh, sum(jumlah) as stock FROM `data_tmp` WHERE (deskripsi LIKE '%ZTE_F670 V2.0%' or deskripsi = 'ONT_ZTE_F670 V2.0' or deskripsi = 'ONT_ZTE_F670') AND (deskripsi like '%ZTE%') AND (status LIKE '%AVAILABLE%' OR status LIKE '%INTECHNICIAN%') GROUP BY lokasi_wh");
	}else if(merk == "fiberhome"){
		rows, err = r.db.Query("SELECT lokasi_wh, sum(jumlah) as stock FROM `data_tmp` WHERE (deskripsi LIKE '%HG6245N%') AND (deskripsi like '%Fiberhome%') AND (status LIKE '%AVAILABLE%' OR status LIKE '%INTECHNICIAN%') GROUP BY lokasi_wh");
	}else if(merk == "huawei"){
		rows, err = r.db.Query("SELECT lokasi_wh, sum(jumlah) as stock FROM `data_tmp` WHERE (deskripsi = 'ONT_HUAWEI HG8245W5-6T') AND (deskripsi like '%Huawei%') AND (status LIKE '%AVAILABLE%' OR status LIKE '%INTECHNICIAN%') GROUP BY lokasi_wh");
	}

	helper.PanicIfError(err)

	for rows.Next(){
		var countPremium domain.CountResponse
		rows.Scan(&countPremium.LokasiWH, &countPremium.Stock)

		countPremiums = append(countPremiums, countPremium)
	}

	return countPremiums
}

func (r *dataTmpRepository) CountSTBPerWitel(merk string) []domain.CountResponse{
	var countSTBs []domain.CountResponse
	var rows *sql.Rows
	var err error

	if(merk == "zte"){
		rows, err = r.db.Query("SELECT lokasi_wh, sum(jumlah) as stock FROM `data_tmp` WHERE (deskripsi = 'SetTopBox_ZTE_ZX10_B866F_V1.1') AND (deskripsi like '%ZTE%') AND (status LIKE '%AVAILABLE%' OR status LIKE '%INTECHNICIAN%') GROUP BY lokasi_wh")
	}

	helper.PanicIfError(err)

	for rows.Next(){
		var countSTB domain.CountResponse
		rows.Scan(&countSTB.LokasiWH, &countSTB.Stock)

		countSTBs = append(countSTBs, countSTB)
	}

	return countSTBs
}

func (r *dataTmpRepository) CountAPPerWitel(merk string) []domain.CountResponse{
	var countAPs []domain.CountResponse
	var rows *sql.Rows
	var err error

	if(merk == "zte"){
		rows, err = r.db.Query("SELECT lokasi_wh, sum(jumlah) as stock FROM `data_tmp` WHERE (deskripsi = 'SetTopBox_ZTE_ZX10_B866F_V1.1') AND (deskripsi like '%ZTE%') AND (status LIKE '%AVAILABLE%' OR status LIKE '%INTECHNICIAN%') GROUP BY lokasi_wh")
	}

	helper.PanicIfError(err)

	for rows.Next(){
		var countAP domain.CountResponse
		rows.Scan(&countAP.LokasiWH, &countAP.Stock)

		countAPs = append(countAPs, countAP)
	}

	return countAPs
}

func (r *dataTmpRepository) GetWitelsFromDataByMerk(merk string) []string{
	var witels [] string

	rows, err := r.db.Query("SELECT DISTINCT witel FROM data WHERE merk = ?", merk)
	helper.PanicIfError(err)

	for rows.Next(){
		var witel string
		rows.Scan(&witel)

		witels = append(witels, witel)
	}

	return witels
}

