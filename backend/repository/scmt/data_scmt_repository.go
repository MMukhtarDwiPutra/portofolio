package repository

import (
	"database/sql"
	"log"
	"portofolio.com/domain/scmt"
	"portofolio.com/api/helper"
	// "fmt"
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
	GetExportDataTmp() []domain.DataTmp
	DeleteAllDataTmp()
	UploadDataTmpBulk(dataTmps []domain.DataTmp)
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
	var waktuUpdate, waktuDibuat sql.NullString

	rows, err := r.db.Query("SELECT UPDATE_TIME as waktu_update, CREATE_TIME as waktu_dibuat FROM information_schema.tables WHERE TABLE_SCHEMA = 'scmt' AND TABLE_NAME = 'data_tmp'")
	helper.PanicIfError(err)

	if rows.Next() {
        err := rows.Scan(&waktuUpdate, &waktuDibuat)
        if err != nil {
            return "", ""
        }
    } else {
        return "", ""
    }

    // Ubah sql.NullString ke string biasa
    waktuUpdateStr := waktuUpdate.String
    waktuDibuatStr := waktuDibuat.String

	return waktuUpdateStr, waktuDibuatStr
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

	if(merk == "ZTE"){
		rows, err = r.db.Query("SELECT lokasi_wh, sum(jumlah) as stock FROM data_tmp WHERE (deskripsi LIKE '%ONT_ZTE_F670L%') AND (deskripsi like '%ZTE%') AND (status LIKE '%AVAILABLE%' OR status LIKE '%INTECHNICIAN%') GROUP BY lokasi_wh")
	}else if(merk == "Nokia"){
		rows, err = r.db.Query("SELECT lokasi_wh, sum(jumlah) as stock FROM `data_tmp` WHERE (deskripsi like '%G240WL%' OR deskripsi like '%2425G%') AND (status LIKE '%AVAILABLE%' OR status LIKE '%INTECHNICIAN%') GROUP BY lokasi_wh")
	}else if(merk == "Fiberhome"){
		rows, err = r.db.Query("SELECT lokasi_wh, sum(jumlah) as stock FROM `data_tmp` WHERE (deskripsi LIKE '%HG6145D2%' or deskripsi LIKE '%HG6145F%') AND (deskripsi like '%Fiberhome%') AND (status LIKE '%AVAILABLE%' OR status LIKE '%INTECHNICIAN%') GROUP BY lokasi_wh")
	}else if(merk == "Huawei"){
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

	if(merk == "ZTE"){
		rows, err = r.db.Query("SELECT lokasi_wh, sum(jumlah) as stock FROM `data_tmp` WHERE (deskripsi LIKE '%ZTE_F670 V2.0%' or deskripsi = 'ONT_ZTE_F670 V2.0' or deskripsi = 'ONT_ZTE_F670') AND (deskripsi like '%ZTE%') AND (status LIKE '%AVAILABLE%' OR status LIKE '%INTECHNICIAN%') GROUP BY lokasi_wh");
	}else if(merk == "Fiberhome"){
		rows, err = r.db.Query("SELECT lokasi_wh, sum(jumlah) as stock FROM `data_tmp` WHERE (deskripsi LIKE '%HG6245N%') AND (deskripsi like '%Fiberhome%') AND (status LIKE '%AVAILABLE%' OR status LIKE '%INTECHNICIAN%') GROUP BY lokasi_wh");
	}else if(merk == "Huawei"){
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

	likePattern := "%" + merk + "%"
	rows, err := r.db.Query("SELECT DISTINCT lokasi_wh FROM data_tmp WHERE deskripsi like ?", likePattern)
	helper.PanicIfError(err)

	for rows.Next(){
		var witel string
		rows.Scan(&witel)

		witels = append(witels, witel)
	}

	return witels
}

func (r *dataTmpRepository) GetExportDataTmp() []domain.DataTmp{
	var dataTmpExports []domain.DataTmp
	
	rows, err := r.db.Query("SELECT * FROM data_tmp")
	helper.PanicIfError(err)

	for rows.Next(){
		var dataTmpExport domain.DataTmp

		err = rows.Scan(&dataTmpExport.ID, &dataTmpExport.Region, &dataTmpExport.LokasiWH, &dataTmpExport.Status, &dataTmpExport.Jumlah, &dataTmpExport.Deskripsi)
		helper.PanicIfError(err)

		dataTmpExports = append(dataTmpExports, dataTmpExport)
	}

	return dataTmpExports
}

func (r *dataTmpRepository) DeleteAllDataTmp(){
	_, err := r.db.Query("DELETE FROM data_tmp")
	helper.PanicIfError(err)
}

func (r *dataTmpRepository) UploadDataTmpBulk(dataTmps []domain.DataTmp){
	query := "INSERT INTO `data_tmp`(`region`, `lokasi_wh`, `status`, `jumlah`, `deskripsi`) VALUES "

	for i := range dataTmps{
		query += "( '"+dataTmps[i].Region+"','"+dataTmps[i].LokasiWH+"','"+ dataTmps[i].Status+"','"+ dataTmps[i].Jumlah+"','"+ dataTmps[i].Deskripsi+"'),"
	}

	query = query[:len(query)-1]
	_, err := r.db.Query(query)
	helper.PanicIfError(err)
}