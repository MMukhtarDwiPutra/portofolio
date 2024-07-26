package service

import(
	"portofolio.com/repository/scmt"
	"portofolio.com/domain/scmt"
	"fmt"
)

type DataTmpService interface{
	GetAllDataTmp() []domain.DataTmp
	InsertDataTmp(data domain.DataTmp)
	GetLastDataTmp() domain.DataTmp
	CountRetailPerWitel(merk string) []domain.CountResponse
	CountPremiumPerWitel(merk string) []domain.CountResponse
	CountSTBPerWitel(merk string) []domain.CountResponse
	CountAPPerWitel(merk string) []domain.CountResponse
	CountDataPerWitel()
}

type dataTmpService struct{
	dataTmpRepository repository.DataTmpRepository
	gudangRepository repository.GudangRepository
}

func NewDataTmpService(dataTmpRepository repository.DataTmpRepository, gudangRepository repository.GudangRepository) *dataTmpService{
	return &dataTmpService{dataTmpRepository, gudangRepository}
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

func (s *dataTmpService) CountDataPerWitel(){
	data := make(map[string]interface{})

    // Get the slices
    tregQtyMinimum := s.gudangRepository.GetTREGQtyMinimum()
    qtyMinimum := s.gudangRepository.GetQtyMinimum()

	data["witel"] = append(qtyMinimum, tregQtyMinimum...)

	fmt.Println(data)

	// merk := "fiberhome";
	// data[merk] = s.dataTmpRepository.CountRetailPerWitel(merk);
	// data = s.AddStockCountTmp(jenis_stock, data, merk);
}

// func AddStockCount(jenis_stock string, data map[string]interface{}, merk string){
// 	arrayStock := data[merk]
// 	witel := s.dataTmpRepository.GetWitelsFromDataByMerk(merk)

// 	for i := 0; i < len(witel); i++ {
//         nama_witel := data["witel"].LokasiWH
//     }
// }

// public function addStockCount($jenis_stock, $data, $merk){
// 	$array_stock = $data[$merk];
// 	$witel = Data::getWitelsFromDataByMerk($merk);
	
// 	for ($i=0; $i < count($data["witel"]); $i++) {
// 		$nama_witel = $data["witel"][$i]->lokasi_wh;

// 		$filtered_array = array_filter($array_stock, function ($obj) use ($nama_witel){
// 		  return str_contains(strtolower($nama_witel), strtolower($obj->witel));
// 		});

// 		if(count($filtered_array) == 0){
// 			$data["witel"][$i]->{$jenis_stock."_stock_".$merk} = 0;
// 		}else{
// 			$tmp_stock = current($filtered_array)->stock;

// 			$data["witel"][$i]->{$jenis_stock."_stock_".$merk} = $tmp_stock;
// 		}
// 	}

// 	return $data;
// }
