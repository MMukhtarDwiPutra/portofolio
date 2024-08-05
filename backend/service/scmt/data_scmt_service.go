package service

import(
	"portofolio.com/repository/scmt"
	"portofolio.com/domain/scmt"
	"portofolio.com/api/helper"
	"strings"
	"fmt"
	"reflect"
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
						witelSlice[i].QtyKirimRetailALU += roundToNearest(result, 10)
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
						witelSlice[i].QtyKirimPremiumFH += roundToNearest(result, 10)
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
							witelSlice[i].QtyKirimRetailALU += roundToNearest(result, 10)
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
							witelSlice[i].QtyKirimPremiumFH += roundToNearest(result, 10)
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
					dataSO[j].QtyKirimRetailALU += roundToNearest(result, 10)
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
					dataSO[j].QtyKirimPremiumFH += roundToNearest(result, 10)
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
							sumALU += roundToNearest(result, 10)
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
					sumAllZTE += roundToNearest(result, 20)
				}

				if(dataSO.RetailStockHuawei < dataSO.BatasBawahRetailHW){
					result = dataSO.BatasAtasRetailHW - dataSO.RetailStockHuawei
					sumAllHW += roundToNearest(result, 20)
				}

				if(dataSO.RetailStockFiberhome < dataSO.BatasBawahRetailFH){
					result = dataSO.BatasAtasRetailFH - dataSO.RetailStockFiberhome
					sumAllFH += roundToNearest(result, 20)
				}

				if(dataSO.RetailStockNokia < dataSO.BatasBawahRetailALU){
					result = dataSO.BatasAtasRetailALU - dataSO.RetailStockNokia
					sumAllALU += roundToNearest(result, 20)
				}

				if(dataSO.PremiumStockZTE < dataSO.BatasBawahPremiumZTE){
					result = dataSO.BatasAtasPremiumZTE - dataSO.PremiumStockZTE
					sumAllPZTE += roundToNearest(result, 20)
				}

				if(dataSO.PremiumStockHuawei < dataSO.BatasBawahPremiumHW){
					result = dataSO.BatasAtasPremiumHW - dataSO.PremiumStockHuawei
					sumAllPHW += roundToNearest(result, 20)
				}

				if(dataSO.PremiumStockFiberhome < dataSO.BatasBawahPremiumFH){
					result = dataSO.BatasAtasPremiumFH - dataSO.PremiumStockFiberhome
					sumAllPFH += roundToNearest(result, 20)
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
				witelSlice[i].QtyKirimRetailFH = roundToNearest(result, 5)
			}

			if(witelSlice[i].RetailStockHuawei < witelSlice[i].BatasBawahRetailHW){
				result = witelSlice[i].BatasAtasRetailHW - witelSlice[i].RetailStockHuawei
				witelSlice[i].QtyKirimRetailHW = roundToNearest(result, 5)
			}

			if(witelSlice[i].RetailStockZTE < witelSlice[i].BatasBawahRetailZTE){
				result = witelSlice[i].BatasAtasRetailZTE - witelSlice[i].RetailStockZTE
				witelSlice[i].QtyKirimRetailZTE = roundToNearest(result, 5)
			}

			if(witelSlice[i].RetailStockNokia < witelSlice[i].BatasBawahRetailALU){
				result = witelSlice[i].BatasAtasRetailALU - witelSlice[i].RetailStockNokia
				witelSlice[i].QtyKirimRetailALU = roundToNearest(result, 5)
			}

			if(witelSlice[i].PremiumStockFiberhome < witelSlice[i].BatasBawahPremiumFH){
				result = witelSlice[i].BatasAtasPremiumFH - witelSlice[i].PremiumStockFiberhome
				witelSlice[i].QtyKirimPremiumFH = roundToNearest(result, 5)
			}

			if(witelSlice[i].PremiumStockHuawei < witelSlice[i].BatasBawahPremiumHW){
				result = witelSlice[i].BatasAtasPremiumHW - witelSlice[i].PremiumStockHuawei
				witelSlice[i].QtyKirimPremiumHW = roundToNearest(result, 5)
			}

			if(witelSlice[i].PremiumStockZTE < witelSlice[i].BatasBawahPremiumZTE){
				result = witelSlice[i].BatasAtasPremiumZTE - witelSlice[i].PremiumStockZTE
				witelSlice[i].QtyKirimPremiumZTE = roundToNearest(result, 5)
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
