package repository

import(
	"database/sql"
	"portofolio.com/domain/scmt"
	"portofolio.com/api/helper"
	// "fmt"
	"time"
)

type PenerimaRepository interface{
	CountRetailPerWitel(merk string) []domain.CountResponse
	CountPremiumPerWitel(merk string) []domain.CountResponse
	GetTableLastUpdate() (string, string)
	GetAllDataONT() []domain.PenerimaResponse
	GetAllPenerimaExport() []domain.PenerimaResponse
	GetAllDataONTExport() []domain.PenerimaResponse
	GetAllDataSTBExport() []domain.PenerimaResponse
	GetAllDataAPExport() []domain.PenerimaResponse
	GetAllSNSTBExist() []domain.PenerimaResponse
	GetAllSNONTExist() []domain.PenerimaResponse
	GetAllSNSTB() []domain.PenerimaResponse
	GetAllSNONT() []domain.PenerimaResponse
	GetSNBatchById(id int) (string, string)
	AddPenerima(penerima domain.PenerimaPost)
	DeletePenerimaById(id int)
	GetSNById(id int) string
	GetDataById(id int) domain.PenerimaResponse
	EditIDOGDById(id string, data domain.PenerimaPost)
	EditTanggalPenerimaanById(id string, data domain.PenerimaPost)
	EditTanggalOnly(id string, data domain.PenerimaPost)
	AddPenerimaBulk(penerima []domain.PenerimaPost)
	DeleteAllPenerima()
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

func (r *penerimaRepository) GetAllPenerimaExport() []domain.PenerimaResponse{
	var penerimaExports []domain.PenerimaResponse
	
	rows, err := r.db.Query("select p.id, p.type, p.qty, p.alamat_pengirim, p.pic_pengirim, p.alamat_penerima, p.warehouse_penerima, p.pic_penerima, p.tanggal_pengiriman, p.tanggal_sampai, p.ido_gd, p.sn_mac_barcode, p.ido_gd_time_added, p.time_added, g.regional, p.batch, p.sn_time_added from penerimaan p JOIN (select gd.lokasi_wh, gd.regional, row_number() over (partition by gd.lokasi_wh) as row_number from gudang gd GROUP BY gd.lokasi_wh,gd.regional) g ON g.lokasi_wh = p.warehouse_penerima where (g.row_number = 1)")
	helper.PanicIfError(err)

	for rows.Next(){
		var penerimaExport domain.PenerimaResponse

		err = rows.Scan(&penerimaExport.ID, &penerimaExport.Type, &penerimaExport.Qty, &penerimaExport.AlamatPengirim, &penerimaExport.PICPengirim, &penerimaExport.AlamatPenerima, &penerimaExport.WarehousePenerima, &penerimaExport.PICPenerima, &penerimaExport.TanggalPengiriman, &penerimaExport.TanggalSampai, &penerimaExport.IDOGD, &penerimaExport.SNMacBarcode, &penerimaExport.IDOGDTimeAdded, &penerimaExport.TimeAdded, &penerimaExport.Regional, &penerimaExport.Batch, &penerimaExport.SNTimeAdded)
		helper.PanicIfError(err)

		penerimaExports = append(penerimaExports, penerimaExport)
	}

	return penerimaExports
}

func (r *penerimaRepository) GetAllDataONTExport() []domain.PenerimaResponse{
	var penerimaExports []domain.PenerimaResponse
	
	rows, err := r.db.Query("select p.id, p.type, p.qty, p.alamat_pengirim, p.pic_pengirim, p.alamat_penerima, p.warehouse_penerima, p.pic_penerima, p.tanggal_pengiriman, p.tanggal_sampai, p.ido_gd, p.sn_mac_barcode, p.batch, p.ido_gd_time_added, p.sn_time_added, g.regional, p.time_added from penerimaan p JOIN (select gd.lokasi_wh, gd.regional, row_number() over (partition by gd.lokasi_wh) as row_number from gudang gd GROUP BY gd.lokasi_wh,gd.regional) g ON g.lokasi_wh = p.warehouse_penerima where (g.row_number = 1) AND ((p.type LIKE '%ONT_ZTE_F670L%') OR (p.type like '%G240WL%' OR type like '%2425G%') OR (p.type LIKE '%HG6145D2%' or p.type LIKE '%HG6145F%') OR (p.type LIKE '%HG8145%') OR (p.type LIKE '%ZTE_F670 V2.0%' or type = 'ONT_ZTE_F670 V2.0' or p.type = 'ONT_ZTE_F670') OR (p.type LIKE '%HG6245N%') OR (p.type LIKE '%HG8245W5-6T') OR (`type` like '%HG8145V5%')) ORDER BY p.batch")
	helper.PanicIfError(err)

	for rows.Next(){
		var penerimaExport domain.PenerimaResponse

		err = rows.Scan(&penerimaExport.ID, &penerimaExport.Type, &penerimaExport.Qty, &penerimaExport.AlamatPengirim, &penerimaExport.PICPengirim, &penerimaExport.AlamatPenerima, &penerimaExport.WarehousePenerima, &penerimaExport.PICPenerima, &penerimaExport.TanggalPengiriman, &penerimaExport.TanggalSampai, &penerimaExport.IDOGD, &penerimaExport.SNMacBarcode, &penerimaExport.Batch, &penerimaExport.IDOGDTimeAdded, &penerimaExport.SNTimeAdded, &penerimaExport.Regional, &penerimaExport.TimeAdded)
		helper.PanicIfError(err)

		penerimaExports = append(penerimaExports, penerimaExport)
	}

	return penerimaExports
}

func (r *penerimaRepository) GetAllDataSTBExport() []domain.PenerimaResponse{
	var penerimaExports []domain.PenerimaResponse
	
	rows, err := r.db.Query("SELECT p.id, p.type, p.qty, p.alamat_pengirim, p.pic_pengirim, p.alamat_penerima, p.warehouse_penerima, p.pic_penerima, p.tanggal_pengiriman, p.tanggal_sampai, p.ido_gd, p.sn_mac_barcode, p.batch, p.ido_gd_time_added, p.sn_time_added, g.regional, p.time_added FROM penerimaan p JOIN gudang g ON p.warehouse_penerima = g.lokasi_wh WHERE (type = 'SetTopBox_ZTE_ZX10_B866F_V1.1')")
	helper.PanicIfError(err)

	for rows.Next(){
		var penerimaExport domain.PenerimaResponse

		err = rows.Scan(&penerimaExport.ID, &penerimaExport.Type, &penerimaExport.Qty, &penerimaExport.AlamatPengirim, &penerimaExport.PICPengirim, &penerimaExport.AlamatPenerima, &penerimaExport.WarehousePenerima, &penerimaExport.PICPenerima, &penerimaExport.TanggalPengiriman, &penerimaExport.TanggalSampai, &penerimaExport.IDOGD, &penerimaExport.SNMacBarcode, &penerimaExport.Batch, &penerimaExport.IDOGDTimeAdded, &penerimaExport.SNTimeAdded, &penerimaExport.Regional, &penerimaExport.TimeAdded)
		helper.PanicIfError(err)

		penerimaExports = append(penerimaExports, penerimaExport)
	}

	return penerimaExports
}

func (r *penerimaRepository) GetAllDataAPExport() []domain.PenerimaResponse{
	var penerimaExports []domain.PenerimaResponse
	
	rows, err := r.db.Query("SELECT p.id, p.type, p.qty, p.alamat_pengirim, p.pic_pengirim, p.alamat_penerima, p.warehouse_penerima, p.pic_penerima, p.tanggal_pengiriman, p.tanggal_sampai, p.ido_gd, p.sn_mac_barcode, p.batch, p.ido_gd_time_added, p.sn_time_added, g.regional, p.time_added FROM penerimaan p JOIN gudang g ON p.warehouse_penerima = g.lokasi_wh WHERE (type = 'AP_CISCO_AIR-AP2802I-F-K9' OR type = 'AP_CISCO_C9105AXI-F' OR type = 'AP_HUAWEI_AIRENGINE5761-11')")
	helper.PanicIfError(err)

	for rows.Next(){
		var penerimaExport domain.PenerimaResponse

		err = rows.Scan(&penerimaExport.ID, &penerimaExport.Type, &penerimaExport.Qty, &penerimaExport.AlamatPengirim, &penerimaExport.PICPengirim, &penerimaExport.AlamatPenerima, &penerimaExport.WarehousePenerima, &penerimaExport.PICPenerima, &penerimaExport.TanggalPengiriman, &penerimaExport.TanggalSampai, &penerimaExport.IDOGD, &penerimaExport.SNMacBarcode, &penerimaExport.Batch, &penerimaExport.IDOGDTimeAdded, &penerimaExport.SNTimeAdded, &penerimaExport.Regional, &penerimaExport.TimeAdded)
		helper.PanicIfError(err)

		penerimaExports = append(penerimaExports, penerimaExport)
	}

	return penerimaExports
}

func (r *penerimaRepository) GetAllSNSTBExist() []domain.PenerimaResponse{
	var penerimaExports []domain.PenerimaResponse
	
	rows, err := r.db.Query("SELECT p.id, p.type, p.qty, p.alamat_pengirim, p.pic_pengirim, p.alamat_penerima, p.warehouse_penerima, p.pic_penerima, p.tanggal_pengiriman, p.tanggal_sampai, p.ido_gd, p.sn_mac_barcode, p.batch, p.ido_gd_time_added, p.sn_time_added, g.regional, p.time_added FROM penerimaan p JOIN gudang g ON p.warehouse_penerima = g.lokasi_wh WHERE p.sn_mac_barcode <> '' AND p.ido_gd = '' AND (p.type = 'SetTopBox_ZTE_ZX10_B866F_V1.1')")
	helper.PanicIfError(err)

	for rows.Next(){
		var penerimaExport domain.PenerimaResponse

		err = rows.Scan(&penerimaExport.ID, &penerimaExport.Type, &penerimaExport.Qty, &penerimaExport.AlamatPengirim, &penerimaExport.PICPengirim, &penerimaExport.AlamatPenerima, &penerimaExport.WarehousePenerima, &penerimaExport.PICPenerima, &penerimaExport.TanggalPengiriman, &penerimaExport.TanggalSampai, &penerimaExport.IDOGD, &penerimaExport.SNMacBarcode, &penerimaExport.Batch, &penerimaExport.IDOGDTimeAdded, &penerimaExport.SNTimeAdded, &penerimaExport.Regional, &penerimaExport.TimeAdded)
		helper.PanicIfError(err)

		penerimaExports = append(penerimaExports, penerimaExport)
	}

	return penerimaExports
}

func (r *penerimaRepository) GetAllSNONTExist() []domain.PenerimaResponse{
	var penerimaExports []domain.PenerimaResponse
	
	rows, err := r.db.Query("select p.id, p.type, p.qty, p.alamat_pengirim, p.pic_pengirim, p.alamat_penerima, p.warehouse_penerima, p.pic_penerima, p.tanggal_pengiriman, p.tanggal_sampai, p.ido_gd, p.sn_mac_barcode, p.ido_gd_time_added, p.time_added, g.regional, p.batch, p.sn_time_added from penerimaan p JOIN gudang g ON g.lokasi_wh = p.warehouse_penerima where p.sn_mac_barcode <> '' AND p.ido_gd = '' AND ((p.type LIKE '%ONT_ZTE_F670L%') OR (p.type like '%G240WL%' OR type like '%2425G%') OR (p.type LIKE '%HG6145D2%' or p.type LIKE '%HG6145F%') OR (p.type LIKE '%HG8145%') OR (p.type LIKE '%ZTE_F670 V2.0%' or type = 'ONT_ZTE_F670 V2.0' or p.type = 'ONT_ZTE_F670') OR (p.type LIKE '%HG6245N%') OR (p.type LIKE '%HG8245W5-6T') OR (`type` like '%HG8145V5%'))")
	helper.PanicIfError(err)

	for rows.Next(){
		var penerimaExport domain.PenerimaResponse

		err = rows.Scan(&penerimaExport.ID, &penerimaExport.Type, &penerimaExport.Qty, &penerimaExport.AlamatPengirim, &penerimaExport.PICPengirim, &penerimaExport.AlamatPenerima, &penerimaExport.WarehousePenerima, &penerimaExport.PICPenerima, &penerimaExport.TanggalPengiriman, &penerimaExport.TanggalSampai, &penerimaExport.IDOGD, &penerimaExport.SNMacBarcode, &penerimaExport.Batch, &penerimaExport.IDOGDTimeAdded, &penerimaExport.SNTimeAdded, &penerimaExport.Regional, &penerimaExport.TimeAdded)
		helper.PanicIfError(err)

		penerimaExports = append(penerimaExports, penerimaExport)
	}

	return penerimaExports
}

func (r *penerimaRepository) GetAllSNSTB() []domain.PenerimaResponse{
	var penerimaExports []domain.PenerimaResponse
	
	rows, err := r.db.Query("SELECT p.id, p.type, p.qty, p.alamat_pengirim, p.pic_pengirim, p.alamat_penerima, p.warehouse_penerima, p.pic_penerima, p.tanggal_pengiriman, p.tanggal_sampai, p.ido_gd, p.sn_mac_barcode, p.batch, p.ido_gd_time_added, p.sn_time_added, g.regional, p.time_added FROM penerimaan p JOIN gudang g ON p.warehouse_penerima = g.lokasi_wh WHERE p.sn_mac_barcode <> '' AND (p.type = 'SetTopBox_ZTE_ZX10_B866F_V1.1')")
	helper.PanicIfError(err)

	for rows.Next(){
		var penerimaExport domain.PenerimaResponse

		err = rows.Scan(&penerimaExport.ID, &penerimaExport.Type, &penerimaExport.Qty, &penerimaExport.AlamatPengirim, &penerimaExport.PICPengirim, &penerimaExport.AlamatPenerima, &penerimaExport.WarehousePenerima, &penerimaExport.PICPenerima, &penerimaExport.TanggalPengiriman, &penerimaExport.TanggalSampai, &penerimaExport.IDOGD, &penerimaExport.SNMacBarcode, &penerimaExport.Batch, &penerimaExport.IDOGDTimeAdded, &penerimaExport.SNTimeAdded, &penerimaExport.Regional, &penerimaExport.TimeAdded)
		helper.PanicIfError(err)

		penerimaExports = append(penerimaExports, penerimaExport)
	}

	return penerimaExports
}

func (r *penerimaRepository) GetAllSNONT() []domain.PenerimaResponse{
	var penerimaExports []domain.PenerimaResponse
	
	rows, err := r.db.Query("select p.id, p.type, p.qty, p.alamat_pengirim, p.pic_pengirim, p.alamat_penerima, p.warehouse_penerima, p.pic_penerima, p.tanggal_pengiriman, p.tanggal_sampai, p.ido_gd, p.sn_mac_barcode, p.ido_gd_time_added, p.time_added, g.regional, p.batch, p.sn_time_added from penerimaan p JOIN gudang g ON g.lokasi_wh = p.warehouse_penerima where p.sn_mac_barcode <> '' AND ((p.type LIKE '%ONT_ZTE_F670L%') OR (p.type like '%G240WL%' OR type like '%2425G%') OR (p.type LIKE '%HG6145D2%' or p.type LIKE '%HG6145F%') OR (p.type LIKE '%HG8145%') OR (p.type LIKE '%ZTE_F670 V2.0%' or type = 'ONT_ZTE_F670 V2.0' or p.type = 'ONT_ZTE_F670') OR (p.type LIKE '%HG6245N%') OR (p.type LIKE '%HG8245W5-6T') OR (`type` like '%HG8145V5%'))")
	helper.PanicIfError(err)

	for rows.Next(){
		var penerimaExport domain.PenerimaResponse

		err = rows.Scan(&penerimaExport.ID, &penerimaExport.Type, &penerimaExport.Qty, &penerimaExport.AlamatPengirim, &penerimaExport.PICPengirim, &penerimaExport.AlamatPenerima, &penerimaExport.WarehousePenerima, &penerimaExport.PICPenerima, &penerimaExport.TanggalPengiriman, &penerimaExport.TanggalSampai, &penerimaExport.IDOGD, &penerimaExport.SNMacBarcode, &penerimaExport.Batch, &penerimaExport.IDOGDTimeAdded, &penerimaExport.SNTimeAdded, &penerimaExport.Regional, &penerimaExport.TimeAdded)
		helper.PanicIfError(err)

		penerimaExports = append(penerimaExports, penerimaExport)
	}

	return penerimaExports
}

func (r *penerimaRepository) GetSNBatchById(id int) (string, string){
	var sn, batch string
	
	rows, err := r.db.Query("SELECT sn_mac_barcode, batch as 'penerimaan_batch' FROM penerimaan WHERE id = ?", id)
	helper.PanicIfError(err)

	if rows.Next(){
		err = rows.Scan(&sn, &batch)
		helper.PanicIfError(err)
	}

	return sn, batch
}

func (r *penerimaRepository) AddPenerima(penerima domain.PenerimaPost){
	_, err := r.db.Query("INSERT INTO `penerimaan`(`type`, `qty`, `alamat_pengirim`, `pic_pengirim`, `alamat_penerima`, `warehouse_penerima`, `pic_penerima`, `tanggal_pengiriman`, `tanggal_sampai`, `batch`) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",penerima.Type, penerima.Qty, penerima.AlamatPengirim, penerima.PICPengirim, penerima.AlamatPenerima, penerima.WarehousePenerima, penerima.PICPenerima, penerima.TanggalPengiriman, penerima.TanggalSampai, penerima.Batch)
	helper.PanicIfError(err)
}

func (r *penerimaRepository) AddPenerimaBulk(penerima []domain.PenerimaPost){
	query := "INSERT INTO `penerimaan`(`type`, `qty`, `alamat_pengirim`, `pic_pengirim`, `alamat_penerima`, `warehouse_penerima`, `pic_penerima`, `tanggal_pengiriman`, `tanggal_sampai`, `ido_gd`, `sn_mac_barcode`, `batch`, `ido_gd_time_added`, `sn_time_added`, `time_added`) VALUES "

	for i := range penerima{
		query += "( '"+penerima[i].Type+"','"+penerima[i].Qty+"','"+ penerima[i].AlamatPengirim+"','"+ penerima[i].PICPengirim+"','"+ penerima[i].AlamatPenerima+"','"+ penerima[i].WarehousePenerima+"','"+ penerima[i].PICPenerima+"','"+ penerima[i].TanggalPengiriman+"','"+ penerima[i].TanggalSampai+"','"+ penerima[i].IDOGD+"','"+ penerima[i].SNMacBarcode+"','"+ penerima[i].Batch+"','"+ penerima[i].IDOGDTimeAdded+"','"+ penerima[i].SNTimeAdded+"','"+ penerima[i].TimeAdded+"'),"
	}

	query = query[:len(query)-1]
	_, err := r.db.Query(query)
	helper.PanicIfError(err)
}
 
func (r *penerimaRepository) DeletePenerimaById(id int){
	_, err := r.db.Query("DELETE FROM penerimaan WHERE id = ?", id)
	helper.PanicIfError(err)
}

func (r *penerimaRepository) GetSNById(id int) string{
	var snMacBarcode string

	rows, err := r.db.Query("SELECT sn_mac_barcode FROM penerimaan WHERE id = ?",id)
	helper.PanicIfError(err)
	if rows.Next(){
		rows.Scan(&snMacBarcode)
	}else{
		return "SN tidak ada!"
	}

	return snMacBarcode
}

func (r *penerimaRepository) GetDataById(id int) domain.PenerimaResponse{
	var penerimaResponse domain.PenerimaResponse
	rows, err := r.db.Query("select p.id, p.type, p.qty, p.alamat_pengirim, p.pic_pengirim, p.alamat_penerima, p.warehouse_penerima, p.pic_penerima, p.tanggal_pengiriman, p.tanggal_sampai, p.ido_gd, p.sn_mac_barcode, p.batch, p.ido_gd_time_added, p.sn_time_added, g.regional, p.time_added from penerimaan p JOIN gudang g ON g.lokasi_wh = p.warehouse_penerima where p.id = ?", id)
	helper.PanicIfError(err)

	if rows.Next(){
		err = rows.Scan(&penerimaResponse.ID, &penerimaResponse.Type, &penerimaResponse.Qty, &penerimaResponse.AlamatPengirim, &penerimaResponse.PICPengirim, &penerimaResponse.AlamatPenerima, &penerimaResponse.WarehousePenerima, &penerimaResponse.PICPenerima, &penerimaResponse.TanggalPengiriman, &penerimaResponse.TanggalSampai, &penerimaResponse.IDOGD, &penerimaResponse.SNMacBarcode, &penerimaResponse.Batch, &penerimaResponse.IDOGDTimeAdded, &penerimaResponse.SNTimeAdded, &penerimaResponse.Regional, &penerimaResponse.TimeAdded)
	}

	return penerimaResponse
}

func (r *penerimaRepository) EditIDOGDById(id string, data domain.PenerimaPost){
	// Get current date and time
	now := time.Now()

	// Format the date and time
	timeNow := now.Format("2006-01-02 15:04:05")

	_, err := r.db.Query("UPDATE `penerimaan` SET `ido_gd`= ?, `ido_gd_time_added` = ? WHERE id = ?", data.IDOGD, timeNow, id)
	helper.PanicIfError(err)
}

func (r *penerimaRepository) EditTanggalPenerimaanById(id string, data domain.PenerimaPost){
	_, err := r.db.Query("UPDATE `penerimaan` SET `tanggal_pengiriman`= ?,`tanggal_sampai`= ?, `sn_mac_barcode` = ?, `sn_time_added` = ? WHERE id = ?", data.TanggalPengiriman, data.TanggalSampai, data.SNMacBarcode, data.SNTimeAdded, id)
	helper.PanicIfError(err)
}

func (r *penerimaRepository) EditTanggalOnly(id string, data domain.PenerimaPost){
	_, err := r.db.Query("UPDATE `penerimaan` SET `tanggal_pengiriman`= ?,`tanggal_sampai`= ? WHERE id = ?", data.TanggalPengiriman, data.TanggalSampai, id)
	helper.PanicIfError(err)
}

func (r *penerimaRepository) DeleteAllPenerima(){
	_, err := r.db.Query("DELETE FROM penerimaan WHERE (`type` LIKE '%ONT_ZTE_F670L%') OR (`type` like '%G240WL%' OR `type` like '%2425G%') OR (`type` LIKE '%HG6145D2%' or `type` LIKE '%HG6145F%') OR (`type` LIKE '%HG8145%') OR (`type` like '%HG8145V5%') OR (`type` LIKE '%ZTE_F670 V2.0%' or `type` = 'ONT_ZTE_F670 V2.0' or `type` = 'ONT_ZTE_F670') OR (`type` LIKE '%HG6245N%') OR (`type` like '%HG8245W5-6T%' or `type` = 'ONT_HW_HG8245W5-6T')")
	helper.PanicIfError(err)
}