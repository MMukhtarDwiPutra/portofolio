package service

import(
	"portofolio.com/repository/scmt"
	"portofolio.com/domain/scmt"
	"portofolio.com/api/helper"
	"strings"
	"fmt"
	"reflect"
	"github.com/xuri/excelize/v2"
	"time"
	"strconv"
)

type DataTmpService interface{
	GetAllDataTmp() []domain.DataTmp
	InsertDataTmp(data domain.DataTmp)
	GetLastDataTmp() domain.DataTmp
	CountRetailPerWitel(merk string) []domain.CountResponse
	CountPremiumPerWitel(merk string) []domain.CountResponse
	CountSTBPerWitel(merk string) []domain.CountResponse
	CountAPPerWitel(merk string) []domain.CountResponse
	CountDataPerWitelTmp() map[string]interface{}
	AddStockCountTmp(jenisStock string, data map[string]interface{}, merk string) map[string]interface{}
	HitungQtyKirim(data map[string]interface{}) map[string]interface{}
	RekapDeliveryTREG() map[string]interface{}
	RekapDeliveryWitel(lokasiWH string) map[string]interface{}
	GetExportDataTmp() ([]byte, string, error)
	GetExportMinimumStockDatabase() ([]byte, string, error)
	DownloadTemplateMinimumStock() ([]byte, string, error)
	DownloadTemplateDataTmp() ([]byte, string, error)
	UploadDataTmp()
	ExportDataTmp(jenisWarna string, jenisExport string) ([]byte, string, error)
}

type dataTmpService struct{
	dataTmpRepository repository.DataTmpRepository
	gudangRepository repository.GudangRepository
	penerimaRepository repository.PenerimaRepository
}

