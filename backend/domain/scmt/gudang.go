package domain

type Gudang struct {
	ID    int `bson:"id" json:"id"`
	Regional  string `bson:"regional" json:"regional"`
	Witel string `bson:"witel" json:"witel"`
	LokasiWH string `bson:"lokasi_wh" json:"lokasi_wh"`
	Lokasi string `bson:"lokasi" json:"lokasi"`
	Wilayah string `bson:"wilayah" json:"wilayah"`
	MinimumQty string `bson:"minimum_qty" json:"minimum_qty"`
	RetailZTE string `bson:"retail_zte" json:"retail_zte"`
	RetailHW string `bson:"retail_hw" json:"retail_hw"`
	RetailFH string `bson:"retail_fh" json:"retail_fh"`
	RetailALU string `bson:"retail_alu" json:"retail_alu"`
	PremiumZTE string `bson:"premium_zte" json:"premium_zte"`
	PremiumFH string `bson:"premium_fh" json:"premium_fh"`
	PremiumHW string `bson:"premium_hw" json:"premium_hw"`
	STBZTE string `bson:"stb_zte" json:"stb_zte"`
	APCisco string `bson:"ap_cisco" json:"ap_cisco"`
	APHuawei string `bson:"ap_huawei" json:"ap_huawei"`
}

type TREGMinimumResponse struct{
	ID    int `bson:"id" json:"id"`
	Regional    string `bson:"regional" json:"regional"`
	MinimumQty    string `bson:"minimum_qty" json:"minimum_qty"`
	LokasiWH    string `bson:"lokasi_wh" json:"lokasi_wh"`
	Witel    string `bson:"witel" json:"witel"`
	Lokasi    string `bson:"lokasi" json:"lokasi"`
	Wilayah    string `bson:"wilayah" json:"wilayah"`
	RetailZTE int `bson:"retail_zte" json:"retail_zte"`
	RetailHW int `bson:"retail_hw" json:"retail_hw"`
	RetailFH int `bson:"retail_fh" json:"retail_fh"`
	RetailALU int `bson:"retail_alu" json:"retail_alu"`
	PremiumZTE int `bson:"premium_zte" json:"premium_zte"`
	PremiumFH int `bson:"premium_fh" json:"premium_fh"`
	PremiumHW int `bson:"premium_hw" json:"premium_hw"`
	BatasAtasRetailZTE int `bson:"batas_atas_retail_zte" json:"batas_atas_retail_zte"`
	BatasAtasRetailHW int `bson:"batas_atas_retail_hw" json:"batas_atas_retail_hw"`
	BatasAtasRetailFH int `bson:"batas_atas_retail_fh" json:"batas_atas_retail_fh"`
	BatasAtasRetailALU int `bson:"batas_atas_retail_alu" json:"batas_atas_retail_alu"`
	BatasAtasPremiumZTE int `bson:"batas_atas_premium_zte" json:"batas_atas_premium_zte"`
	BatasAtasPremiumFH int `bson:"batas_atas_premium_fh" json:"batas_atas_premium_fh"`
	BatasAtasPremiumHW int `bson:"batas_atas_premium_hw" json:"batas_atas_premium_hw"`
	BatasBawahRetailZTE int `bson:"batas_bawah_retail_zte" json:"batas_bawah_retail_zte"`
	BatasBawahRetailHW int `bson:"batas_bawah_retail_hw" json:"batas_bawah_retail_hw"`
	BatasBawahRetailFH int `bson:"batas_bawah_retail_fh" json:"batas_bawah_retail_fh"`
	BatasBawahRetailALU int `bson:"batas_bawah_retail_alu" json:"batas_bawah_retail_alu"`
	BatasBawahPremiumZTE int `bson:"batas_bawah_premium_zte" json:"batas_bawah_premium_zte"`
	BatasBawahPremiumFH int `bson:"batas_bawah_premium_fh" json:"batas_bawah_premium_fh"`
	BatasBawahPremiumHW int `bson:"batas_bawah_premium_hw" json:"batas_bawah_premium_hw"`
	RetailStockZTE int `bson:"retail_stock_zte" json:"retail_stock_zte"`
	RetailStockHuawei int `bson:"retail_stock_hw" json:"retail_stock_hw"`
	RetailStockFiberhome int `bson:"retail_stock_fh" json:"retail_stock_fh"`
	RetailStockNokia int `bson:"retail_stock_alu" json:"retail_stock_alu"`
	PremiumStockZTE int `bson:"premium_stock_zte" json:"premium_stock_zte"`
	PremiumStockFiberhome int `bson:"premium_stock_fh" json:"premium_stock_fh"`
	PremiumStockHuawei int `bson:"premium_stock_hw" json:"premium_stock_hw"`
	QtyKirimRetailZTE int `bson:"qty_kirim_retail_zte" json:"qty_kirim_retail_zte"`
	QtyKirimRetailHW int `bson:"qty_kirim_retail_hw" json:"qty_kirim_retail_hw"`
	QtyKirimRetailFH int `bson:"qty_kirim_retail_fh" json:"qty_kirim_retail_fh"`
	QtyKirimRetailALU int `bson:"qty_kirim_retail_alu" json:"qty_kirim_retail_alu"`
	QtyKirimPremiumZTE int `bson:"qty_kirim_premium_zte" json:"qty_kirim_premium_zte"`
	QtyKirimPremiumFH int `bson:"qty_kirim_premium_fh" json:"qty_kirim_premium_fh"`
	QtyKirimPremiumHW int `bson:"qty_kirim_premium_hw" json:"qty_kirim_premium_hw"`
	OnDeliveryRetailZTE int `bson:"on_delivery_retail_zte" json:"on_delivery_retail_zte"`
	OnDeliveryRetailHuawei int `bson:"on_delivery_retail_hw" json:"on_delivery_retail_hw"`
	OnDeliveryRetailFiberhome int `bson:"on_delivery_retail_fh" json:"on_delivery_retail_fh"`
	OnDeliveryRetailNokia int `bson:"on_delivery_retail_alu" json:"on_delivery_retail_alu"`
	OnDeliveryPremiumZTE int `bson:"on_delivery_premium_zte" json:"on_delivery_premium_zte"`
	OnDeliveryPremiumFiberhome int `bson:"on_delivery_premium_fh" json:"on_delivery_premium_fh"`
	OnDeliveryPremiumHuawei int `bson:"on_delivery_premium_hw" json:"on_delivery_premium_hw"`
	BlinkRetailZTE int `bson:"blink_retail_zte" json:"blink_retail_zte"`
	BlinkRetailHW int `bson:"blink_retail_hw" json:"blink_retail_hw"`
	BlinkRetailFH int `bson:"blink_retail_fh" json:"blink_retail_fh"`
	BlinkRetailALU int `bson:"blink_retail_alu" json:"blink_retail_alu"`
	BlinkPremiumZTE int `bson:"blink_premium_zte" json:"blink_premium_zte"`
	BlinkPremiumFH int `bson:"blink_premium_fh" json:"blink_premium_fh"`
	BlinkPremiumHW int `bson:"blink_premium_hw" json:"blink_premium_hw"`
	TotalRetailStock int `bson:"total_retail_stock" json:"total_retail_stock"`
	TotalRetail int `bson:"total_retail" json:"total_retail"`
	OnDeliveryTotalRetail int `bson:"on_delivery_total_retail" json:"on_delivery_total_retail"`
	TotalPremiumStock int `bson:"total_premium_stock" json:"total_premium_stock"`
	TotalPremium int `bson:"total_premium" json:"total_premium"`
	OnDeliveryTotalPremium int `bson:"on_delivery_total_premium" json:"on_delivery_total_premium"`
}

