package domain

type PenerimaResponse struct{
	ID    int `bson:"id" json:"id"`
	Type  string `bson:"type" json:"type"`
	Qty  int `bson:"qty" json:"qty"`
	AlamatPengirim  string `bson:"alamat_pengirim" json:"alamat_pengirim"`
	PICPengirim  string `bson:"pic_pengirim" json:"pic_pengirim"`
	AlamatPenerima  string `bson:"alamat_penerima" json:"alamat_penerima"`
	WarehousePenerima  string `bson:"warehouse_penerima" json:"warehouse_penerima"`
	PICPenerima  string `bson:"pic_penerima" json:"pic_penerima"`
	TanggalPengiriman  string `bson:"tanggal_pengiriman" json:"tanggal_pengiriman"`
	TanggalSampai  string `bson:"tanggal_sampai" json:"tanggal_sampai"`
	IDOGD  string `bson:"ido_gd" json:"ido_gd"`
	SNMacBarcode  string `bson:"sn_mac_barcode" json:"sn_mac_barcode"`
	Batch  string `bson:"batch" json:"batch"`
	IDOGDTimeAdded  string `bson:"ido_gd_time_added" json:"ido_gd_time_added"`
	TimeAdded  string `bson:"time_added" json:"time_added"`
	SNTimeAdded  string `bson:"sn_time_added" json:"sn_time_added"`
	Regional string `bson:"regional" json:"regional"`
}