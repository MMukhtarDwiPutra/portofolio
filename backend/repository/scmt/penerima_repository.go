package repository

import(
	"database/sql"
	"portofolio.com/domain/scmt"
	"portofolio.com/api/helper"
	"fmt"
)

type PenerimaRepository interface{
	CountRetailPerWitel(merk string) []domain.CountResponse
	CountPremiumPerWitel(merk string) []domain.CountResponse
	GetTableLastUpdate() (string, string)
	GetAllDataONT() []domain.PenerimaResponse
}

type penerimaRepository struct{
	db *sql.DB
}

func NewPenerimaRepository(db *sql.DB) *penerimaRepository{
	return &penerimaRepository{db}
}

func (r *penerimaRepository) CountRetailPerWitel(merk string) []domain.CountResponse{
	var countPenerimas []domain.CountResponse
	var rows *sql.Rows
	var err error

	if (merk == "ZTE"){
		rows, err = r.db.Query("SELECT `warehouse_penerima`, sum(qty) as stock FROM `penerimaan` WHERE (`type` LIKE '%ONT_ZTE_F670L%') AND (`type` like '%ZTE%') GROUP BY `warehouse_penerima`")
	}else if(merk == "Nokia"){
		rows, err = r.db.Query("SELECT `warehouse_penerima`, sum(qty) as stock FROM `penerimaan` WHERE (`type` like '%G240WL%' OR `type` like '%2425G%') GROUP BY `warehouse_penerima`")
	}else if(merk == "Fiberhome"){
		rows, err = r.db.Query("SELECT `warehouse_penerima`, sum(qty) as stock FROM `penerimaan` WHERE (`type` LIKE '%HG6145D2%' or `type` LIKE '%HG6145F%') AND (`type` like '%Fiberhome%') GROUP BY `warehouse_penerima`")
	}else if(merk == "Huawei"){
		rows, err = r.db.Query("SELECT `warehouse_penerima`, sum(qty) as stock FROM `penerimaan` WHERE ((`type` LIKE '%HG8145%') OR (`type` like '%HG8145V5%')) GROUP BY `warehouse_penerima`")
	}
	helper.PanicIfError(err)

	for rows.Next(){
		var countPenerima domain.CountResponse
		rows.Scan(&countPenerima.LokasiWH, &countPenerima.Stock)

		countPenerimas = append(countPenerimas, countPenerima)
	}

	return countPenerimas
}

func (r *penerimaRepository) CountPremiumPerWitel(merk string) []domain.CountResponse{
	var countPenerimas []domain.CountResponse
	var rows *sql.Rows
	var err error

	if (merk == "ZTE"){
		rows, err = r.db.Query("SELECT `warehouse_penerima`, sum(qty) as stock FROM `penerimaan` WHERE (`type` LIKE '%ZTE_F670 V2.0%' or `type` = 'ONT_ZTE_F670 V2.0' or `type` = 'ONT_ZTE_F670') AND (`type` like '%ZTE%')  GROUP BY `warehouse_penerima`")
	}else if(merk == "Fiberhome"){
		rows, err = r.db.Query("SELECT `warehouse_penerima`, sum(qty) as stock FROM `penerimaan` WHERE (`type` LIKE '%HG6245N%') AND (`type` like '%Fiberhome%')  GROUP BY `warehouse_penerima`")
	}else if(merk == "Huawei"){
		rows, err = r.db.Query("SELECT `warehouse_penerima`, sum(qty) as stock FROM `penerimaan` WHERE (`type` = 'ONT_HUAWEI HG8245W5-6T' or `type` = 'ONT_HW_HG8245W5-6T') GROUP BY `warehouse_penerima`")
	}
	helper.PanicIfError(err)

	for rows.Next(){
		var countPenerima domain.CountResponse
		rows.Scan(&countPenerima.LokasiWH, &countPenerima.Stock)

		countPenerimas = append(countPenerimas, countPenerima)
	}

	return countPenerimas
}

func (r *penerimaRepository) GetAllDataONT() []domain.PenerimaResponse{
	var dataONTs []domain.PenerimaResponse
	rows, err := r.db.Query("select p.id, p.type, p.qty, p.alamat_pengirim, p.pic_pengirim, p.alamat_penerima, p.warehouse_penerima, p.pic_penerima, p.tanggal_pengiriman, p.tanggal_sampai, p.ido_gd, p.sn_mac_barcode, p.batch, p.ido_gd_time_added, p.sn_time_added, g.regional from penerimaan p JOIN (select gd.lokasi_wh, gd.regional, row_number() over (partition by gd.lokasi_wh) as row_number from gudang gd GROUP BY gd.lokasi_wh,gd.regional) g ON g.lokasi_wh = p.warehouse_penerima where (g.row_number = 1) AND ((p.type LIKE '%ONT_ZTE_F670L%') OR (p.type like '%G240WL%' OR type like '%2425G%') OR (p.type LIKE '%HG6145D2%' or p.type LIKE '%HG6145F%') OR (p.type LIKE '%HG8145%') OR (p.type LIKE '%ZTE_F670 V2.0%' or type = 'ONT_ZTE_F670 V2.0' or p.type = 'ONT_ZTE_F670') OR (p.type LIKE '%HG6245N%') OR (p.type LIKE '%HG8245W5-6T') OR (`type` like '%HG8145V5%')) ORDER BY p.batch")
	helper.PanicIfError(err)

	for rows.Next(){
		var dataONT domain.PenerimaResponse
		err = rows.Scan(&dataONT.ID, &dataONT.Type, &dataONT.Qty, &dataONT.AlamatPengirim, &dataONT.PICPengirim, &dataONT.AlamatPenerima, &dataONT.WarehousePenerima, &dataONT.PICPenerima, &dataONT.TanggalPengiriman, &dataONT.TanggalSampai, &dataONT.IDOGD, &dataONT.SNMacBarcode, &dataONT.Batch, &dataONT.IDOGDTimeAdded, &dataONT.SNTimeAdded, &dataONT.Regional)
		helper.PanicIfError(err)

		dataONTs = append(dataONTs, dataONT)
	}
	fmt.Println(dataONTs)

	return dataONTs
}

func (r *penerimaRepository) GetTableLastUpdate() (string, string){
	var waktuUpdate, waktuDibuat sql.NullString
	
	rows, err := r.db.Query("SELECT UPDATE_TIME as waktu_update, CREATE_TIME as waktu_dibuat FROM information_schema.tables WHERE TABLE_SCHEMA = 'scmt' AND TABLE_NAME = 'penerimaan'")
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