type LokasiWarehouseResponse struct{
	LokasiWH string `bson:"lokasi_wh" json:"lokasi_wh"`
	Lokasi string `bson:"lokasi" json:"lokasi"`
}

type QtyMinimumResponse struct{
	ID    int `bson:"id" json:"id"`
	Regional    int `bson:"regional" json:"regional"`
	Witel    int `bson:"witel" json:"witel"`
	LokasiWH    int `bson:"lokasi_wh" json:"lokasi_wh"`
	Lokasi    int `bson:"lokasi" json:"lokasi"`
	Wilayah    int `bson:"wilayah" json:"wilayah"`
	MinimumQty    int `bson:"minimum_qty" json:"minimum_qty"`
	RetailZTE string `bson:"retail_zte" json:"retail_zte"`
	RetailHW string `bson:"retail_hw" json:"retail_hw"`
	RetailFH string `bson:"retail_fh" json:"retail_fh"`
	RetailALU string `bson:"retail_alu" json:"retail_alu"`
	PremiumZTE string `bson:"premium_zte" json:"premium_zte"`
	PremiumFH string `bson:"premium_fh" json:"premium_fh"`
	PremiumHW string `bson:"premium_hw" json:"premium_hw"`
	BatasAtasRetailZTE string `bson:"batas_atas_retail_zte" json:"batas_atas_retail_zte"`
	BatasAtasRetailHW string `bson:"batas_atas_retail_hw" json:"batas_atas_retail_hw"`
	BatasAtasRetailFH string `bson:"batas_atas_retail_fh" json:"batas_atas_retail_fh"`
	BatasAtasRetailALU string `bson:"batas_atas_retail_alu" json:"batas_atas_retail_alu"`
	BatasAtasPremiumZTE string `bson:"batas_atas_premium_zte" json:"batas_atas_premium_zte"`
	BatasAtasPremiumFH string `bson:"batas_atas_premium_fh" json:"batas_atas_premium_fh"`
	BatasAtasPremiumHW string `bson:"batas_atas_premium_hw" json:"batas_atas_premium_hw"`
	BatasBawahRetailZTE string `bson:"batas_bawah_retail_zte" json:"batas_bawah_retail_zte"`
	BatasBawahRetailHW string `bson:"batas_bawah_retail_hw" json:"batas_bawah_retail_hw"`
	BatasBawahRetailFH string `bson:"batas_bawah_retail_fh" json:"batas_bawah_retail_fh"`
	BatasBawahRetailALU string `bson:"batas_bawah_retail_alu" json:"batas_bawah_retail_alu"`
	BatasBawahPremiumZTE string `bson:"batas_bawah_premium_zte" json:"batas_bawah_premium_zte"`
	BatasBawahPremiumFH string `bson:"batas_bawah_premium_fh" json:"batas_bawah_premium_fh"`
	BatasBawahPremiumHW string `bson:"batas_bawah_premium_hw" json:"batas_bawah_premium_hw"`
}