func NewDataTmpService(dataTmpRepository repository.DataTmpRepository, gudangRepository repository.GudangRepository, penerimaRepository repository.PenerimaRepository) *dataTmpService{
	return &dataTmpService{dataTmpRepository, gudangRepository, penerimaRepository}
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

func (s *dataTmpService) CountDataPerWitelTmp() map[string]interface{}{
	data := make(map[string]interface{})

    // Get the slices
    tregQtyMinimum := s.gudangRepository.GetTREGQtyMinimum()
    qtyMinimum := s.gudangRepository.GetQtyMinimum()

	data["witel"] = append(qtyMinimum, tregQtyMinimum...)

	var merk, jenisStock string

	jenisStock = "Retail"
	merk = "Fiberhome"

	data[merk] = s.dataTmpRepository.CountRetailPerWitel(merk);
	data["penerima_"+strings.ToLower(merk)] = s.penerimaRepository.CountRetailPerWitel(merk);
	data = s.AddStockCountTmp(jenisStock, data, merk);

	merk = "ZTE"

	data[merk] = s.dataTmpRepository.CountRetailPerWitel(merk);
	data["penerima_"+strings.ToLower(merk)] = s.penerimaRepository.CountRetailPerWitel(merk);
	data = s.AddStockCountTmp(jenisStock, data, merk);

	merk = "Nokia"

	data[merk] = s.dataTmpRepository.CountRetailPerWitel(merk);
	data["penerima_"+strings.ToLower(merk)] = s.penerimaRepository.CountRetailPerWitel(merk);
	data = s.AddStockCountTmp(jenisStock, data, merk);

	merk = "Huawei"

	data[merk] = s.dataTmpRepository.CountRetailPerWitel(merk);
	data["penerima_"+strings.ToLower(merk)] = s.penerimaRepository.CountRetailPerWitel(merk);
	data = s.AddStockCountTmp(jenisStock, data, merk);

	data = AddTotalStockCount(jenisStock, data)

	jenisStock = "Premium"
	merk = "Fiberhome"

	data[merk] = s.dataTmpRepository.CountPremiumPerWitel(merk);
	data["penerima_"+strings.ToLower(merk)] = s.penerimaRepository.CountPremiumPerWitel(merk);
	data = s.AddStockCountTmp(jenisStock, data, merk);

	merk = "ZTE"

	data[merk] = s.dataTmpRepository.CountPremiumPerWitel(merk);
	data["penerima_"+strings.ToLower(merk)] = s.penerimaRepository.CountPremiumPerWitel(merk);
	data = s.AddStockCountTmp(jenisStock, data, merk);

	merk = "Huawei"

	data[merk] = s.dataTmpRepository.CountPremiumPerWitel(merk);
	data["penerima_"+strings.ToLower(merk)] = s.penerimaRepository.CountPremiumPerWitel(merk);
	data = s.AddStockCountTmp(jenisStock, data, merk);

	data = AddTotalStockCount(jenisStock, data)
	data = AddStockTregCountTmp(data)
	// if witelSlice, ok := data["witel"].([]domain.TREGMinimumResponse); ok {
	// 	return witelSlice
	// }

	// var witelSlice []domain.TREGMinimumResponse
	return data
}

func AddTotalStockCount(jenisStock string, data map[string]interface{}) map[string]interface{}{
	if witelSlice, ok := data["witel"].([]domain.TREGMinimumResponse); ok {
		if(jenisStock == "Retail"){		
			for i := range witelSlice {
				witelSlice[i].TotalRetailStock = witelSlice[i].RetailStockHuawei + witelSlice[i].RetailStockFiberhome + witelSlice[i].RetailStockNokia + witelSlice[i].RetailStockZTE
				witelSlice[i].TotalRetail = witelSlice[i].RetailHW + witelSlice[i].RetailFH + witelSlice[i].RetailALU + witelSlice[i].RetailZTE
				witelSlice[i].OnDeliveryTotalRetail = witelSlice[i].OnDeliveryRetailHuawei + witelSlice[i].OnDeliveryRetailFiberhome + witelSlice[i].OnDeliveryRetailNokia + witelSlice[i].OnDeliveryRetailZTE
			}
		}else if(jenisStock == "Premium"){
			for i := range witelSlice {
				witelSlice[i].TotalPremiumStock = witelSlice[i].PremiumStockHuawei + witelSlice[i].PremiumStockFiberhome + + witelSlice[i].PremiumStockZTE
				witelSlice[i].TotalPremium = witelSlice[i].PremiumHW + witelSlice[i].PremiumFH + witelSlice[i].PremiumZTE
				witelSlice[i].OnDeliveryTotalPremium = witelSlice[i].OnDeliveryPremiumHuawei + witelSlice[i].OnDeliveryPremiumFiberhome + witelSlice[i].OnDeliveryPremiumZTE
			}
		}
		data["witel"] = witelSlice
	}

	return data
}

// Generic sum function for summing a field in a slice of structs
func SumWitelArrayByField[T any](items []T, selector func(T) int) int {
    total := 0
    for _, item := range items {
        total += selector(item)
    }
    return total
}

func AddStockTregCountTmp(data map[string]interface{}) map[string]interface{}{
	if witelSlice, ok := data["witel"].([]domain.TREGMinimumResponse); ok{
		tregArray := filterDataTREGBySO(witelSlice, "WH TR TREG")

		for i:= 0; i < len(tregArray); i++{
			regional := tregArray[i].Regional

			whArray := filterDataWitelByTREG(witelSlice, regional)

			var idx int
			for j, dataWitel := range witelSlice{
				if(strings.Contains(dataWitel.LokasiWH, regional)){
					idx = j;
					break
				}
			}

			witelSlice[idx].RetailStockZTE = SumWitelArrayByField(whArray, func(i domain.TREGMinimumResponse) int { return i.RetailStockZTE })
			witelSlice[idx].RetailStockFiberhome = SumWitelArrayByField(whArray, func(i domain.TREGMinimumResponse) int { return i.RetailStockFiberhome })
			witelSlice[idx].RetailStockHuawei = SumWitelArrayByField(whArray, func(i domain.TREGMinimumResponse) int { return i.RetailStockHuawei })
			witelSlice[idx].RetailStockNokia = SumWitelArrayByField(whArray, func(i domain.TREGMinimumResponse) int { return i.RetailStockNokia })

			witelSlice[idx].RetailZTE = SumWitelArrayByField(whArray, func(i domain.TREGMinimumResponse) int { return i.RetailZTE })
			witelSlice[idx].RetailFH = SumWitelArrayByField(whArray, func(i domain.TREGMinimumResponse) int { return i.RetailFH })
			witelSlice[idx].RetailHW = SumWitelArrayByField(whArray, func(i domain.TREGMinimumResponse) int { return i.RetailHW })
			witelSlice[idx].RetailALU = SumWitelArrayByField(whArray, func(i domain.TREGMinimumResponse) int { return i.RetailALU })

			witelSlice[idx].OnDeliveryRetailZTE = SumWitelArrayByField(whArray, func(i domain.TREGMinimumResponse) int { return i.OnDeliveryRetailZTE })
			witelSlice[idx].OnDeliveryRetailFiberhome = SumWitelArrayByField(whArray, func(i domain.TREGMinimumResponse) int { return i.OnDeliveryRetailFiberhome })
			witelSlice[idx].OnDeliveryRetailHuawei = SumWitelArrayByField(whArray, func(i domain.TREGMinimumResponse) int { return i.OnDeliveryRetailHuawei })
			witelSlice[idx].OnDeliveryRetailNokia = SumWitelArrayByField(whArray, func(i domain.TREGMinimumResponse) int { return i.OnDeliveryRetailNokia })

			witelSlice[idx].TotalRetailStock = witelSlice[idx].RetailStockHuawei + witelSlice[idx].RetailStockFiberhome + witelSlice[idx].RetailStockNokia + witelSlice[idx].RetailStockZTE
			witelSlice[idx].TotalRetail = witelSlice[idx].RetailHW + witelSlice[idx].RetailFH + witelSlice[idx].RetailALU + witelSlice[idx].RetailZTE
			witelSlice[idx].OnDeliveryTotalRetail = witelSlice[idx].OnDeliveryRetailHuawei + witelSlice[idx].OnDeliveryRetailFiberhome + witelSlice[idx].OnDeliveryRetailNokia + witelSlice[idx].OnDeliveryRetailZTE

			witelSlice[idx].PremiumZTE = SumWitelArrayByField(whArray, func(i domain.TREGMinimumResponse) int { return i.PremiumZTE })
			witelSlice[idx].PremiumFH = SumWitelArrayByField(whArray, func(i domain.TREGMinimumResponse) int { return i.PremiumFH })
			witelSlice[idx].PremiumHW = SumWitelArrayByField(whArray, func(i domain.TREGMinimumResponse) int { return i.PremiumHW })

			witelSlice[idx].PremiumStockZTE = SumWitelArrayByField(whArray, func(i domain.TREGMinimumResponse) int { return i.PremiumStockZTE })
			witelSlice[idx].PremiumStockFiberhome = SumWitelArrayByField(whArray, func(i domain.TREGMinimumResponse) int { return i.PremiumStockFiberhome })
			witelSlice[idx].PremiumStockHuawei = SumWitelArrayByField(whArray, func(i domain.TREGMinimumResponse) int { return i.PremiumStockHuawei })

			witelSlice[idx].OnDeliveryPremiumZTE = SumWitelArrayByField(whArray, func(i domain.TREGMinimumResponse) int { return i.OnDeliveryPremiumZTE })
			witelSlice[idx].OnDeliveryPremiumFiberhome = SumWitelArrayByField(whArray, func(i domain.TREGMinimumResponse) int { return i.OnDeliveryPremiumFiberhome })
			witelSlice[idx].OnDeliveryPremiumHuawei = SumWitelArrayByField(whArray, func(i domain.TREGMinimumResponse) int { return i.OnDeliveryPremiumHuawei })
			
			witelSlice[idx].TotalPremiumStock = witelSlice[idx].PremiumStockHuawei + witelSlice[idx].PremiumStockFiberhome + witelSlice[idx].PremiumStockZTE
			witelSlice[idx].TotalPremium = witelSlice[idx].PremiumHW + witelSlice[idx].PremiumFH + witelSlice[idx].PremiumZTE
			witelSlice[idx].OnDeliveryTotalPremium = witelSlice[idx].OnDeliveryPremiumHuawei + witelSlice[idx].OnDeliveryPremiumFiberhome + witelSlice[idx].OnDeliveryPremiumZTE
		}
	}
	return data
}

func filterDataTREGBySO(arrayDataWitel []domain.TREGMinimumResponse, namaWitel string) []domain.TREGMinimumResponse {
    var dataTREGArray []domain.TREGMinimumResponse
    for _, dataWitel := range arrayDataWitel {
        if strings.Contains(dataWitel.LokasiWH, namaWitel){
            dataTREGArray = append(dataTREGArray, dataWitel)
        }
    }
    return dataTREGArray
}

func filterArrayStockByWitel(arrayStocks []domain.CountResponse, namaWitel string) domain.CountResponse {
    var stockObj domain.CountResponse
    for _, stock := range arrayStocks {
        if strings.Contains(stock.LokasiWH, namaWitel){
            return stock
        }
    }
    return stockObj
}

func filterArrayStockTASOByWitel(arrayStocks []domain.CountResponse, namaWitel string) domain.CountResponse{
	var stockObj domain.CountResponse
    for _, stock := range arrayStocks {
        if strings.Contains(stock.LokasiWH, namaWitel) && strings.Contains(stock.LokasiWH, "SO"){
            return stock
        }
    }
    return stockObj
}

func filterArrayStockTASOByLokasiWH(arrayStocks []domain.CountResponse, namaWitel string) domain.CountResponse{
	var stockObj domain.CountResponse
    for _, stock := range arrayStocks {
        if stock.LokasiWH == namaWitel{
            return stock
        }
    }
    return stockObj
}

func filterArrayStockWitelByWitel(arrayStocks []domain.CountResponse, namaWitel string) domain.CountResponse {
    var stockObj domain.CountResponse
    for _, stock := range arrayStocks {
        if strings.Contains(stock.LokasiWH, namaWitel) && strings.Contains(stock.LokasiWH, "WITEL"){
            return stock
        }
    }
    return stockObj
}

func filterArrayPenerimaByWitel(arrayPenerima []domain.PenerimaResponse, namaWitel string) []domain.PenerimaResponse {
    var penerimaObj []domain.PenerimaResponse
    for _, penerima := range arrayPenerima {
        if strings.Contains(penerima.WarehousePenerima, namaWitel){
            penerimaObj = append(penerimaObj, penerima)
        }
    }
    return penerimaObj
}

func filterDataWitelByWitel(arrayDataWitel []domain.TREGMinimumResponse, namaWitel string) []domain.TREGMinimumResponse {
    var dataWitelObj []domain.TREGMinimumResponse
    for _, dataWitel := range arrayDataWitel {
        if strings.Contains(dataWitel.Witel, namaWitel){
            dataWitelObj = append(dataWitelObj, dataWitel)
        }
    }
    return dataWitelObj
}

func filterDataWitelByWitelWithoutWitelSO(arrayDataWitel []domain.TREGMinimumResponse, namaWitel string) []domain.TREGMinimumResponse {
    var dataWitelObj []domain.TREGMinimumResponse
    for _, dataWitel := range arrayDataWitel {
        if strings.Contains(dataWitel.Witel, namaWitel) && !strings.Contains(dataWitel.LokasiWH, namaWitel){
            dataWitelObj = append(dataWitelObj, dataWitel)
        }
    }
    return dataWitelObj
}

func filterDataTREGMinimumResponseByTREG(arrayDataWitel []domain.TREGMinimumResponse) []domain.TREGMinimumResponse {
    var dataTREGObj []domain.TREGMinimumResponse
    for _, dataWitel := range arrayDataWitel {
        if strings.Contains(dataWitel.LokasiWH, "WH TR TREG"){
            dataTREGObj = append(dataTREGObj, dataWitel)
        }
    }
    return dataTREGObj
}

func filterDataWitelByTREG(arrayDataWitel []domain.TREGMinimumResponse, namaWitel string) []domain.TREGMinimumResponse {
    var dataWitelObj []domain.TREGMinimumResponse
    for _, dataWitel := range arrayDataWitel {
        if strings.Contains(dataWitel.Regional, namaWitel) && !strings.Contains(dataWitel.LokasiWH, "TREG"){
            dataWitelObj = append(dataWitelObj, dataWitel)
        }
    }
    return dataWitelObj
}

func filterDataWitelByWitelTREG(arrayDataWitel []domain.TREGMinimumResponse, namaWitel string) []domain.TREGMinimumResponse {
    var dataWitelObj []domain.TREGMinimumResponse
    for _, dataWitel := range arrayDataWitel {
        if strings.Contains(dataWitel.Regional, namaWitel) && strings.Contains(dataWitel.LokasiWH, "WITEL"){
            dataWitelObj = append(dataWitelObj, dataWitel)
        }
    }
    return dataWitelObj
}

func filterDataWitelBySO(arrayDataWitel []domain.TREGMinimumResponse, namaWitel string) domain.TREGMinimumResponse {
    var dataWitelObj domain.TREGMinimumResponse
    for _, dataWitel := range arrayDataWitel {
        if strings.Contains(dataWitel.LokasiWH, namaWitel){
            return dataWitel
        }
    }
    return dataWitelObj
}

func (s *dataTmpService) AddStockCountTmp(jenisStock string, data map[string]interface{}, merk string) map[string]interface{}{
	arrayStock := data[merk]
	arrayPenerima := data["penerima_"+strings.ToLower(merk)]
	// witel := s.dataTmpRepository.GetWitelsFromDataByMerk(merk)

	if stockSlice, ok := arrayStock.([]domain.CountResponse); ok {
		if witelSlice, ok := data["witel"].([]domain.TREGMinimumResponse); ok {
			if penerimaSlice, ok := arrayPenerima.([]domain.CountResponse); ok {
		        for i := range witelSlice {
		        	objWitel := &witelSlice[i]
					namaWitel := objWitel.LokasiWH

					filteredArray := filterArrayStockTASOByWitel(stockSlice, namaWitel)
					filteredArrayPenerima := filterArrayStockTASOByWitel(penerimaSlice, namaWitel)

					if strings.Contains(namaWitel, "TA SO CCAN"){
						startPos := strings.Index(namaWitel, "CCAN")
					    startPos += 5
						namaWitel = namaWitel[startPos:]

						countWord := len(strings.Split(namaWitel, " "))

						if(countWord > 2){
							filteredArray = filterArrayStockTASOByWitel(stockSlice, namaWitel)
							filteredArrayPenerima = filterArrayStockTASOByWitel(penerimaSlice, namaWitel)
						}else{
							namaWitel = witelSlice[i].LokasiWH
							filteredArray = filterArrayStockTASOByLokasiWH(stockSlice, namaWitel)
							filteredArrayPenerima = filterArrayStockTASOByLokasiWH(penerimaSlice, namaWitel)
							
						}
					}else if strings.Contains(namaWitel, "TA SO"){
						startPos := strings.Index(namaWitel, "TA SO")
					    startPos += 5
						namaWitel = namaWitel[startPos:]

						filteredArray = filterArrayStockTASOByWitel(stockSlice, namaWitel)
						filteredArrayPenerima = filterArrayStockTASOByWitel(penerimaSlice, namaWitel)
					}else if strings.Contains(namaWitel, "WITEL CCAN"){
						startPos := strings.Index(namaWitel, "TA SO")
					    startPos += 11
						namaWitel = namaWitel[startPos:]

						filteredArray = filterArrayStockWitelByWitel(stockSlice, namaWitel)
						filteredArrayPenerima = filterArrayStockWitelByWitel(penerimaSlice, namaWitel)
					}else if strings.Contains(namaWitel, "WITEL"){
						startPos := strings.Index(namaWitel, "WITEL")
					    startPos += 6
						namaWitel = namaWitel[startPos:]

						filteredArray = filterArrayStockWitelByWitel(stockSlice, namaWitel)
						filteredArrayPenerima = filterArrayStockWitelByWitel(penerimaSlice, namaWitel)
					}

		    		if(filteredArray.LokasiWH != ""){
						fieldName := jenisStock + "Stock" + merk

						stock := filteredArray.Stock
						err := setField(objWitel, fieldName, stock)

						helper.PanicIfError(err)
					}

					if (filteredArrayPenerima.LokasiWH != ""){
						fieldName := "OnDelivery"+ jenisStock + merk
						stock := filteredArrayPenerima.Stock
						err := setField(objWitel, fieldName, stock)

						helper.PanicIfError(err)
					}
		        }
		    }else{
	    		fmt.Println("data['penerima'] is not of type PenerimaResponse")
		    }
	    }else{
	    	fmt.Println("data['witel'] is not of type TREGMinimumResponse")
		}
    } else {
        fmt.Println("arrayStock is not of type CountResponse")
    }

    return data
}

func (s *dataTmpService) RekapDeliveryWitel(lokasiWH string) map[string]interface{}{
	data := s.CountDataPerWitelTmp()
	var result int
	// var sumAllZTE, sumAllFH, sumAllHW, sumAllALU, sumAllPZTE, sumAllPFH, sumAllPHW int
	// var sumZTE, sumFH, sumHW, sumALU, sumPZTE, sumPFH, sumPHW int
	if tregSlice, ok := data["witel"].([]domain.TREGMinimumResponse); ok{
		if(strings.Contains(lokasiWH, "TREG")){
			data["response"] = filterDataWitelByWitelTREG(tregSlice, lokasiWH)

			if witelSlice, ok := data["response"].([]domain.TREGMinimumResponse); ok{
				for i := range witelSlice{
					lokasiWH := witelSlice[i].LokasiWH

					dataSO := filterDataWitelByWitelWithoutWitelSO(tregSlice, lokasiWH)

					if(witelSlice[i].RetailStockZTE < witelSlice[i].BatasBawahRetailZTE){
						result = witelSlice[i].BatasAtasRetailZTE - witelSlice[i].RetailStockZTE
						witelSlice[i].QtyKirimRetailZTE += roundToNearest(result, 10)
					}

					if(witelSlice[i].RetailStockHuawei < witelSlice[i].BatasBawahRetailHW){
						result = witelSlice[i].BatasAtasRetailHW - witelSlice[i].RetailStockHuawei
						witelSlice[i].QtyKirimRetailHW += roundToNearest(result, 10)
					}

					if(witelSlice[i].RetailStockFiberhome < witelSlice[i].BatasBawahRetailFH){
						result = witelSlice[i].BatasAtasRetailFH - witelSlice[i].RetailStockFiberhome
						witelSlice[i].QtyKirimRetailFH += roundToNearest(result, 10)
					}

					if(witelSlice[i].RetailStockNokia < witelSlice[i].BatasBawahRetailALU){
						result = witelSlice[i].BatasAtasRetailALU - witelSlice[i].RetailStockNokia
						witelSlice[i].QtyKirimRetailALU += roundToNearest(result, 6)
					}

					if(witelSlice[i].PremiumStockZTE < witelSlice[i].BatasBawahPremiumZTE){
						result = witelSlice[i].BatasAtasPremiumZTE - witelSlice[i].PremiumStockZTE
						witelSlice[i].QtyKirimPremiumZTE += roundToNearest(result, 10)
					}

					if(witelSlice[i].PremiumStockHuawei < witelSlice[i].BatasBawahPremiumHW){
						result = witelSlice[i].BatasAtasPremiumHW - witelSlice[i].PremiumStockHuawei
						witelSlice[i].QtyKirimPremiumHW += roundToNearest(result, 12)
					}

					if(witelSlice[i].PremiumStockFiberhome < witelSlice[i].BatasBawahPremiumFH){
						result = witelSlice[i].BatasAtasPremiumFH - witelSlice[i].PremiumStockFiberhome
						witelSlice[i].QtyKirimPremiumFH += roundToNearest(result, 8)
					}

					for j := range dataSO {
						if(dataSO[j].RetailStockZTE < dataSO[j].BatasBawahRetailZTE){
							result = dataSO[j].BatasAtasRetailZTE - dataSO[j].RetailStockZTE
							witelSlice[i].QtyKirimRetailZTE += roundToNearest(result, 10)
						}

						if(dataSO[j].RetailStockHuawei < dataSO[j].BatasBawahRetailHW){
							result = dataSO[j].BatasAtasRetailHW - dataSO[j].RetailStockHuawei
							witelSlice[i].QtyKirimRetailHW += roundToNearest(result, 10)
						}

						if(dataSO[j].RetailStockFiberhome < dataSO[j].BatasBawahRetailFH){
							result = dataSO[j].BatasAtasRetailFH - dataSO[j].RetailStockFiberhome
							witelSlice[i].QtyKirimRetailFH += roundToNearest(result, 10)
						}

						if(dataSO[j].RetailStockNokia < dataSO[j].BatasBawahRetailALU){
							result = dataSO[j].BatasAtasRetailALU - dataSO[j].RetailStockNokia
							witelSlice[i].QtyKirimRetailALU += roundToNearest(result, 6)
						}

						if(dataSO[j].PremiumStockZTE < dataSO[j].BatasBawahPremiumZTE){
							result = dataSO[j].BatasAtasPremiumZTE - dataSO[j].PremiumStockZTE
							witelSlice[i].QtyKirimPremiumZTE += roundToNearest(result, 10)
						}

						if(dataSO[j].PremiumStockHuawei < dataSO[j].BatasBawahPremiumHW){
							result = dataSO[j].BatasAtasPremiumHW - dataSO[j].PremiumStockHuawei
							witelSlice[i].QtyKirimPremiumHW += roundToNearest(result, 12)
						}

						if(dataSO[j].PremiumStockFiberhome < dataSO[j].BatasBawahPremiumFH){
							result = dataSO[j].BatasAtasPremiumFH - dataSO[j].PremiumStockFiberhome
							witelSlice[i].QtyKirimPremiumFH += roundToNearest(result, 8)
						}

						if(float64(dataSO[j].RetailStockZTE - dataSO[j].RetailZTE + dataSO[j].OnDeliveryRetailZTE) < -(float64(dataSO[j].RetailZTE) * 0.75)){
							witelSlice[i].BlinkRetailZTE = 1
						}

						if(float64(dataSO[j].RetailStockHuawei - dataSO[j].RetailHW + dataSO[j].OnDeliveryRetailHuawei) < -(float64(dataSO[j].RetailHW) * 0.75)){
							witelSlice[i].BlinkRetailHW = 1
						}

						if(float64(dataSO[j].RetailStockFiberhome - dataSO[j].RetailFH + dataSO[j].OnDeliveryRetailFiberhome) < -(float64(dataSO[j].RetailFH) * 0.75)){
							witelSlice[i].BlinkRetailFH = 1
						}

						if(float64(dataSO[j].RetailStockNokia - dataSO[j].RetailALU + dataSO[j].OnDeliveryRetailNokia) < -(float64(dataSO[j].RetailALU) * 0.75)){
							witelSlice[i].BlinkRetailALU = 1
						}

						if(float64(dataSO[j].PremiumStockZTE - dataSO[j].PremiumZTE + dataSO[j].OnDeliveryPremiumZTE) < -(float64(dataSO[j].PremiumZTE) * 0.75)){
							witelSlice[i].BlinkPremiumZTE = 1
						}

						if(float64(dataSO[j].PremiumStockHuawei - dataSO[j].PremiumHW + dataSO[j].OnDeliveryPremiumHuawei) < -(float64(dataSO[j].PremiumHW) * 0.75)){
							witelSlice[i].BlinkPremiumHW = 1
						}

						if(float64(dataSO[j].PremiumStockFiberhome - dataSO[j].PremiumFH + dataSO[j].OnDeliveryPremiumFiberhome) < -(float64(dataSO[j].PremiumFH) * 0.75)){
							witelSlice[i].BlinkPremiumFH = 1
						}

						witelSlice[i].TotalRetail += dataSO[j].TotalRetail
						witelSlice[i].TotalPremium += dataSO[j].TotalPremium
						witelSlice[i].TotalRetailStock += dataSO[j].TotalRetailStock
						witelSlice[i].TotalPremiumStock += dataSO[j].TotalPremiumStock
						witelSlice[i].OnDeliveryTotalRetail += dataSO[j].OnDeliveryRetailHuawei + dataSO[j].OnDeliveryRetailFiberhome + dataSO[j].OnDeliveryRetailZTE + dataSO[j].OnDeliveryRetailNokia
						witelSlice[i].OnDeliveryTotalPremium += dataSO[j].OnDeliveryPremiumHuawei + dataSO[j].OnDeliveryPremiumFiberhome + dataSO[j].OnDeliveryPremiumZTE
					}
				}
			}
		}else{
			dataSO := filterDataWitelByWitel(tregSlice,lokasiWH)

			for j:= 0; j < len(dataSO); j++ {
				if(dataSO[j].RetailStockZTE < dataSO[j].BatasBawahRetailZTE){
					result = dataSO[j].BatasAtasRetailZTE - dataSO[j].RetailStockZTE
					dataSO[j].QtyKirimRetailZTE += roundToNearest(result, 10)
				}

				if(dataSO[j].RetailStockHuawei < dataSO[j].BatasBawahRetailHW){
					result = dataSO[j].BatasAtasRetailHW - dataSO[j].RetailStockHuawei
					dataSO[j].QtyKirimRetailHW += roundToNearest(result, 10)
				}

				if(dataSO[j].RetailStockFiberhome < dataSO[j].BatasBawahRetailFH){
					result = dataSO[j].BatasAtasRetailFH - dataSO[j].RetailStockFiberhome
					dataSO[j].QtyKirimRetailFH += roundToNearest(result, 10)
				}

				if(dataSO[j].RetailStockNokia < dataSO[j].BatasBawahRetailALU){
					result = dataSO[j].BatasAtasRetailALU - dataSO[j].RetailStockNokia
					dataSO[j].QtyKirimRetailALU += roundToNearest(result, 6)
				}

				if(dataSO[j].PremiumStockZTE < dataSO[j].BatasBawahPremiumZTE){
					result = dataSO[j].BatasAtasPremiumZTE - dataSO[j].PremiumStockZTE
					dataSO[j].QtyKirimPremiumZTE += roundToNearest(result, 10)
				}

				if(dataSO[j].PremiumStockHuawei < dataSO[j].BatasBawahPremiumHW){
					result = dataSO[j].BatasAtasPremiumHW - dataSO[j].PremiumStockHuawei
					dataSO[j].QtyKirimPremiumHW += roundToNearest(result, 12)
				}

				if(dataSO[j].PremiumStockFiberhome < dataSO[j].BatasBawahPremiumFH){
					result = dataSO[j].BatasAtasPremiumFH - dataSO[j].PremiumStockFiberhome
					dataSO[j].QtyKirimPremiumFH += roundToNearest(result, 8)
				}

				if(float64(dataSO[j].RetailStockZTE - dataSO[j].RetailZTE + dataSO[j].OnDeliveryRetailZTE) < -(float64(dataSO[j].RetailZTE) * 0.75)){
					dataSO[j].BlinkRetailZTE = 1
				}

				if(float64(dataSO[j].RetailStockHuawei - dataSO[j].RetailHW + dataSO[j].OnDeliveryRetailHuawei) < -(float64(dataSO[j].RetailHW) * 0.75)){
					dataSO[j].BlinkRetailHW = 1
				}

				if(float64(dataSO[j].RetailStockFiberhome - dataSO[j].RetailFH + dataSO[j].OnDeliveryRetailFiberhome) < -(float64(dataSO[j].RetailFH) * 0.75)){
					dataSO[j].BlinkRetailFH = 1
				}

				if(float64(dataSO[j].RetailStockNokia - dataSO[j].RetailALU + dataSO[j].OnDeliveryRetailNokia) < -(float64(dataSO[j].RetailALU) * 0.75)){
					dataSO[j].BlinkRetailALU = 1
				}

				if(float64(dataSO[j].PremiumStockZTE - dataSO[j].PremiumZTE + dataSO[j].OnDeliveryPremiumZTE) < -(float64(dataSO[j].PremiumZTE) * 0.75)){
					dataSO[j].BlinkPremiumZTE = 1
				}

				if(float64(dataSO[j].PremiumStockHuawei - dataSO[j].PremiumHW + dataSO[j].OnDeliveryPremiumHuawei) < -(float64(dataSO[j].PremiumHW) * 0.75)){
					dataSO[j].BlinkPremiumHW = 1
				}

				if(float64(dataSO[j].PremiumStockFiberhome - dataSO[j].PremiumFH + dataSO[j].OnDeliveryPremiumFiberhome) < -(float64(dataSO[j].PremiumFH) * 0.75)){
					dataSO[j].BlinkPremiumFH = 1
				}
			}

			data["response"] = dataSO
		}
		
	}
	return data
}

func (s *dataTmpService) RekapDeliveryTREG() map[string]interface{}{
	data := s.CountDataPerWitelTmp()
	// data = s.HitungQtyKirim(data)
	var result int
	var sumAllZTE, sumAllFH, sumAllHW, sumAllALU, sumAllPZTE, sumAllPFH, sumAllPHW int
	waktuUpdate, waktuDibuat := s.dataTmpRepository.GetTableLastUpdate()
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
		data["last_update"] = waktuDibuat
	}

	if witelSlice, ok := data["witel"].([]domain.TREGMinimumResponse); ok {
		for i := range witelSlice{
			lokasiWH := witelSlice[i].LokasiWH

			var witel []string

			sumAllZTE = 0
			sumAllFH = 0
			sumAllHW = 0
			sumAllALU = 0

			sumAllPZTE = 0
			sumAllPFH = 0
			sumAllPHW = 0

			if(strings.Contains(lokasiWH, "TREG")){
				witel = s.gudangRepository.GetAllSOFromTREG(lokasiWH)

				for k := range witel {
					dataSO := filterDataWitelByTREG(witelSlice ,witel[k])

					sumZTE := 0
					sumFH := 0
					sumHW := 0
					sumALU := 0

					sumPZTE := 0
					sumPFH := 0
					sumPHW := 0

					for j := range dataSO{
						if(dataSO[j].RetailStockZTE < dataSO[j].BatasBawahRetailZTE){
							result = dataSO[j].BatasAtasRetailZTE - dataSO[j].RetailStockZTE
							sumZTE += roundToNearest(result, 10)
						}

						if(dataSO[j].RetailStockHuawei < dataSO[j].BatasBawahRetailHW){
							result = dataSO[j].BatasAtasRetailHW - dataSO[j].RetailStockHuawei
							sumHW += roundToNearest(result, 10)
						}

						if(dataSO[j].RetailStockFiberhome < dataSO[j].BatasBawahRetailFH){
							result = dataSO[j].BatasAtasRetailFH - dataSO[j].RetailStockFiberhome
							sumFH += roundToNearest(result, 10)
						}

						if(dataSO[j].RetailStockNokia < dataSO[j].BatasBawahRetailALU){
							result = dataSO[j].BatasAtasRetailALU - dataSO[j].RetailStockNokia
							sumALU += roundToNearest(result, 6)
						}

						if(dataSO[j].PremiumStockZTE < dataSO[j].BatasBawahPremiumZTE){
							result = dataSO[j].BatasAtasPremiumZTE - dataSO[j].PremiumStockZTE
							sumPZTE += roundToNearest(result, 10)
						}

						if(dataSO[j].PremiumStockHuawei < dataSO[j].BatasBawahPremiumHW){
							result = dataSO[j].BatasAtasPremiumHW - dataSO[j].PremiumStockHuawei
							sumPHW += roundToNearest(result, 12)
						}

						if(dataSO[j].PremiumStockFiberhome < dataSO[j].BatasBawahPremiumFH){
							result = dataSO[j].BatasAtasPremiumFH - dataSO[j].PremiumStockFiberhome
							sumPFH += roundToNearest(result, 8)
						}

						if(float64(dataSO[j].RetailStockZTE - dataSO[j].RetailZTE + dataSO[j].OnDeliveryRetailZTE) < -(float64(dataSO[j].RetailZTE) * 0.75)){
							witelSlice[i].BlinkRetailZTE = 1
						}

						if(float64(dataSO[j].RetailStockHuawei - dataSO[j].RetailHW + dataSO[j].OnDeliveryRetailHuawei) < -(float64(dataSO[j].RetailHW) * 0.75)){
							witelSlice[i].BlinkRetailHW = 1
						}

						if(float64(dataSO[j].RetailStockFiberhome - dataSO[j].RetailFH + dataSO[j].OnDeliveryRetailFiberhome) < -(float64(dataSO[j].RetailFH) * 0.75)){
							witelSlice[i].BlinkRetailFH = 1
						}

						if(float64(dataSO[j].RetailStockNokia - dataSO[j].RetailALU + dataSO[j].OnDeliveryRetailNokia) < -(float64(dataSO[j].RetailALU) * 0.75)){
							witelSlice[i].BlinkRetailALU = 1
						}

						if(float64(dataSO[j].PremiumStockZTE - dataSO[j].PremiumZTE + dataSO[j].OnDeliveryPremiumZTE) < -(float64(dataSO[j].PremiumZTE) * 0.75)){
							witelSlice[i].BlinkPremiumZTE = 1
						}

						if(float64(dataSO[j].PremiumStockHuawei - dataSO[j].PremiumHW + dataSO[j].OnDeliveryPremiumHuawei) < -(float64(dataSO[j].PremiumHW) * 0.75)){
							witelSlice[i].BlinkPremiumHW = 1
						}

						if(float64(dataSO[j].PremiumStockFiberhome - dataSO[j].PremiumFH + dataSO[j].OnDeliveryPremiumFiberhome) < -(float64(dataSO[j].PremiumFH) * 0.75)){
							witelSlice[i].BlinkPremiumFH = 1
						}
					}

					sumAllZTE += sumZTE;
					sumAllFH += sumFH;
					sumAllHW += sumHW;
					sumAllALU += sumALU;

					sumAllPZTE += sumPZTE;
					sumAllPFH += sumPFH;
					sumAllPHW += sumPHW;
				}
				witelSlice[i].BatasAtasRetailFH = int(witelSlice[i].RetailFH * 120 / 100)
				witelSlice[i].BatasAtasRetailHW = int(witelSlice[i].RetailHW * 120 / 100)
				witelSlice[i].BatasAtasRetailZTE = int(witelSlice[i].RetailZTE * 120 / 100)
				witelSlice[i].BatasAtasRetailALU = int(witelSlice[i].RetailALU * 120 / 100)

				witelSlice[i].BatasAtasPremiumFH = int(witelSlice[i].PremiumFH * 120 / 100)
				witelSlice[i].BatasAtasPremiumHW = int(witelSlice[i].PremiumHW * 120 / 100)
				witelSlice[i].BatasAtasPremiumZTE = int(witelSlice[i].PremiumZTE * 120 / 100)

				witelSlice[i].BatasBawahRetailFH = int(witelSlice[i].RetailFH * 70 / 100)
				witelSlice[i].BatasBawahRetailHW = int(witelSlice[i].RetailHW * 70 / 100)
				witelSlice[i].BatasBawahRetailZTE = int(witelSlice[i].RetailZTE * 70 / 100)
				witelSlice[i].BatasBawahRetailALU = int(witelSlice[i].RetailALU * 70 / 100)

				witelSlice[i].BatasBawahPremiumFH = int(witelSlice[i].PremiumFH * 70 / 100)
				witelSlice[i].BatasBawahPremiumHW = int(witelSlice[i].PremiumHW * 70 / 100)
				witelSlice[i].BatasBawahPremiumZTE = int(witelSlice[i].PremiumZTE * 70 / 100)
				
			}else if(strings.Contains(lokasiWH, "WITEL")){
				witel = s.gudangRepository.GetAllSOFromWitel(lokasiWH)

				for k := range witel {
					dataSO := filterDataWitelByWitel(witelSlice ,witel[k])

					sumZTE := 0
					sumFH := 0
					sumHW := 0
					sumALU := 0

					sumPZTE := 0
					sumPFH := 0
					sumPHW := 0

					for j := range dataSO{
						if(dataSO[j].RetailStockZTE < dataSO[j].BatasBawahRetailZTE){
							result = dataSO[j].BatasAtasRetailZTE - dataSO[j].RetailStockZTE
							sumZTE += roundToNearest(result, 10)
						}

						if(dataSO[j].RetailStockHuawei < dataSO[j].BatasBawahRetailHW){
							result = dataSO[j].BatasAtasRetailHW - dataSO[j].RetailStockHuawei
							sumHW += roundToNearest(result, 10)
						}

						if(dataSO[j].RetailStockFiberhome < dataSO[j].BatasBawahRetailFH){
							result = dataSO[j].BatasAtasRetailFH - dataSO[j].RetailStockFiberhome
							sumFH += roundToNearest(result, 10)
						}

						if(dataSO[j].RetailStockNokia < dataSO[j].BatasBawahRetailALU){
							result = dataSO[j].BatasAtasRetailALU - dataSO[j].RetailStockNokia
							sumALU += roundToNearest(result, 6)
						}

						if(dataSO[j].PremiumStockZTE < dataSO[j].BatasBawahPremiumZTE){
							result = dataSO[j].BatasAtasPremiumZTE - dataSO[j].PremiumStockZTE
							sumPZTE += roundToNearest(result, 10)
						}

						if(dataSO[j].PremiumStockHuawei < dataSO[j].BatasBawahPremiumHW){
							result = dataSO[j].BatasAtasPremiumHW - dataSO[j].PremiumStockHuawei
							sumPHW += roundToNearest(result, 12)
						}

						if(dataSO[j].PremiumStockFiberhome < dataSO[j].BatasBawahPremiumFH){
							result = dataSO[j].BatasAtasPremiumFH - dataSO[j].PremiumStockFiberhome
							sumPHW += roundToNearest(result, 8)
						}

						if(float64(dataSO[j].RetailStockZTE - dataSO[j].RetailZTE + dataSO[j].OnDeliveryRetailZTE) < -(float64(dataSO[j].RetailZTE) * 0.75)){
							witelSlice[i].BlinkRetailZTE = 1
						}

						if(float64(dataSO[j].RetailStockHuawei - dataSO[j].RetailHW + dataSO[j].OnDeliveryRetailHuawei) < -(float64(dataSO[j].RetailHW) * 0.75)){
							witelSlice[i].BlinkRetailHW = 1
						}

						if(float64(dataSO[j].RetailStockFiberhome - dataSO[j].RetailFH + dataSO[j].OnDeliveryRetailFiberhome) < -(float64(dataSO[j].RetailFH) * 0.75)){
							witelSlice[i].BlinkRetailFH = 1
						}

						if(float64(dataSO[j].RetailStockNokia - dataSO[j].RetailALU + dataSO[j].OnDeliveryRetailNokia) < -(float64(dataSO[j].RetailALU) * 0.75)){
							witelSlice[i].BlinkRetailALU = 1
						}

						if(float64(dataSO[j].PremiumStockZTE - dataSO[j].PremiumZTE + dataSO[j].OnDeliveryPremiumZTE) < -(float64(dataSO[j].PremiumZTE) * 0.75)){
							witelSlice[i].BlinkPremiumZTE = 1
						}

						if(float64(dataSO[j].PremiumStockHuawei - dataSO[j].PremiumHW + dataSO[j].OnDeliveryPremiumHuawei) < -(float64(dataSO[j].PremiumHW) * 0.75)){
							witelSlice[i].BlinkPremiumHW = 1
						}

						if(float64(dataSO[j].PremiumStockFiberhome - dataSO[j].PremiumFH + dataSO[j].OnDeliveryPremiumFiberhome) < -(float64(dataSO[j].PremiumFH) * 0.75)){
							witelSlice[i].BlinkPremiumFH = 1
						}
					}

					sumAllZTE += sumZTE;
					sumAllFH += sumFH;
					sumAllHW += sumHW;
					sumAllALU += sumALU;

					sumAllPZTE += sumPZTE;
					sumAllPFH += sumPFH;
					sumAllPHW += sumPHW;
				}
			}else{
				dataSO := filterDataWitelBySO(witelSlice ,lokasiWH)
				
				if(dataSO.RetailStockZTE < dataSO.BatasBawahRetailZTE){
					result = dataSO.BatasAtasRetailZTE - dataSO.RetailStockZTE
					sumAllZTE += roundToNearest(result, 10)
				}

				if(dataSO.RetailStockHuawei < dataSO.BatasBawahRetailHW){
					result = dataSO.BatasAtasRetailHW - dataSO.RetailStockHuawei
					sumAllHW += roundToNearest(result, 10)
				}

				if(dataSO.RetailStockFiberhome < dataSO.BatasBawahRetailFH){
					result = dataSO.BatasAtasRetailFH - dataSO.RetailStockFiberhome
					sumAllFH += roundToNearest(result, 10)
				}

				if(dataSO.RetailStockNokia < dataSO.BatasBawahRetailALU){
					result = dataSO.BatasAtasRetailALU - dataSO.RetailStockNokia
					sumAllALU += roundToNearest(result, 6)
				}

				if(dataSO.PremiumStockZTE < dataSO.BatasBawahPremiumZTE){
					result = dataSO.BatasAtasPremiumZTE - dataSO.PremiumStockZTE
					sumAllPZTE += roundToNearest(result, 10)
				}

				if(dataSO.PremiumStockHuawei < dataSO.BatasBawahPremiumHW){
					result = dataSO.BatasAtasPremiumHW - dataSO.PremiumStockHuawei
					sumAllPHW += roundToNearest(result, 12)
				}

				if(dataSO.PremiumStockFiberhome < dataSO.BatasBawahPremiumFH){
					result = dataSO.BatasAtasPremiumFH - dataSO.PremiumStockFiberhome
					sumAllPFH += roundToNearest(result, 8)
				}

				if(float64(dataSO.RetailStockZTE - dataSO.RetailZTE + dataSO.OnDeliveryRetailZTE) < -(float64(dataSO.RetailZTE) * 0.75)){
					witelSlice[i].BlinkRetailZTE = 1
				}

				if(float64(dataSO.RetailStockHuawei - dataSO.RetailHW + dataSO.OnDeliveryRetailHuawei) < -(float64(dataSO.RetailHW) * 0.75)){
					witelSlice[i].BlinkRetailHW = 1
				}

				if(float64(dataSO.RetailStockFiberhome - dataSO.RetailFH + dataSO.OnDeliveryRetailFiberhome) < -(float64(dataSO.RetailFH) * 0.75)){
					witelSlice[i].BlinkRetailFH = 1
				}

				if(float64(dataSO.RetailStockNokia - dataSO.RetailALU + dataSO.OnDeliveryRetailNokia) < -(float64(dataSO.RetailALU) * 0.75)){
					witelSlice[i].BlinkRetailALU = 1
				}

				if(float64(dataSO.PremiumStockZTE - dataSO.PremiumZTE + dataSO.OnDeliveryPremiumZTE) < -(float64(dataSO.PremiumZTE) * 0.75)){
					witelSlice[i].BlinkPremiumZTE = 1
				}

				if(float64(dataSO.PremiumStockHuawei - dataSO.PremiumHW + dataSO.OnDeliveryPremiumHuawei) < -(float64(dataSO.PremiumHW) * 0.75)){
					witelSlice[i].BlinkPremiumHW = 1
				}

				if(float64(dataSO.PremiumStockFiberhome - dataSO.PremiumFH + dataSO.OnDeliveryPremiumFiberhome) < -(float64(dataSO.PremiumFH) * 0.75)){
					witelSlice[i].BlinkPremiumFH = 1
				}
			}	

			witelSlice[i].QtyKirimRetailZTE = sumAllZTE
			witelSlice[i].QtyKirimRetailFH = sumAllFH
			witelSlice[i].QtyKirimRetailHW = sumAllHW
			witelSlice[i].QtyKirimRetailALU = sumAllALU

			witelSlice[i].QtyKirimPremiumZTE = sumAllPZTE
			witelSlice[i].QtyKirimPremiumFH = sumAllPFH
			witelSlice[i].QtyKirimPremiumHW = sumAllPHW

			witelSlice[i].MinimumQty = strconv.Itoa(witelSlice[i].TotalRetail + witelSlice[i].TotalPremium)
		}

		data["jenis_warehouse"] = "Witel"
		data["treg"] = filterDataTREGMinimumResponseByTREG(witelSlice)

		return data
	}

	return data
}

