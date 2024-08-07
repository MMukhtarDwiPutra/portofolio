package repository

import (
	"database/sql"
	"portofolio.com/domain/scmt"
	"portofolio.com/api/helper"
)

type GudangRepository interface{
	GetTREGQtyMinimum() []domain.TREGMinimumResponse
	GetQtyMinimum() []domain.TREGMinimumResponse
	GetAllSOFromTREG(witel string) []string
	GetAllSOFromWitel(witel string) []string
	GetSOFromSO(witel string) []string
	GetAllWarehouse() []domain.LokasiWarehouseResponse
	GetAllData() []domain.Gudang
	UploadGudangBulk(gudang []domain.Gudang)
	DeleteAllDataGudang()
}

type gudangRepository struct{
	db *sql.DB
}

func NewGudangRepository(db *sql.DB) *gudangRepository{
	return &gudangRepository{db};
}

func (r *gudangRepository) GetTREGQtyMinimum() []domain.TREGMinimumResponse{
	var TREGMinimumResponses []domain.TREGMinimumResponse
	rows, err := r.db.Query("SELECT CONCAT('1000', ROW_NUMBER() OVER(ORDER BY regional ASC)) AS id, regional, SUM(minimum_qty) as minimum_qty, regional as 'lokasi_wh', regional as 'witel', regional as 'lokasi', regional as 'wilayah', SUM(retail_fh) as 'retail_fh', SUM(retail_zte) as 'retail_zte', SUM(retail_hw) as 'retail_hw', SUM(retail_alu) as 'retail_alu', SUM(premium_fh) as 'premium_fh', SUM(premium_zte) as 'premium_zte', SUM(premium_hw) as 'premium_hw', CAST(ROUND(SUM(retail_fh) * 120 / 100) as int) as 'batas_atas_retail_fh', CAST(ROUND(SUM(retail_zte) * 120 / 100) as int) as 'batas_atas_retail_zte', CAST(ROUND(SUM(retail_hw) * 120 / 100) as int) as 'batas_atas_retail_hw', CAST(ROUND(SUM(retail_alu) * 120 / 100) as int) as 'batas_atas_retail_alu', CAST(ROUND(SUM(premium_fh) * 120 / 100) as int) as 'batas_atas_premium_fh', CAST(ROUND(SUM(premium_hw) * 120 / 100) as int) as 'batas_atas_premium_hw', CAST(ROUND(SUM(premium_zte) * 120 / 100) as int) as 'batas_atas_premium_zte', CAST(ROUND(SUM(retail_fh) * 70 / 100) as int) as 'batas_bawah_retail_fh', CAST(ROUND(SUM(retail_zte) * 70 / 100) as int) as 'batas_bawah_retail_zte', CAST(ROUND(SUM(retail_hw) * 70 / 100) as int) as 'batas_bawah_retail_hw', CAST(ROUND(SUM(retail_alu) * 70 / 100) as int) as 'batas_bawah_retail_alu', CAST(ROUND(SUM(premium_fh) * 70 / 100) as int) as 'batas_bawah_premium_fh', CAST(ROUND(SUM(premium_hw) * 70 / 100) as int) as 'batas_bawah_premium_hw', CAST(ROUND(SUM(premium_zte) * 70 / 100) as int) as 'batas_bawah_premium_zte' FROM gudang GROUP BY regional")
	helper.PanicIfError(err)

	for rows.Next(){
		var TREGMinimumResponse domain.TREGMinimumResponse
		rows.Scan(&TREGMinimumResponse.ID, &TREGMinimumResponse.Regional, &TREGMinimumResponse.MinimumQty, &TREGMinimumResponse.Witel, &TREGMinimumResponse.Lokasi, &TREGMinimumResponse.Wilayah, &TREGMinimumResponse.RetailFH, &TREGMinimumResponse.RetailZTE, &TREGMinimumResponse.RetailHW, &TREGMinimumResponse.RetailALU, &TREGMinimumResponse.PremiumFH, &TREGMinimumResponse.PremiumZTE, &TREGMinimumResponse.PremiumHW, &TREGMinimumResponse.BatasAtasRetailFH, &TREGMinimumResponse.BatasAtasRetailZTE, &TREGMinimumResponse.BatasAtasRetailHW, &TREGMinimumResponse.BatasAtasRetailALU, &TREGMinimumResponse.BatasAtasPremiumFH, &TREGMinimumResponse.BatasAtasPremiumZTE, &TREGMinimumResponse.BatasAtasPremiumHW, &TREGMinimumResponse.BatasBawahRetailFH, &TREGMinimumResponse.BatasBawahRetailZTE, &TREGMinimumResponse.BatasBawahRetailHW, &TREGMinimumResponse.BatasBawahRetailALU, &TREGMinimumResponse.BatasBawahPremiumFH, &TREGMinimumResponse.BatasBawahPremiumZTE, &TREGMinimumResponse.BatasBawahPremiumHW)

		TREGMinimumResponses = append(TREGMinimumResponses, TREGMinimumResponse)
	}

	return TREGMinimumResponses
}

