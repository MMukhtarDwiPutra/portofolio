package service

import(
	"time"
	"portofolio.com/repository/scmt"
	// "portofolio.com/domain/scmt"
	"portofolio.com/api/helper"
	"fmt"
)

type PenerimaService interface{
	GetPengirimanONT() map[string]interface{}
}

type penerimaService struct{
	penerimaRepository repository.PenerimaRepository
	gudangRepository repository.GudangRepository
	fiturRepository repository.FiturRepository
}

func NewPenerimaService(penerimaRepository repository.PenerimaRepository, gudangRepository repository.GudangRepository, fiturRepository repository.FiturRepository) *penerimaService{
	return &penerimaService{penerimaRepository, gudangRepository, fiturRepository}
}

func (s *penerimaService) GetPengirimanONT() map[string]interface{}{
	data := make(map[string]interface{})

	data["warehouse"] = s.gudangRepository.GetAllWarehouse();
	data["penerima"] = s.penerimaRepository.GetAllDataONT();

	namaFitur := "report_delivery_ont"
	data["status_filling_disable"] = s.fiturRepository.GetFitur(namaFitur)

	waktuUpdate, waktuDibuat := s.penerimaRepository.GetTableLastUpdate()
	
	layout := "2006-01-02 15:04:05"

	if(waktuUpdate != ""){
		waktuDibuatTimeFormat, err := time.Parse(layout, waktuDibuat)
		helper.PanicIfError(err)

		waktuUpdateTimeFormat, err := time.Parse(layout, waktuUpdate)
		helper.PanicIfError(err)

		data["last_update"] = waktuUpdate
		if(waktuDibuatTimeFormat.After(waktuUpdateTimeFormat)){
			data["last_update"] = waktuDibuat
		}
	}else{
		fmt.Println(waktuDibuat)
		data["last_update"] = waktuDibuat
	}

	return data
}