func roundToNearest(n, m int) int {
    // Add m/2 to n before integer division for proper rounding
    return ((n + m/2) / m) * m
}

func (s *dataTmpService) HitungQtyKirim(data map[string]interface{}) map[string]interface{}{
	var result int
	if witelSlice, ok := data["witel"].([]domain.TREGMinimumResponse); ok {
		for i := range witelSlice {
			if(witelSlice[i].RetailStockFiberhome < witelSlice[i].BatasBawahRetailFH){
				result = witelSlice[i].BatasAtasRetailFH - witelSlice[i].RetailStockFiberhome
				witelSlice[i].QtyKirimRetailFH = roundToNearest(result, 10)
			}

			if(witelSlice[i].RetailStockHuawei < witelSlice[i].BatasBawahRetailHW){
				result = witelSlice[i].BatasAtasRetailHW - witelSlice[i].RetailStockHuawei
				witelSlice[i].QtyKirimRetailHW = roundToNearest(result, 10)
			}

			if(witelSlice[i].RetailStockZTE < witelSlice[i].BatasBawahRetailZTE){
				result = witelSlice[i].BatasAtasRetailZTE - witelSlice[i].RetailStockZTE
				witelSlice[i].QtyKirimRetailZTE = roundToNearest(result, 10)
			}

			if(witelSlice[i].RetailStockNokia < witelSlice[i].BatasBawahRetailALU){
				result = witelSlice[i].BatasAtasRetailALU - witelSlice[i].RetailStockNokia
				witelSlice[i].QtyKirimRetailALU = roundToNearest(result, 6)
			}

			if(witelSlice[i].PremiumStockFiberhome < witelSlice[i].BatasBawahPremiumFH){
				result = witelSlice[i].BatasAtasPremiumFH - witelSlice[i].PremiumStockFiberhome
				witelSlice[i].QtyKirimPremiumFH = roundToNearest(result, 8)
			}

			if(witelSlice[i].PremiumStockHuawei < witelSlice[i].BatasBawahPremiumHW){
				result = witelSlice[i].BatasAtasPremiumHW - witelSlice[i].PremiumStockHuawei
				witelSlice[i].QtyKirimPremiumHW = roundToNearest(result, 12)
			}

			if(witelSlice[i].PremiumStockZTE < witelSlice[i].BatasBawahPremiumZTE){
				result = witelSlice[i].BatasAtasPremiumZTE - witelSlice[i].PremiumStockZTE
				witelSlice[i].QtyKirimPremiumZTE = roundToNearest(result, 10)
			}
		}
	}else {
        fmt.Println("data['witel'] is not of type []domain.TREGMinimumResponse")
    }

    return data
}