func (r *gudangRepository) GetQtyMinimum() []domain.TREGMinimumResponse{
	var QtyMinimumResponses []domain.TREGMinimumResponse

	rows, err := r.db.Query("SELECT `id`, `regional`, `witel`, `lokasi_wh`, `lokasi`, `wilayah`, `minimum_qty`, `retail_fh`, `retail_zte`, `retail_hw`, `retail_alu`, `premium_zte`, `premium_fh`, `premium_hw`, CAST(ROUND(retail_fh * 120 / 100) as int) as 'batas_atas_retail_fh', CAST(ROUND(retail_zte * 120 / 100) as int) as 'batas_atas_retail_zte', CAST(ROUND(retail_hw * 120 / 100) as int) as 'batas_atas_retail_hw', CAST(ROUND(retail_alu * 120 / 100) as int) as 'batas_atas_retail_alu', CAST(ROUND(premium_fh * 120 / 100) as int) as 'batas_atas_premium_fh', CAST(ROUND(premium_hw * 120 / 100) as int) as 'batas_atas_premium_hw', CAST(ROUND(premium_zte * 120 / 100) as int) as 'batas_atas_premium_zte', CAST(ROUND(retail_fh * 70 / 100) as int) as 'batas_bawah_retail_fh', CAST(ROUND(retail_zte * 70 / 100) as int) as 'batas_bawah_retail_zte', CAST(ROUND(retail_hw * 70 / 100) as int) as 'batas_bawah_retail_hw', CAST(ROUND(retail_alu * 70 / 100) as int) as 'batas_bawah_retail_alu', CAST(ROUND(premium_fh * 70 / 100) as int) as 'batas_bawah_premium_fh', CAST(ROUND(premium_hw * 70 / 100) as int) as 'batas_bawah_premium_hw', CAST(ROUND(premium_zte * 70 / 100) as int) as 'batas_bawah_premium_zte' FROM gudang ORDER BY id")
	helper.PanicIfError(err)

	for rows.Next(){
		var QtyMinimumResponse domain.TREGMinimumResponse
		rows.Scan(&QtyMinimumResponse.ID, &QtyMinimumResponse.Regional, &QtyMinimumResponse.Witel, &QtyMinimumResponse.LokasiWH, &QtyMinimumResponse.Lokasi, &QtyMinimumResponse.Wilayah, &QtyMinimumResponse.MinimumQty, &QtyMinimumResponse.RetailFH, &QtyMinimumResponse.RetailZTE, &QtyMinimumResponse.RetailHW, &QtyMinimumResponse.RetailALU, &QtyMinimumResponse.PremiumZTE, &QtyMinimumResponse.PremiumFH, &QtyMinimumResponse.PremiumHW, &QtyMinimumResponse.BatasAtasRetailFH, &QtyMinimumResponse.BatasAtasRetailZTE, &QtyMinimumResponse.BatasAtasRetailHW, &QtyMinimumResponse.BatasAtasRetailALU, &QtyMinimumResponse.BatasAtasPremiumFH, &QtyMinimumResponse.BatasAtasPremiumHW, &QtyMinimumResponse.BatasAtasPremiumZTE, &QtyMinimumResponse.BatasBawahRetailFH, &QtyMinimumResponse.BatasBawahRetailZTE, &QtyMinimumResponse.BatasBawahRetailHW, &QtyMinimumResponse.BatasBawahRetailALU, &QtyMinimumResponse.BatasBawahPremiumFH, &QtyMinimumResponse.BatasBawahPremiumHW, &QtyMinimumResponse.BatasBawahPremiumZTE)

		QtyMinimumResponses = append(QtyMinimumResponses, QtyMinimumResponse)
	}

	return QtyMinimumResponses
}

