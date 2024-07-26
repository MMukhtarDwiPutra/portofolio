package service

import(
	"portofolio.com/repository/scmt"
	"portofolio.com/domain/scmt"
)

type DataTmpService interface{
	GetAllDataTmp() []domain.DataTmp
	InsertDataTmp(data domain.DataTmp)
	GetLastDataTmp() domain.DataTmp
	CountRetailPerWitel(merk string) []domain.CountResponse
	CountPremiumPerWitel(merk string) []domain.CountResponse
	CountSTBPerWitel(merk string) []domain.CountResponse
	CountAPPerWitel(merk string) []domain.CountResponse
}

type dataTmpService struct{
	dataTmpRepository repository.DataTmpRepository
}

func NewDataTmpService(dataTmpRepository repository.DataTmpRepository) *dataTmpService{
	return &dataTmpService{dataTmpRepository}
}

func (s *dataTmpService) GetAllDataTmp() []domain.DataTmp{
	return s.dataTmpRepository.GetAllDataTmp()
}

func (s *dataTmpService) InsertDataTmp(data domain.DataTmp){
	s.dataTmpRepository.InsertData(data)
}

func (s *dataTmpService) GetLastDataTmp() domain.DataTmp{
	return s.dataTmpRepository.GetLastDataTmp()
}

func (s *dataTmpService) CountRetailPerWitel(merk string) []domain.CountResponse{
	return s.dataTmpRepository.CountRetailPerWitel(merk)
}

func (s *dataTmpService) CountPremiumPerWitel(merk string) []domain.CountResponse{
	return s.dataTmpRepository.CountPremiumPerWitel(merk)
}

func (s *dataTmpService) CountSTBPerWitel(merk string) []domain.CountResponse{
	return s.dataTmpRepository.CountSTBPerWitel(merk)
}

func (s *dataTmpService) CountAPPerWitel(merk string) []domain.CountResponse{
	return s.dataTmpRepository.CountAPPerWitel(merk)
}