func (s *dataTmpService) HitungQtyKirimTreg(data map[string]interface{}) map[string]interface{}{
	if witelSlice, ok := data["witel"].([]domain.TREGMinimumResponse); ok {
		for i := range witelSlice {
			valueRegional := witelSlice[i].Regional
			witel := s.gudangRepository.GetAllSOFromTREG(valueRegional);

			sumAllZTE := 0
			sumAllFH := 0
			sumAllHW := 0
			sumAllALU := 0

			sumAllPZTE := 0
			sumAllPFH := 0
			sumAllPHW := 0

			for k := range witel {
				dataSO := filterDataWitelByWitel(witelSlice, witel[k])

				sumZTE := 0
				sumFH := 0
				sumHW := 0
				sumALU := 0

				sumPZTE := 0
				sumPFH := 0
				sumPHW := 0

				for j := range dataSO{
					if(dataSO[j].RetailStockZTE < dataSO[j].BatasBawahRetailZTE){
						sumZTE += dataSO[j].BatasAtasRetailZTE - dataSO[j].RetailStockZTE
					}

					if(dataSO[j].RetailStockHuawei < dataSO[j].BatasBawahRetailHW){
						sumHW += dataSO[j].BatasAtasRetailHW - dataSO[j].RetailStockHuawei
					}

					if(dataSO[j].RetailStockFiberhome < dataSO[j].BatasBawahRetailFH){
						sumFH += dataSO[j].BatasAtasRetailFH - dataSO[j].RetailStockFiberhome
					}

					if(dataSO[j].RetailStockNokia < dataSO[j].BatasBawahRetailALU){
						sumALU += dataSO[j].BatasAtasRetailALU - dataSO[j].RetailStockNokia
					}

					if(dataSO[j].PremiumStockZTE < dataSO[j].BatasBawahPremiumZTE){
						sumPZTE += dataSO[j].BatasAtasPremiumZTE - dataSO[j].PremiumStockZTE
					}

					if(dataSO[j].PremiumStockHuawei < dataSO[j].BatasBawahPremiumHW){
						sumPHW += dataSO[j].BatasAtasPremiumHW - dataSO[j].PremiumStockHuawei
					}

					if(dataSO[j].PremiumStockFiberhome < dataSO[j].BatasBawahPremiumFH){
						sumPFH += dataSO[j].BatasAtasPremiumFH - dataSO[j].PremiumStockFiberhome
					}
				}

				sumAllZTE += sumZTE;
				sumAllFH += sumFH;
				sumAllHW += sumHW;
				sumAllALU += sumALU;

				sumAllPZTE += sumPZTE;
				sumAllPFH += sumPFH;
				sumAllPHW += sumPHW;
			}
			objWitel := &witelSlice[i]

			fieldName := "QtyKirimRetailZTE"
			err := setField(objWitel, fieldName, sumAllZTE)
			helper.PanicIfError(err)

			fieldName = "QtyKirimRetailFH"
			err = setField(objWitel, fieldName, sumAllFH)
			helper.PanicIfError(err)

			fieldName = "QtyKirimRetailHW"
			err = setField(objWitel, fieldName, sumAllHW)
			helper.PanicIfError(err)

			fieldName = "QtyKirimRetailALU"
			err = setField(objWitel, fieldName, sumAllALU)
			helper.PanicIfError(err)


			fieldName = "QtyKirimPremiumZTE"
			err = setField(objWitel, fieldName, sumAllPZTE)
			helper.PanicIfError(err)

			fieldName = "QtyKirimPremiumFH"
			err = setField(objWitel, fieldName, sumAllPFH)
			helper.PanicIfError(err)

			fieldName = "QtyKirimPremiumHW"
			err = setField(objWitel, fieldName, sumAllPHW)
			helper.PanicIfError(err)
		}
	}else {
        fmt.Println("data['witel'] is not of type []domain.TREGMinimumResponse")
    }

    return data
}