func (r *gudangRepository) GetAllSOFromTREG(witel string) []string{
	var witels []string

	rows, err := r.db.Query("SELECT DISTINCT `lokasi_wh` FROM gudang WHERE `regional` = ?", witel)
	helper.PanicIfError(err)

	for rows.Next(){
		var witel string
		rows.Scan(&witel)

		witels = append(witels, witel)
	}

	return witels
}

func (r *gudangRepository) GetAllSOFromWitel(witel string) []string{
	var witels []string

	rows, err := r.db.Query("SELECT DISTINCT `lokasi_wh` FROM gudang WHERE `witel` = ?", witel)
	helper.PanicIfError(err)

	for rows.Next(){
		var witel string
		rows.Scan(&witel)

		witels = append(witels, witel)
	}

	return witels
}

func (r *gudangRepository) GetSOFromSO(witel string) []string{
	var witels []string

	rows, err := r.db.Query("SELECT DISTINCT `lokasi_wh` FROM gudang WHERE `lokasi_wh` = ?", witel)
	helper.PanicIfError(err)

	for rows.Next(){
		var witel string
		rows.Scan(&witel)

		witels = append(witels, witel)
	}

	return witels
}

func (r *gudangRepository) GetAllWarehouse() []domain.LokasiWarehouseResponse{
	var warehouses []domain.LokasiWarehouseResponse

	rows, err := r.db.Query("SELECT lokasi_wh, lokasi FROM gudang")
	helper.PanicIfError(err)

	for rows.Next(){
		var warehouse domain.LokasiWarehouseResponse
		rows.Scan(&warehouse.LokasiWH, &warehouse.Lokasi)

		warehouses = append(warehouses, warehouse)
	}

	return warehouses
}

func (r *gudangRepository) GetAllData() []domain.Gudang{
	var dataGudangs []domain.Gudang

	rows, err := r.db.Query("SELECT * FROM GUDANG")
	helper.PanicIfError(err)

	for rows.Next(){
		var dataGudang domain.Gudang
		rows.Scan(&dataGudang.ID, &dataGudang.Regional, &dataGudang.Witel, &dataGudang.LokasiWH, &dataGudang.Lokasi, &dataGudang.Wilayah, &dataGudang.MinimumQty, &dataGudang.RetailZTE, &dataGudang.RetailHW, &dataGudang.RetailFH,&dataGudang.RetailALU, &dataGudang.PremiumZTE, &dataGudang.PremiumFH, &dataGudang.PremiumHW, &dataGudang.STBZTE, &dataGudang.APCisco, &dataGudang.APHuawei)
		dataGudangs = append(dataGudangs, dataGudang)
	}

	return dataGudangs
}

func (r *gudangRepository) UploadGudangBulk(gudang []domain.Gudang){
	query := "INSERT INTO `gudang`(`regional`, `witel`, `lokasi_wh`, `lokasi`, `wilayah`, `minimum_qty`, `retail_zte`, `retail_hw`, `retail_fh`, `retail_alu`, `premium_zte`, `premium_fh`, `premium_hw`, `stb_zte`, `ap_cisco`, `ap_huawei`) VALUE "

	for i := range gudang{
		query += "( '"+gudang[i].Regional+"','"+gudang[i].Witel+"','"+ gudang[i].LokasiWH+"','"+ gudang[i].Lokasi+"','"+ gudang[i].Wilayah+"','"+ gudang[i].MinimumQty+"','"+ gudang[i].RetailZTE+"','"+ gudang[i].RetailHW+"','"+ gudang[i].RetailFH+"','"+ gudang[i].RetailALU+"','"+ gudang[i].PremiumZTE+"','"+ gudang[i].PremiumFH+"','"+ gudang[i].PremiumHW+"','"+ gudang[i].STBZTE+"','"+ gudang[i].APCisco+"','"+ gudang[i].APHuawei+"'),"
	}

	query = query[:len(query)-1]
	_, err := r.db.Query(query)
	helper.PanicIfError(err)
}

func (r *gudangRepository) DeleteAllDataGudang(){
	_, err := r.db.Query("DELETE FROM gudang")
	helper.PanicIfError(err)
}