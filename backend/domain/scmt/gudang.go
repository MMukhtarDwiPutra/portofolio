package domain

type Gudang struct {
	ID    int `bson:"id" json:"id"`
	Regional  string `bson:"regional" json:"regional"`
	Witel string `bson:"witel" json:"witel"`
	LokasiWH string `bson:"lokasi_wh" json:"lokasi_wh"`
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