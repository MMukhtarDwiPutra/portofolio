package domain

import (
)

const (
	CollectionData = "data_scmt"
)

type Data struct {
	ID    int `bson:"_id" json:"-"`
	StatusPerangkat  string `bson:"status_perangkat" json:"status_perangkat"`
	Witel string `bson:"witel" json:"witel"`
	TREGWH string `bson:"treg_wh" json:"treg_wh"`
	Tipe string `bson:"tipe" json:"tipe"`
	Merk string `bson:"merk" json:"merk"`
	DeviceID string `bson:"device_id" json:"device_id"`
	ItemCode string `bson:"item_code" json:"item_code"`
}

type DataTmp struct {
	ID    int `bson:"_id" json:"-"`
	Region  string `bson:"region" json:"region"`
	LokasiWH string `bson:"lokasi_wh" json:"lokasi_wh"`
	Status string `bson:"status" json:"status"`
	Jumlah string `bson:"jumlah" json:"jumlah"`
	Deskripsi string `bson:"deskripsi" json:"deskripsi"`
}

type CountResponse struct{
	LokasiWH string `bson:"lokasi_wh" json:"lokasi_wh"`
	Stock int `bson:"stock" json:"stock"`
}