func setField(obj interface{}, fieldName string, value int) error {
    v := reflect.ValueOf(obj).Elem()
    f := v.FieldByName(fieldName)
    if !f.IsValid() {
        return fmt.Errorf("No such field: %s in obj", fieldName)
    }
    if !f.CanSet() {
        return fmt.Errorf("Cannot set %s field value", fieldName)
    }
    f.SetInt(int64(value))

    return nil
}

func (s *dataTmpService) GetExportDataTmp() ([]byte, string, error){
	var filename string
	var dataTmp []domain.DataTmp

	templatePath := "template/template_data_tmp.xlsx"
    fileExport, err := excelize.OpenFile(templatePath)
    helper.PanicIfError(err)

    // Get the active sheet index
    activeSheetIndexExport := fileExport.GetActiveSheetIndex()

    // Get the name of the active sheet using the index
    activeSheetNameExport := fileExport.GetSheetName(activeSheetIndexExport)

	// Get the current time
	currentTime := time.Now()

	// Format the time as YYYY-MM-DD
	formattedDate := currentTime.Format("2006-01-02")

	filename = "Exported_Data_Tmp "+formattedDate+".xlsx"

	dataTmp = s.dataTmpRepository.GetExportDataTmp();

    // Set values in cells
    for i := range dataTmp{
	    fileExport.SetCellValue(activeSheetNameExport, "A"+strconv.Itoa(i+2), dataTmp[i].Region)
	    fileExport.SetCellValue(activeSheetNameExport, "B"+strconv.Itoa(i+2), dataTmp[i].LokasiWH)
	    fileExport.SetCellValue(activeSheetNameExport, "C"+strconv.Itoa(i+2), dataTmp[i].Status)
	    fileExport.SetCellValue(activeSheetNameExport, "D"+strconv.Itoa(i+2), dataTmp[i].Jumlah)
	    fileExport.SetCellValue(activeSheetNameExport, "E"+strconv.Itoa(i+2), dataTmp[i].Deskripsi)
	}

    // Set active sheet of the workbook
    fileExport.SetActiveSheet(activeSheetIndexExport)

    bytesBuffer, err := fileExport.WriteToBuffer()

	//Export to download
	return bytesBuffer.Bytes(), filename, err
}

func (s *dataTmpService) DownloadTemplateMinimumStock() ([]byte, string, error){
	filename := "template_database_minimum_stock.xlsx"
	templatePath := "template/template_database.xlsx"
    fileExport, err := excelize.OpenFile(templatePath)
    helper.PanicIfError(err)

    bytesBuffer, err := fileExport.WriteToBuffer()

	//Export to download
	return bytesBuffer.Bytes(), filename, err
}

func (s *dataTmpService) DownloadTemplateDataTmp() ([]byte, string, error){
	filename := "template_data_tmp.xlsx"
	templatePath := "template/template_data_tmp.xlsx"
    fileExport, err := excelize.OpenFile(templatePath)
    helper.PanicIfError(err)

    bytesBuffer, err := fileExport.WriteToBuffer()

	//Export to download
	return bytesBuffer.Bytes(), filename, err
}



func (s *dataTmpService) GetExportMinimumStockDatabase() ([]byte, string, error){
	var filename string

	templatePath := "template/template_database.xlsx"
    fileExport, err := excelize.OpenFile(templatePath)
    helper.PanicIfError(err)

    // Get the active sheet index
    activeSheetIndexExport := fileExport.GetActiveSheetIndex()

    // Get the name of the active sheet using the index
    activeSheetNameExport := fileExport.GetSheetName(activeSheetIndexExport)

	// Get the current time
	currentTime := time.Now()

	// Format the time as YYYY-MM-DD
	formattedDate := currentTime.Format("2006-01-02")

	filename = "Exported_Database_Minimum_Stock "+formattedDate+".xlsx"

	dataGudang := s.gudangRepository.GetAllData();

	j := 0
    // Set values in cells
    for i:=4; i < len(dataGudang); i++{
	    fileExport.SetCellValue(activeSheetNameExport, "A"+strconv.Itoa(i), dataGudang[j].Regional)
	    fileExport.SetCellValue(activeSheetNameExport, "B"+strconv.Itoa(i), dataGudang[j].Witel)
	    fileExport.SetCellValue(activeSheetNameExport, "C"+strconv.Itoa(i), dataGudang[j].LokasiWH)
	    fileExport.SetCellValue(activeSheetNameExport, "D"+strconv.Itoa(i), dataGudang[j].Lokasi)
	    fileExport.SetCellValue(activeSheetNameExport, "E"+strconv.Itoa(i), dataGudang[j].Wilayah)
	    fileExport.SetCellValue(activeSheetNameExport, "F"+strconv.Itoa(i), dataGudang[j].MinimumQty)
	    fileExport.SetCellValue(activeSheetNameExport, "G"+strconv.Itoa(i), dataGudang[j].RetailFH)
	    fileExport.SetCellValue(activeSheetNameExport, "H"+strconv.Itoa(i), dataGudang[j].RetailHW)
	    fileExport.SetCellValue(activeSheetNameExport, "I"+strconv.Itoa(i), dataGudang[j].RetailZTE)
	    fileExport.SetCellValue(activeSheetNameExport, "J"+strconv.Itoa(i), dataGudang[j].RetailALU)
	    fileExport.SetCellValue(activeSheetNameExport, "K"+strconv.Itoa(i), dataGudang[j].PremiumFH)
	    fileExport.SetCellValue(activeSheetNameExport, "L"+strconv.Itoa(i), dataGudang[j].PremiumHW)
	    fileExport.SetCellValue(activeSheetNameExport, "M"+strconv.Itoa(i), dataGudang[j].PremiumZTE)
	    fileExport.SetCellValue(activeSheetNameExport, "N"+strconv.Itoa(i), dataGudang[j].STBZTE)
	    fileExport.SetCellValue(activeSheetNameExport, "O"+strconv.Itoa(i), dataGudang[j].APCisco)
	    fileExport.SetCellValue(activeSheetNameExport, "P"+strconv.Itoa(i), dataGudang[j].APHuawei)
	    j+=1
	}

    // Set active sheet of the workbook
    fileExport.SetActiveSheet(activeSheetIndexExport)

    bytesBuffer, err := fileExport.WriteToBuffer()

	//Export to download
	return bytesBuffer.Bytes(), filename, err
}

func (s *dataTmpService) UploadDataTmp(){
	s.dataTmpRepository.DeleteAllDataTmp()
	templatePath := "template/uploaded_data_tmp.xlsx"
    spreadsheetImport, err := excelize.OpenFile(templatePath)
    helper.PanicIfError(err)

    // Get the active sheet index
    activeSheetIndexUploaded := spreadsheetImport.GetActiveSheetIndex()

    // Get the name of the active sheet using the index
    activeSheetName := spreadsheetImport.GetSheetName(activeSheetIndexUploaded)

    rows, err := spreadsheetImport.GetRows(activeSheetName)
    helper.PanicIfError(err)

    var dataTmps []domain.DataTmp

    // Get the maximum number of rows
    maxRows := len(rows)
    for i := 2; i < maxRows+1; i++{
    	var dataTmp domain.DataTmp

    	dataTmp.Region, err = spreadsheetImport.GetCellValue(activeSheetName, "A"+strconv.Itoa(i))
    	dataTmp.LokasiWH, err = spreadsheetImport.GetCellValue(activeSheetName, "B"+strconv.Itoa(i))
    	dataTmp.Status, err = spreadsheetImport.GetCellValue(activeSheetName, "C"+strconv.Itoa(i))
    	dataTmp.Jumlah, err = spreadsheetImport.GetCellValue(activeSheetName, "D"+strconv.Itoa(i))
    	dataTmp.Deskripsi, err = spreadsheetImport.GetCellValue(activeSheetName, "E"+strconv.Itoa(i))
    	
    	dataTmps = append(dataTmps, dataTmp)
    }

	s.dataTmpRepository.UploadDataTmpBulk(dataTmps)

	// Ensure the file is closed properly
    defer func() {
        if err := spreadsheetImport.Close(); err != nil {
            helper.PanicIfError(err)
        }
    }()
}

