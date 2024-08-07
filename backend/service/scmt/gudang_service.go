package service

import(
	"portofolio.com/repository/scmt"
	"portofolio.com/domain/scmt"
	"math"
	"github.com/xuri/excelize/v2"
	"strconv"
	"portofolio.com/api/helper"
)

type GudangService interface{
	UploadNewGudang()
}

type gudangService struct{
	gudangRepository repository.GudangRepository
}

func NewGudangService(gudangRepository repository.GudangRepository) *gudangService{
	return &gudangService{gudangRepository}
}

func (s *gudangService) UploadNewGudang(){
	s.gudangRepository.DeleteAllDataGudang()
	templatePath := "template/uploaded_gudang.xlsx"
    spreadsheetImport, err := excelize.OpenFile(templatePath)
    helper.PanicIfError(err)

    // Get the active sheet index
    activeSheetIndexUploaded := spreadsheetImport.GetActiveSheetIndex()

    // Get the name of the active sheet using the index
    activeSheetName := spreadsheetImport.GetSheetName(activeSheetIndexUploaded)

    rows, err := spreadsheetImport.GetRows(activeSheetName)
    helper.PanicIfError(err)

    var gudang []domain.Gudang

    // Get the maximum number of rows
    maxRows := len(rows)
    for i := 4; i < maxRows+1; i++{
    	var gudangUpload domain.Gudang

    	gudangUpload.Regional, err = spreadsheetImport.GetCellValue(activeSheetName, "A"+strconv.Itoa(i))
    	if(gudangUpload.Regional == ""){
			break;
		}
    	gudangUpload.Witel, err = spreadsheetImport.GetCellValue(activeSheetName, "B"+strconv.Itoa(i))
    	gudangUpload.LokasiWH, err = spreadsheetImport.GetCellValue(activeSheetName, "C"+strconv.Itoa(i))
    	gudangUpload.Lokasi, err = spreadsheetImport.GetCellValue(activeSheetName, "D"+strconv.Itoa(i))
    	gudangUpload.Wilayah, err = spreadsheetImport.GetCellValue(activeSheetName, "E"+strconv.Itoa(i))
    	gudangUpload.MinimumQty, err = spreadsheetImport.GetCellValue(activeSheetName, "F"+strconv.Itoa(i))
    	gudangUpload.RetailFH, err = spreadsheetImport.GetCellValue(activeSheetName, "G"+strconv.Itoa(i))
    	gudangUpload.RetailHW, err = spreadsheetImport.GetCellValue(activeSheetName, "H"+strconv.Itoa(i))
    	gudangUpload.RetailZTE, err = spreadsheetImport.GetCellValue(activeSheetName, "I"+strconv.Itoa(i))
    	gudangUpload.RetailALU, err = spreadsheetImport.GetCellValue(activeSheetName, "J"+strconv.Itoa(i))
    	gudangUpload.PremiumFH, err = spreadsheetImport.GetCellValue(activeSheetName, "K"+strconv.Itoa(i))
    	gudangUpload.PremiumHW, err = spreadsheetImport.GetCellValue(activeSheetName, "L"+strconv.Itoa(i))
    	gudangUpload.PremiumZTE, err = spreadsheetImport.GetCellValue(activeSheetName, "M"+strconv.Itoa(i))
    	gudangUpload.STBZTE, err = spreadsheetImport.GetCellValue(activeSheetName, "F"+strconv.Itoa(i))
    	stbZTE, err := strconv.Atoi(gudangUpload.STBZTE)
    	helper.PanicIfError(err)
    	gudangUpload.STBZTE = strconv.Itoa(int(math.Ceil(float64(stbZTE) * 32 / 100)))
    	gudangUpload.APCisco, err = spreadsheetImport.GetCellValue(activeSheetName, "O"+strconv.Itoa(i))
    	gudangUpload.APHuawei, err = spreadsheetImport.GetCellValue(activeSheetName, "P"+strconv.Itoa(i))

    	gudang = append(gudang, gudangUpload)
    }

	s.gudangRepository.UploadGudangBulk(gudang)

	// Ensure the file is closed properly
    defer func() {
        if err := spreadsheetImport.Close(); err != nil {
            helper.PanicIfError(err)
        }
    }()
}