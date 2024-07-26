package repository

import (
	"database/sql"
	"portofolio.com/domain/scmt"
	"portofolio.com/api/helper"
)

type GudangRepository interface{
	GetTREGQtyMinimum() []domain.TREGMinimumResponse
	GetQtyMinimum() []domain.TREGMinimumResponse
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