func (s *dataTmpService) ExportDataTmp(jenisWarna string, jenisExport string) ([]byte, string, error){
	var filename string

	templatePath := "template/template_qty_kirim.xlsx"
    fileExport, err := excelize.OpenFile(templatePath)
    helper.PanicIfError(err)

	// Get the current time
	currentTime := time.Now()

	// Format the time as YYYY-MM-DD
	formattedDate := currentTime.Format("2006-01-02")
	filename = "Hasil_Rekap_Stock_SCMT_"+formattedDate+".xlsx";

	// Get the active sheet index
    activeSheetIndexExport := fileExport.GetActiveSheetIndex()

    // Get the name of the active sheet using the index
    activeSheetNameExport := fileExport.GetSheetName(activeSheetIndexExport)

    data := s.CountDataPerWitelTmp()
    data = s.HitungQtyKirim(data)

    idxAwal := 0

    styleRed, err := fileExport.NewStyle(&excelize.Style{
        Fill: excelize.Fill{
            Type:    "pattern",
            Color:   []string{"#FF0000"}, // Hex color code for red
            Pattern: 1,
        },
    })

    styleYellow, err := fileExport.NewStyle(&excelize.Style{
        Fill: excelize.Fill{
            Type:    "pattern",
            Color:   []string{"#FFFF00"}, // Hex color code for red
            Pattern: 1,
        },
    })

    if(jenisExport != "all"){
    	if(jenisExport == "treg_only"){
    		data = s.RekapDeliveryTREG()
    		data["witel"] = data["treg"]
    	}else{
    		data = s.RekapDeliveryWitel(jenisExport)
    		data["witel"] = data["response"]
    	}
    }


    if witelSlice, ok := data["witel"].([]domain.TREGMinimumResponse); ok {
		if(jenisWarna == "all"){
			i := 4;
			for j:=idxAwal; j < len(witelSlice); j++ {
				fileExport.SetCellValue(activeSheetNameExport, "A"+strconv.Itoa(i), witelSlice[j].Regional)
				fileExport.SetCellValue(activeSheetNameExport, "B"+strconv.Itoa(i), witelSlice[j].LokasiWH)
				fileExport.SetCellValue(activeSheetNameExport, "C"+strconv.Itoa(i), witelSlice[j].MinimumQty)
				fileExport.SetCellValue(activeSheetNameExport, "D"+strconv.Itoa(i), witelSlice[j].RetailFH)
				fileExport.SetCellValue(activeSheetNameExport, "E"+strconv.Itoa(i), witelSlice[j].RetailHW)
				fileExport.SetCellValue(activeSheetNameExport, "F"+strconv.Itoa(i), witelSlice[j].RetailZTE)
				fileExport.SetCellValue(activeSheetNameExport, "G"+strconv.Itoa(i), witelSlice[j].RetailALU)
				fileExport.SetCellValue(activeSheetNameExport, "H"+strconv.Itoa(i), witelSlice[j].TotalRetail)
				fileExport.SetCellValue(activeSheetNameExport, "I"+strconv.Itoa(i), witelSlice[j].PremiumFH)
				fileExport.SetCellValue(activeSheetNameExport, "J"+strconv.Itoa(i), witelSlice[j].PremiumHW)
				fileExport.SetCellValue(activeSheetNameExport, "K"+strconv.Itoa(i), witelSlice[j].PremiumZTE)
				fileExport.SetCellValue(activeSheetNameExport, "L"+strconv.Itoa(i), witelSlice[j].TotalPremium)
				fileExport.SetCellValue(activeSheetNameExport, "M"+strconv.Itoa(i), witelSlice[j].RetailStockFiberhome)
				fileExport.SetCellValue(activeSheetNameExport, "N"+strconv.Itoa(i), witelSlice[j].RetailStockHuawei)
				fileExport.SetCellValue(activeSheetNameExport, "O"+strconv.Itoa(i), witelSlice[j].RetailStockZTE)
				fileExport.SetCellValue(activeSheetNameExport, "P"+strconv.Itoa(i), witelSlice[j].RetailStockNokia)
				fileExport.SetCellValue(activeSheetNameExport, "Q"+strconv.Itoa(i), witelSlice[j].TotalRetailStock)
				fileExport.SetCellValue(activeSheetNameExport, "R"+strconv.Itoa(i), witelSlice[j].PremiumStockFiberhome)
				fileExport.SetCellValue(activeSheetNameExport, "S"+strconv.Itoa(i), witelSlice[j].PremiumStockHuawei)
				fileExport.SetCellValue(activeSheetNameExport, "T"+strconv.Itoa(i), witelSlice[j].PremiumStockZTE)
				fileExport.SetCellValue(activeSheetNameExport, "U"+strconv.Itoa(i), witelSlice[j].TotalPremiumStock)
				fileExport.SetCellValue(activeSheetNameExport, "V"+strconv.Itoa(i), witelSlice[j].RetailStockFiberhome - witelSlice[j].RetailFH + witelSlice[j].OnDeliveryRetailFiberhome)
				fileExport.SetCellValue(activeSheetNameExport, "W"+strconv.Itoa(i), witelSlice[j].RetailStockHuawei- witelSlice[j].RetailHW + witelSlice[j].OnDeliveryRetailHuawei)
				fileExport.SetCellValue(activeSheetNameExport, "X"+strconv.Itoa(i), witelSlice[j].RetailStockZTE - witelSlice[j].RetailZTE + witelSlice[j].OnDeliveryRetailZTE)
				fileExport.SetCellValue(activeSheetNameExport, "Y"+strconv.Itoa(i), witelSlice[j].RetailStockNokia - witelSlice[j].RetailALU + witelSlice[j].OnDeliveryRetailNokia)
				fileExport.SetCellValue(activeSheetNameExport, "Z"+strconv.Itoa(i), witelSlice[j].TotalRetailStock - witelSlice[j].TotalRetail + witelSlice[j].OnDeliveryTotalRetail)

				fileExport.SetCellValue(activeSheetNameExport, "AA"+strconv.Itoa(i), witelSlice[j].PremiumStockFiberhome - witelSlice[j].PremiumFH + witelSlice[j].OnDeliveryPremiumFiberhome)
				fileExport.SetCellValue(activeSheetNameExport, "AB"+strconv.Itoa(i), witelSlice[j].PremiumStockHuawei - witelSlice[j].PremiumHW + witelSlice[j].OnDeliveryPremiumHuawei)
				fileExport.SetCellValue(activeSheetNameExport, "AC"+strconv.Itoa(i), witelSlice[j].PremiumStockZTE - witelSlice[j].PremiumZTE + witelSlice[j].OnDeliveryPremiumZTE)
				fileExport.SetCellValue(activeSheetNameExport, "AD"+strconv.Itoa(i), witelSlice[j].TotalPremiumStock - witelSlice[j].TotalPremium + witelSlice[j].OnDeliveryTotalPremium)

				fileExport.SetCellValue(activeSheetNameExport, "AE"+strconv.Itoa(i), witelSlice[j].BatasAtasRetailFH)
				fileExport.SetCellValue(activeSheetNameExport, "AF"+strconv.Itoa(i), witelSlice[j].BatasAtasRetailHW)
				fileExport.SetCellValue(activeSheetNameExport, "AG"+strconv.Itoa(i), witelSlice[j].BatasAtasRetailZTE)
				fileExport.SetCellValue(activeSheetNameExport, "AH"+strconv.Itoa(i), witelSlice[j].BatasAtasRetailALU)
				fileExport.SetCellValue(activeSheetNameExport, "AI"+strconv.Itoa(i), witelSlice[j].BatasAtasRetailFH + witelSlice[j].BatasAtasRetailHW + witelSlice[j].BatasAtasRetailZTE + witelSlice[j].BatasAtasRetailALU)

				fileExport.SetCellValue(activeSheetNameExport, "AJ"+strconv.Itoa(i), witelSlice[j].BatasAtasPremiumFH)
				fileExport.SetCellValue(activeSheetNameExport, "AK"+strconv.Itoa(i), witelSlice[j].BatasAtasPremiumHW)
				fileExport.SetCellValue(activeSheetNameExport, "AL"+strconv.Itoa(i), witelSlice[j].BatasAtasPremiumZTE)
				fileExport.SetCellValue(activeSheetNameExport, "AM"+strconv.Itoa(i), witelSlice[j].BatasAtasPremiumFH + witelSlice[j].BatasAtasPremiumHW + witelSlice[j].BatasAtasPremiumZTE)

				fileExport.SetCellValue(activeSheetNameExport, "AN"+strconv.Itoa(i), witelSlice[j].BatasBawahRetailFH)
				fileExport.SetCellValue(activeSheetNameExport, "AO"+strconv.Itoa(i), witelSlice[j].BatasBawahRetailHW)
				fileExport.SetCellValue(activeSheetNameExport, "AP"+strconv.Itoa(i), witelSlice[j].BatasBawahRetailZTE)
				fileExport.SetCellValue(activeSheetNameExport, "AQ"+strconv.Itoa(i), witelSlice[j].BatasBawahRetailALU)
				fileExport.SetCellValue(activeSheetNameExport, "AR"+strconv.Itoa(i), witelSlice[j].BatasBawahRetailFH + witelSlice[j].BatasBawahRetailHW + witelSlice[j].BatasBawahRetailZTE + witelSlice[j].BatasBawahRetailALU)

				fileExport.SetCellValue(activeSheetNameExport, "AS"+strconv.Itoa(i), witelSlice[j].BatasBawahPremiumFH)
				fileExport.SetCellValue(activeSheetNameExport, "AT"+strconv.Itoa(i), witelSlice[j].BatasBawahPremiumHW)
				fileExport.SetCellValue(activeSheetNameExport, "AU"+strconv.Itoa(i), witelSlice[j].BatasBawahPremiumZTE)
				fileExport.SetCellValue(activeSheetNameExport, "AV"+strconv.Itoa(i), witelSlice[j].BatasBawahPremiumFH + witelSlice[j].BatasBawahPremiumHW + witelSlice[j].BatasBawahPremiumZTE)

				fileExport.SetCellValue(activeSheetNameExport, "AW"+strconv.Itoa(i), witelSlice[j].QtyKirimRetailFH)
				fileExport.SetCellValue(activeSheetNameExport, "AX"+strconv.Itoa(i), witelSlice[j].QtyKirimRetailHW)
				fileExport.SetCellValue(activeSheetNameExport, "AY"+strconv.Itoa(i), witelSlice[j].QtyKirimRetailZTE)
				fileExport.SetCellValue(activeSheetNameExport, "AZ"+strconv.Itoa(i), witelSlice[j].QtyKirimRetailALU)
				fileExport.SetCellValue(activeSheetNameExport, "BA"+strconv.Itoa(i), witelSlice[j].QtyKirimRetailFH + witelSlice[j].QtyKirimRetailHW + witelSlice[j].QtyKirimRetailZTE + witelSlice[j].QtyKirimRetailALU)

				fileExport.SetCellValue(activeSheetNameExport, "BB"+strconv.Itoa(i), witelSlice[j].QtyKirimPremiumFH)
				fileExport.SetCellValue(activeSheetNameExport, "BC"+strconv.Itoa(i), witelSlice[j].QtyKirimPremiumHW)
				fileExport.SetCellValue(activeSheetNameExport, "BD"+strconv.Itoa(i), witelSlice[j].QtyKirimPremiumZTE)
				fileExport.SetCellValue(activeSheetNameExport, "BE"+strconv.Itoa(i), witelSlice[j].QtyKirimPremiumFH + witelSlice[j].QtyKirimPremiumHW + witelSlice[j].QtyKirimPremiumZTE)

				fileExport.SetCellValue(activeSheetNameExport, "BF"+strconv.Itoa(i), witelSlice[j].OnDeliveryRetailFiberhome)
				fileExport.SetCellValue(activeSheetNameExport, "BG"+strconv.Itoa(i), witelSlice[j].OnDeliveryRetailHuawei)
				fileExport.SetCellValue(activeSheetNameExport, "BH"+strconv.Itoa(i), witelSlice[j].OnDeliveryRetailZTE)
				fileExport.SetCellValue(activeSheetNameExport, "BI"+strconv.Itoa(i), witelSlice[j].OnDeliveryRetailNokia)
				fileExport.SetCellValue(activeSheetNameExport, "BJ"+strconv.Itoa(i), witelSlice[j].OnDeliveryTotalRetail)

				fileExport.SetCellValue(activeSheetNameExport, "BK"+strconv.Itoa(i), witelSlice[j].OnDeliveryPremiumFiberhome)
				fileExport.SetCellValue(activeSheetNameExport, "BL"+strconv.Itoa(i), witelSlice[j].OnDeliveryPremiumHuawei)
				fileExport.SetCellValue(activeSheetNameExport, "BM"+strconv.Itoa(i), witelSlice[j].OnDeliveryPremiumZTE)
				fileExport.SetCellValue(activeSheetNameExport, "BN"+strconv.Itoa(i), witelSlice[j].OnDeliveryTotalPremium)

				if(float64(witelSlice[j].TotalRetailStock - witelSlice[j].TotalRetail + witelSlice[j].OnDeliveryTotalRetail) < -(float64(witelSlice[j].TotalRetail) * 0.75)){
					if err := fileExport.SetCellStyle(activeSheetNameExport, "Z"+strconv.Itoa(i), "Z"+strconv.Itoa(i), styleRed); err != nil {
				        helper.PanicIfError(err)
				    }
				}else if(witelSlice[j].TotalRetailStock - witelSlice[j].TotalRetail + witelSlice[j].OnDeliveryTotalRetail < 0){
					if err := fileExport.SetCellStyle(activeSheetNameExport, "Z"+strconv.Itoa(i), "Z"+strconv.Itoa(i), styleYellow); err != nil {
				        helper.PanicIfError(err)
				    }
				}

				if(float64(witelSlice[j].TotalPremiumStock - witelSlice[j].TotalPremium + witelSlice[j].OnDeliveryTotalPremium) < -(float64(witelSlice[j].TotalPremium) * 0.75)){
					if err := fileExport.SetCellStyle(activeSheetNameExport, "AD"+strconv.Itoa(i), "AD"+strconv.Itoa(i), styleRed); err != nil {
				        helper.PanicIfError(err)
				    }
				}else if(witelSlice[j].TotalPremiumStock - witelSlice[j].TotalPremium + witelSlice[j].OnDeliveryTotalPremium < 0){
					if err := fileExport.SetCellStyle(activeSheetNameExport, "AD"+strconv.Itoa(i), "AD"+strconv.Itoa(i), styleYellow); err != nil {
				        helper.PanicIfError(err)
				    }
				}
				i += 1
			}
		}else if(jenisWarna == "kuning"){
			i := 4
			for j:=idxAwal; j < len(witelSlice); j ++{
				if((witelSlice[j].TotalRetailStock - witelSlice[j].TotalRetail + witelSlice[j].OnDeliveryTotalRetail < 0 && float64(witelSlice[j].TotalRetailStock - witelSlice[j].TotalRetail + witelSlice[j].OnDeliveryTotalRetail) > -(float64(witelSlice[j].TotalRetail) * 0.75)) || (witelSlice[j].TotalPremiumStock - witelSlice[j].TotalPremium + witelSlice[j].OnDeliveryTotalPremium < 0 && float64(witelSlice[j].TotalPremiumStock - witelSlice[j].TotalPremium + witelSlice[j].OnDeliveryTotalPremium) > - (float64(witelSlice[j].TotalPremium) * 0.75))){
					fileExport.SetCellValue(activeSheetNameExport, "A"+strconv.Itoa(i), witelSlice[j].Regional)
					fileExport.SetCellValue(activeSheetNameExport, "B"+strconv.Itoa(i), witelSlice[j].LokasiWH)
					fileExport.SetCellValue(activeSheetNameExport, "C"+strconv.Itoa(i), witelSlice[j].MinimumQty)
					fileExport.SetCellValue(activeSheetNameExport, "D"+strconv.Itoa(i), witelSlice[j].RetailFH)
					fileExport.SetCellValue(activeSheetNameExport, "E"+strconv.Itoa(i), witelSlice[j].RetailHW)
					fileExport.SetCellValue(activeSheetNameExport, "F"+strconv.Itoa(i), witelSlice[j].RetailZTE)
					fileExport.SetCellValue(activeSheetNameExport, "G"+strconv.Itoa(i), witelSlice[j].RetailALU)
					fileExport.SetCellValue(activeSheetNameExport, "H"+strconv.Itoa(i), witelSlice[j].TotalRetail)
					fileExport.SetCellValue(activeSheetNameExport, "I"+strconv.Itoa(i), witelSlice[j].PremiumFH)
					fileExport.SetCellValue(activeSheetNameExport, "J"+strconv.Itoa(i), witelSlice[j].PremiumHW)
					fileExport.SetCellValue(activeSheetNameExport, "K"+strconv.Itoa(i), witelSlice[j].PremiumZTE)
					fileExport.SetCellValue(activeSheetNameExport, "L"+strconv.Itoa(i), witelSlice[j].TotalPremium)
					fileExport.SetCellValue(activeSheetNameExport, "M"+strconv.Itoa(i), witelSlice[j].RetailStockFiberhome)
					fileExport.SetCellValue(activeSheetNameExport, "N"+strconv.Itoa(i), witelSlice[j].RetailStockHuawei)
					fileExport.SetCellValue(activeSheetNameExport, "O"+strconv.Itoa(i), witelSlice[j].RetailStockZTE)
					fileExport.SetCellValue(activeSheetNameExport, "P"+strconv.Itoa(i), witelSlice[j].RetailStockNokia)
					fileExport.SetCellValue(activeSheetNameExport, "Q"+strconv.Itoa(i), witelSlice[j].TotalRetailStock)
					fileExport.SetCellValue(activeSheetNameExport, "R"+strconv.Itoa(i), witelSlice[j].PremiumStockFiberhome)
					fileExport.SetCellValue(activeSheetNameExport, "S"+strconv.Itoa(i), witelSlice[j].PremiumStockHuawei)
					fileExport.SetCellValue(activeSheetNameExport, "T"+strconv.Itoa(i), witelSlice[j].PremiumStockZTE)
					fileExport.SetCellValue(activeSheetNameExport, "U"+strconv.Itoa(i), witelSlice[j].TotalPremiumStock)
					fileExport.SetCellValue(activeSheetNameExport, "V"+strconv.Itoa(i), witelSlice[j].RetailStockFiberhome - witelSlice[j].RetailFH + witelSlice[j].OnDeliveryRetailFiberhome)
					fileExport.SetCellValue(activeSheetNameExport, "W"+strconv.Itoa(i), witelSlice[j].RetailStockHuawei- witelSlice[j].RetailHW + witelSlice[j].OnDeliveryRetailHuawei)
					fileExport.SetCellValue(activeSheetNameExport, "X"+strconv.Itoa(i), witelSlice[j].RetailStockZTE - witelSlice[j].RetailZTE + witelSlice[j].OnDeliveryRetailZTE)
					fileExport.SetCellValue(activeSheetNameExport, "Y"+strconv.Itoa(i), witelSlice[j].RetailStockNokia - witelSlice[j].RetailALU + witelSlice[j].OnDeliveryRetailNokia)
					fileExport.SetCellValue(activeSheetNameExport, "Z"+strconv.Itoa(i), witelSlice[j].TotalRetailStock - witelSlice[j].TotalRetail + witelSlice[j].OnDeliveryTotalRetail)

					fileExport.SetCellValue(activeSheetNameExport, "AA"+strconv.Itoa(i), witelSlice[j].PremiumStockFiberhome - witelSlice[j].PremiumFH + witelSlice[j].OnDeliveryPremiumFiberhome)
					fileExport.SetCellValue(activeSheetNameExport, "AB"+strconv.Itoa(i), witelSlice[j].PremiumStockHuawei - witelSlice[j].PremiumHW + witelSlice[j].OnDeliveryPremiumHuawei)
					fileExport.SetCellValue(activeSheetNameExport, "AC"+strconv.Itoa(i), witelSlice[j].PremiumStockZTE - witelSlice[j].PremiumZTE + witelSlice[j].OnDeliveryPremiumZTE)
					fileExport.SetCellValue(activeSheetNameExport, "AD"+strconv.Itoa(i), witelSlice[j].TotalPremiumStock - witelSlice[j].TotalPremium + witelSlice[j].OnDeliveryTotalPremium)

					fileExport.SetCellValue(activeSheetNameExport, "AE"+strconv.Itoa(i), witelSlice[j].BatasAtasRetailFH)
					fileExport.SetCellValue(activeSheetNameExport, "AF"+strconv.Itoa(i), witelSlice[j].BatasAtasRetailHW)
					fileExport.SetCellValue(activeSheetNameExport, "AG"+strconv.Itoa(i), witelSlice[j].BatasAtasRetailZTE)
					fileExport.SetCellValue(activeSheetNameExport, "AH"+strconv.Itoa(i), witelSlice[j].BatasAtasRetailALU)
					fileExport.SetCellValue(activeSheetNameExport, "AI"+strconv.Itoa(i), witelSlice[j].BatasAtasRetailFH + witelSlice[j].BatasAtasRetailHW + witelSlice[j].BatasAtasRetailZTE + witelSlice[j].BatasAtasRetailALU)

					fileExport.SetCellValue(activeSheetNameExport, "AJ"+strconv.Itoa(i), witelSlice[j].BatasAtasPremiumFH)
					fileExport.SetCellValue(activeSheetNameExport, "AK"+strconv.Itoa(i), witelSlice[j].BatasAtasPremiumHW)
					fileExport.SetCellValue(activeSheetNameExport, "AL"+strconv.Itoa(i), witelSlice[j].BatasAtasPremiumZTE)
					fileExport.SetCellValue(activeSheetNameExport, "AM"+strconv.Itoa(i), witelSlice[j].BatasAtasPremiumFH + witelSlice[j].BatasAtasPremiumHW + witelSlice[j].BatasAtasPremiumZTE)

					fileExport.SetCellValue(activeSheetNameExport, "AN"+strconv.Itoa(i), witelSlice[j].BatasBawahRetailFH)
					fileExport.SetCellValue(activeSheetNameExport, "AO"+strconv.Itoa(i), witelSlice[j].BatasBawahRetailHW)
					fileExport.SetCellValue(activeSheetNameExport, "AP"+strconv.Itoa(i), witelSlice[j].BatasBawahRetailZTE)
					fileExport.SetCellValue(activeSheetNameExport, "AQ"+strconv.Itoa(i), witelSlice[j].BatasBawahRetailALU)
					fileExport.SetCellValue(activeSheetNameExport, "AR"+strconv.Itoa(i), witelSlice[j].BatasBawahRetailFH + witelSlice[j].BatasBawahRetailHW + witelSlice[j].BatasBawahRetailZTE + witelSlice[j].BatasBawahRetailALU)

					fileExport.SetCellValue(activeSheetNameExport, "AS"+strconv.Itoa(i), witelSlice[j].BatasBawahPremiumFH)
					fileExport.SetCellValue(activeSheetNameExport, "AT"+strconv.Itoa(i), witelSlice[j].BatasBawahPremiumHW)
					fileExport.SetCellValue(activeSheetNameExport, "AU"+strconv.Itoa(i), witelSlice[j].BatasBawahPremiumZTE)
					fileExport.SetCellValue(activeSheetNameExport, "AV"+strconv.Itoa(i), witelSlice[j].BatasBawahPremiumFH + witelSlice[j].BatasBawahPremiumHW + witelSlice[j].BatasBawahPremiumZTE)

					fileExport.SetCellValue(activeSheetNameExport, "AW"+strconv.Itoa(i), witelSlice[j].QtyKirimRetailFH)
					fileExport.SetCellValue(activeSheetNameExport, "AX"+strconv.Itoa(i), witelSlice[j].QtyKirimRetailHW)
					fileExport.SetCellValue(activeSheetNameExport, "AY"+strconv.Itoa(i), witelSlice[j].QtyKirimRetailZTE)
					fileExport.SetCellValue(activeSheetNameExport, "AZ"+strconv.Itoa(i), witelSlice[j].QtyKirimRetailALU)
					fileExport.SetCellValue(activeSheetNameExport, "BA"+strconv.Itoa(i), witelSlice[j].QtyKirimRetailFH + witelSlice[j].QtyKirimRetailHW + witelSlice[j].QtyKirimRetailZTE + witelSlice[j].QtyKirimRetailALU)

					fileExport.SetCellValue(activeSheetNameExport, "BB"+strconv.Itoa(i), witelSlice[j].QtyKirimPremiumFH)
					fileExport.SetCellValue(activeSheetNameExport, "BC"+strconv.Itoa(i), witelSlice[j].QtyKirimPremiumHW)
					fileExport.SetCellValue(activeSheetNameExport, "BD"+strconv.Itoa(i), witelSlice[j].QtyKirimPremiumZTE)
					fileExport.SetCellValue(activeSheetNameExport, "BE"+strconv.Itoa(i), witelSlice[j].QtyKirimPremiumFH + witelSlice[j].QtyKirimPremiumHW + witelSlice[j].QtyKirimPremiumZTE)

					fileExport.SetCellValue(activeSheetNameExport, "BF"+strconv.Itoa(i), witelSlice[j].OnDeliveryRetailFiberhome)
					fileExport.SetCellValue(activeSheetNameExport, "BG"+strconv.Itoa(i), witelSlice[j].OnDeliveryRetailHuawei)
					fileExport.SetCellValue(activeSheetNameExport, "BH"+strconv.Itoa(i), witelSlice[j].OnDeliveryRetailZTE)
					fileExport.SetCellValue(activeSheetNameExport, "BI"+strconv.Itoa(i), witelSlice[j].OnDeliveryRetailNokia)
					fileExport.SetCellValue(activeSheetNameExport, "BJ"+strconv.Itoa(i), witelSlice[j].OnDeliveryTotalRetail)

					fileExport.SetCellValue(activeSheetNameExport, "BK"+strconv.Itoa(i), witelSlice[j].OnDeliveryPremiumFiberhome)
					fileExport.SetCellValue(activeSheetNameExport, "BL"+strconv.Itoa(i), witelSlice[j].OnDeliveryPremiumHuawei)
					fileExport.SetCellValue(activeSheetNameExport, "BM"+strconv.Itoa(i), witelSlice[j].OnDeliveryPremiumZTE)
					fileExport.SetCellValue(activeSheetNameExport, "BN"+strconv.Itoa(i), witelSlice[j].OnDeliveryTotalPremium)

					if(float64(witelSlice[j].TotalRetailStock - witelSlice[j].TotalRetail + witelSlice[j].OnDeliveryTotalRetail) < 0 && float64(witelSlice[j].TotalRetailStock - witelSlice[j].TotalRetail + witelSlice[j].OnDeliveryTotalRetail) > -(float64(witelSlice[j].TotalRetail) * 0.75)){
						if err := fileExport.SetCellStyle(activeSheetNameExport, "Z"+strconv.Itoa(i), "Z"+strconv.Itoa(i), styleYellow); err != nil {
					        helper.PanicIfError(err)
					    }
					}

					if(float64(witelSlice[j].TotalPremiumStock - witelSlice[j].TotalPremium + witelSlice[j].OnDeliveryTotalPremium) < 0 && float64(witelSlice[j].TotalPremiumStock - witelSlice[j].TotalPremium + witelSlice[j].OnDeliveryTotalPremium) > -(float64(witelSlice[j].TotalPremium) * 0.75)){
						if err := fileExport.SetCellStyle(activeSheetNameExport, "AD"+strconv.Itoa(i), "AD"+strconv.Itoa(i), styleYellow); err != nil {
					        helper.PanicIfError(err)
					    }
					}

					i += 1;
				}
			}
			filename = "Hasil_Rekap_Stock_SCMT_Kuning_"+formattedDate+".xlsx"
		}else if(jenisWarna == "merah"){
			i := 4;
			for j := idxAwal; j < len(witelSlice); j ++ {
				if((float64(witelSlice[j].TotalRetailStock - witelSlice[j].TotalRetail + witelSlice[j].OnDeliveryTotalRetail) < -(float64(witelSlice[j].TotalRetail) * 0.75)) || (float64(witelSlice[j].TotalPremiumStock - witelSlice[j].TotalPremium + witelSlice[j].OnDeliveryTotalPremium) < - (float64(witelSlice[j].TotalPremium) * 0.75))){
					fileExport.SetCellValue(activeSheetNameExport, "A"+strconv.Itoa(i), witelSlice[j].Regional)
					fileExport.SetCellValue(activeSheetNameExport, "B"+strconv.Itoa(i), witelSlice[j].LokasiWH)
					fileExport.SetCellValue(activeSheetNameExport, "C"+strconv.Itoa(i), witelSlice[j].MinimumQty)
					fileExport.SetCellValue(activeSheetNameExport, "D"+strconv.Itoa(i), witelSlice[j].RetailFH)
					fileExport.SetCellValue(activeSheetNameExport, "E"+strconv.Itoa(i), witelSlice[j].RetailHW)
					fileExport.SetCellValue(activeSheetNameExport, "F"+strconv.Itoa(i), witelSlice[j].RetailZTE)
					fileExport.SetCellValue(activeSheetNameExport, "G"+strconv.Itoa(i), witelSlice[j].RetailALU)
					fileExport.SetCellValue(activeSheetNameExport, "H"+strconv.Itoa(i), witelSlice[j].TotalRetail)
					fileExport.SetCellValue(activeSheetNameExport, "I"+strconv.Itoa(i), witelSlice[j].PremiumFH)
					fileExport.SetCellValue(activeSheetNameExport, "J"+strconv.Itoa(i), witelSlice[j].PremiumHW)
					fileExport.SetCellValue(activeSheetNameExport, "K"+strconv.Itoa(i), witelSlice[j].PremiumZTE)
					fileExport.SetCellValue(activeSheetNameExport, "L"+strconv.Itoa(i), witelSlice[j].TotalPremium)
					fileExport.SetCellValue(activeSheetNameExport, "M"+strconv.Itoa(i), witelSlice[j].RetailStockFiberhome)
					fileExport.SetCellValue(activeSheetNameExport, "N"+strconv.Itoa(i), witelSlice[j].RetailStockHuawei)
					fileExport.SetCellValue(activeSheetNameExport, "O"+strconv.Itoa(i), witelSlice[j].RetailStockZTE)
					fileExport.SetCellValue(activeSheetNameExport, "P"+strconv.Itoa(i), witelSlice[j].RetailStockNokia)
					fileExport.SetCellValue(activeSheetNameExport, "Q"+strconv.Itoa(i), witelSlice[j].TotalRetailStock)
					fileExport.SetCellValue(activeSheetNameExport, "R"+strconv.Itoa(i), witelSlice[j].PremiumStockFiberhome)
					fileExport.SetCellValue(activeSheetNameExport, "S"+strconv.Itoa(i), witelSlice[j].PremiumStockHuawei)
					fileExport.SetCellValue(activeSheetNameExport, "T"+strconv.Itoa(i), witelSlice[j].PremiumStockZTE)
					fileExport.SetCellValue(activeSheetNameExport, "U"+strconv.Itoa(i), witelSlice[j].TotalPremiumStock)
					fileExport.SetCellValue(activeSheetNameExport, "V"+strconv.Itoa(i), witelSlice[j].RetailStockFiberhome - witelSlice[j].RetailFH + witelSlice[j].OnDeliveryRetailFiberhome)
					fileExport.SetCellValue(activeSheetNameExport, "W"+strconv.Itoa(i), witelSlice[j].RetailStockHuawei- witelSlice[j].RetailHW + witelSlice[j].OnDeliveryRetailHuawei)
					fileExport.SetCellValue(activeSheetNameExport, "X"+strconv.Itoa(i), witelSlice[j].RetailStockZTE - witelSlice[j].RetailZTE + witelSlice[j].OnDeliveryRetailZTE)
					fileExport.SetCellValue(activeSheetNameExport, "Y"+strconv.Itoa(i), witelSlice[j].RetailStockNokia - witelSlice[j].RetailALU + witelSlice[j].OnDeliveryRetailNokia)
					fileExport.SetCellValue(activeSheetNameExport, "Z"+strconv.Itoa(i), witelSlice[j].TotalRetailStock - witelSlice[j].TotalRetail + witelSlice[j].OnDeliveryTotalRetail)

					fileExport.SetCellValue(activeSheetNameExport, "AA"+strconv.Itoa(i), witelSlice[j].PremiumStockFiberhome - witelSlice[j].PremiumFH + witelSlice[j].OnDeliveryPremiumFiberhome)
					fileExport.SetCellValue(activeSheetNameExport, "AB"+strconv.Itoa(i), witelSlice[j].PremiumStockHuawei - witelSlice[j].PremiumHW + witelSlice[j].OnDeliveryPremiumHuawei)
					fileExport.SetCellValue(activeSheetNameExport, "AC"+strconv.Itoa(i), witelSlice[j].PremiumStockZTE - witelSlice[j].PremiumZTE + witelSlice[j].OnDeliveryPremiumZTE)
					fileExport.SetCellValue(activeSheetNameExport, "AD"+strconv.Itoa(i), witelSlice[j].TotalPremiumStock - witelSlice[j].TotalPremium + witelSlice[j].OnDeliveryTotalPremium)

					fileExport.SetCellValue(activeSheetNameExport, "AE"+strconv.Itoa(i), witelSlice[j].BatasAtasRetailFH)
					fileExport.SetCellValue(activeSheetNameExport, "AF"+strconv.Itoa(i), witelSlice[j].BatasAtasRetailHW)
					fileExport.SetCellValue(activeSheetNameExport, "AG"+strconv.Itoa(i), witelSlice[j].BatasAtasRetailZTE)
					fileExport.SetCellValue(activeSheetNameExport, "AH"+strconv.Itoa(i), witelSlice[j].BatasAtasRetailALU)
					fileExport.SetCellValue(activeSheetNameExport, "AI"+strconv.Itoa(i), witelSlice[j].BatasAtasRetailFH + witelSlice[j].BatasAtasRetailHW + witelSlice[j].BatasAtasRetailZTE + witelSlice[j].BatasAtasRetailALU)

					fileExport.SetCellValue(activeSheetNameExport, "AJ"+strconv.Itoa(i), witelSlice[j].BatasAtasPremiumFH)
					fileExport.SetCellValue(activeSheetNameExport, "AK"+strconv.Itoa(i), witelSlice[j].BatasAtasPremiumHW)
					fileExport.SetCellValue(activeSheetNameExport, "AL"+strconv.Itoa(i), witelSlice[j].BatasAtasPremiumZTE)
					fileExport.SetCellValue(activeSheetNameExport, "AM"+strconv.Itoa(i), witelSlice[j].BatasAtasPremiumFH + witelSlice[j].BatasAtasPremiumHW + witelSlice[j].BatasAtasPremiumZTE)

					fileExport.SetCellValue(activeSheetNameExport, "AN"+strconv.Itoa(i), witelSlice[j].BatasBawahRetailFH)
					fileExport.SetCellValue(activeSheetNameExport, "AO"+strconv.Itoa(i), witelSlice[j].BatasBawahRetailHW)
					fileExport.SetCellValue(activeSheetNameExport, "AP"+strconv.Itoa(i), witelSlice[j].BatasBawahRetailZTE)
					fileExport.SetCellValue(activeSheetNameExport, "AQ"+strconv.Itoa(i), witelSlice[j].BatasBawahRetailALU)
					fileExport.SetCellValue(activeSheetNameExport, "AR"+strconv.Itoa(i), witelSlice[j].BatasBawahRetailFH + witelSlice[j].BatasBawahRetailHW + witelSlice[j].BatasBawahRetailZTE + witelSlice[j].BatasBawahRetailALU)

					fileExport.SetCellValue(activeSheetNameExport, "AS"+strconv.Itoa(i), witelSlice[j].BatasBawahPremiumFH)
					fileExport.SetCellValue(activeSheetNameExport, "AT"+strconv.Itoa(i), witelSlice[j].BatasBawahPremiumHW)
					fileExport.SetCellValue(activeSheetNameExport, "AU"+strconv.Itoa(i), witelSlice[j].BatasBawahPremiumZTE)
					fileExport.SetCellValue(activeSheetNameExport, "AV"+strconv.Itoa(i), witelSlice[j].BatasBawahPremiumFH + witelSlice[j].BatasBawahPremiumHW + witelSlice[j].BatasBawahPremiumZTE)

					fileExport.SetCellValue(activeSheetNameExport, "AW"+strconv.Itoa(i), witelSlice[j].QtyKirimRetailFH)
					fileExport.SetCellValue(activeSheetNameExport, "AX"+strconv.Itoa(i), witelSlice[j].QtyKirimRetailHW)
					fileExport.SetCellValue(activeSheetNameExport, "AY"+strconv.Itoa(i), witelSlice[j].QtyKirimRetailZTE)
					fileExport.SetCellValue(activeSheetNameExport, "AZ"+strconv.Itoa(i), witelSlice[j].QtyKirimRetailALU)
					fileExport.SetCellValue(activeSheetNameExport, "BA"+strconv.Itoa(i), witelSlice[j].QtyKirimRetailFH + witelSlice[j].QtyKirimRetailHW + witelSlice[j].QtyKirimRetailZTE + witelSlice[j].QtyKirimRetailALU)

					fileExport.SetCellValue(activeSheetNameExport, "BB"+strconv.Itoa(i), witelSlice[j].QtyKirimPremiumFH)
					fileExport.SetCellValue(activeSheetNameExport, "BC"+strconv.Itoa(i), witelSlice[j].QtyKirimPremiumHW)
					fileExport.SetCellValue(activeSheetNameExport, "BD"+strconv.Itoa(i), witelSlice[j].QtyKirimPremiumZTE)
					fileExport.SetCellValue(activeSheetNameExport, "BE"+strconv.Itoa(i), witelSlice[j].QtyKirimPremiumFH + witelSlice[j].QtyKirimPremiumHW + witelSlice[j].QtyKirimPremiumZTE)

					fileExport.SetCellValue(activeSheetNameExport, "BF"+strconv.Itoa(i), witelSlice[j].OnDeliveryRetailFiberhome)
					fileExport.SetCellValue(activeSheetNameExport, "BG"+strconv.Itoa(i), witelSlice[j].OnDeliveryRetailHuawei)
					fileExport.SetCellValue(activeSheetNameExport, "BH"+strconv.Itoa(i), witelSlice[j].OnDeliveryRetailZTE)
					fileExport.SetCellValue(activeSheetNameExport, "BI"+strconv.Itoa(i), witelSlice[j].OnDeliveryRetailNokia)
					fileExport.SetCellValue(activeSheetNameExport, "BJ"+strconv.Itoa(i), witelSlice[j].OnDeliveryTotalRetail)

					fileExport.SetCellValue(activeSheetNameExport, "BK"+strconv.Itoa(i), witelSlice[j].OnDeliveryPremiumFiberhome)
					fileExport.SetCellValue(activeSheetNameExport, "BL"+strconv.Itoa(i), witelSlice[j].OnDeliveryPremiumHuawei)
					fileExport.SetCellValue(activeSheetNameExport, "BM"+strconv.Itoa(i), witelSlice[j].OnDeliveryPremiumZTE)
					fileExport.SetCellValue(activeSheetNameExport, "BN"+strconv.Itoa(i), witelSlice[j].OnDeliveryTotalPremium)

					if(float64(witelSlice[j].TotalRetailStock - witelSlice[j].TotalRetail + witelSlice[j].OnDeliveryTotalRetail) < -(float64(witelSlice[j].TotalRetail) * 0.75)){
						if err := fileExport.SetCellStyle(activeSheetNameExport, "Z"+strconv.Itoa(i), "Z"+strconv.Itoa(i), styleRed); err != nil {
					        helper.PanicIfError(err)
					    }
					}

					if(float64(witelSlice[j].TotalPremiumStock - witelSlice[j].TotalPremium + witelSlice[j].OnDeliveryTotalPremium) < -(float64(witelSlice[j].TotalPremium) * 0.75)){
						if err := fileExport.SetCellStyle(activeSheetNameExport, "AD"+strconv.Itoa(i), "AD"+strconv.Itoa(i), styleRed); err != nil {
					        helper.PanicIfError(err)
					    }
					}
					i += 1;
				}
			}
			filename = "Hasil_Rekap_Stock_SCMT_Merah_"+formattedDate+".xlsx"
		}
	}

	// Set active sheet of the workbook
    fileExport.SetActiveSheet(activeSheetIndexExport)

    bytesBuffer, err := fileExport.WriteToBuffer()

	//Export to download
	return bytesBuffer.Bytes(), filename, err
}