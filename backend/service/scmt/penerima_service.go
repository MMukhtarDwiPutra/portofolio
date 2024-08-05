package service

import(
	"time"
	"portofolio.com/repository/scmt"
	"portofolio.com/domain/scmt"
	// "portofolio.com/domain/scmt"
	"portofolio.com/api/helper"
	"fmt"
	"github.com/xuri/excelize/v2"
	"strconv"
)

type PenerimaService interface{
	GetPengirimanONT() map[string]interface{}
	ExportPenerima(jenisExport string) ([]byte, string, error)
	DownloadAllSN(jenisExport string, jenisDownload string) ([]byte, string, error)
	DownloadTemplatePenerima() ([]byte, string, error)
	DownloadTemplateSerialNumber(jenisDelivery string) ([]byte, string, error)
	AddPenerima(penerima domain.PenerimaPost)
	DeletePenerimaById(id int)
	DownloadSerialNumber(id int) ([]byte, string, error)
	UploadPenerimaan()
	DeleteAllPenerima()
	EditIDOGDById(id string, data domain.PenerimaPost)
	GetFitur(namaFitur string) string
	EditOnDeliveryById(penerima domain.PenerimaPost, id string) string
	EditTanggalOnly(id string, penerima domain.PenerimaPost)
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

func (s *penerimaService) ExportPenerima(jenisExport string) ([]byte, string, error){
	var filename string
	var penerima []domain.PenerimaResponse

	templatePath := "template/template_export_pengiriman.xlsx"
    f, err := excelize.OpenFile(templatePath)
    helper.PanicIfError(err)

	// Get the current time
	currentTime := time.Now()

	// Format the time as YYYY-MM-DD
	formattedDate := currentTime.Format("2006-01-02")

	if(jenisExport == "All"){
		penerima = s.penerimaRepository.GetAllPenerimaExport();
		filename = "Rekap_Delivery_All_"+ formattedDate +".xlsx";
	}else if(jenisExport == "ONT"){
		penerima = s.penerimaRepository.GetAllDataONTExport();
		filename = "Rekap_Delivery_ONT_"+formattedDate+".xlsx";
	}else if(jenisExport == "STB"){
		penerima = s.penerimaRepository.GetAllDataSTBExport();
		filename = "Rekap_Delivery_STB_"+formattedDate+".xlsx";
	}else if(jenisExport == "AP"){
		penerima = s.penerimaRepository.GetAllDataAPExport();
		filename = "Rekap_Delivery_AP_"+formattedDate+".xlsx";
	}

	//Write to excel
	 // Create a new Excel file
    // f := excelize.NewFile()

    // Create a new sheet
    namaSheet := "Sheet1"
    indexSheet, err := f.NewSheet(namaSheet)
    helper.PanicIfError(err)

    // Set values in cells
    for i := range penerima{
	    f.SetCellValue(namaSheet, "A"+strconv.Itoa(i+3), penerima[i].Type)
	    f.SetCellValue(namaSheet, "B"+strconv.Itoa(i+3), penerima[i].Qty)
	    f.SetCellValue(namaSheet, "C"+strconv.Itoa(i+3), penerima[i].AlamatPengirim)
	    f.SetCellValue(namaSheet, "D"+strconv.Itoa(i+3), penerima[i].PICPengirim)
	    f.SetCellValue(namaSheet, "E"+strconv.Itoa(i+3), penerima[i].AlamatPenerima)
	    f.SetCellValue(namaSheet, "F"+strconv.Itoa(i+3), penerima[i].WarehousePenerima)	
	    f.SetCellValue(namaSheet, "G"+strconv.Itoa(i+3), penerima[i].Regional)
	    f.SetCellValue(namaSheet, "H"+strconv.Itoa(i+3), penerima[i].PICPenerima)
	    f.SetCellValue(namaSheet, "I"+strconv.Itoa(i+3), penerima[i].TanggalPengiriman)
	    f.SetCellValue(namaSheet, "J"+strconv.Itoa(i+3), penerima[i].TanggalSampai)
	    f.SetCellValue(namaSheet, "K"+strconv.Itoa(i+3), penerima[i].IDOGD)
	    f.SetCellValue(namaSheet, "L"+strconv.Itoa(i+3), penerima[i].SNMacBarcode)
	    f.SetCellValue(namaSheet, "M"+strconv.Itoa(i+3), penerima[i].Batch)
	    f.SetCellValue(namaSheet, "N"+strconv.Itoa(i+3), penerima[i].IDOGDTimeAdded)
	    f.SetCellValue(namaSheet, "O"+strconv.Itoa(i+3), penerima[i].SNTimeAdded)
	    f.SetCellValue(namaSheet, "P"+strconv.Itoa(i+3), penerima[i].TimeAdded)
    }

    // Set active sheet of the workbook
    f.SetActiveSheet(indexSheet)

    bytesBuffer, err := f.WriteToBuffer()

	//Export to download
	return bytesBuffer.Bytes(), filename, err
}

func (s *penerimaService) DownloadAllSN(jenisExport string, jenisDownload string) ([]byte, string, error){
	var filename string
	var penerima []domain.PenerimaResponse

	templatePath := "template/template_export_sn.xlsx"
    fileExport, err := excelize.OpenFile(templatePath)
    helper.PanicIfError(err)

    // Get the active sheet index
    activeSheetIndexExport := fileExport.GetActiveSheetIndex()

    // Get the name of the active sheet using the index
    activeSheetNameExport := fileExport.GetSheetName(activeSheetIndexExport)

    if(jenisDownload == "exist"){
    	if(jenisExport == "stb"){
    		penerima = s.penerimaRepository.GetAllSNSTBExist();
    	}else if(jenisExport == "ont"){
    		penerima = s.penerimaRepository.GetAllSNONTExist();
    	}
    }else if(jenisDownload == "all"){
    	if(jenisExport == "stb"){
    		penerima = s.penerimaRepository.GetAllSNSTB();
    	}else if(jenisExport == "ont"){
    		penerima = s.penerimaRepository.GetAllSNONT();
    	}
    }

    // if(session["jenis_akun"] == "treg")
    // filter penerima.regional by session(asal)

    i := 2
    for j := range penerima{
    	sn_mac_barcode, batch := s.penerimaRepository.GetSNBatchById(penerima[j].ID)
    	spreadsheetImport, err := excelize.OpenFile("Uploaded SN/"+sn_mac_barcode)
    	helper.PanicIfError(err)

    	// Get the active sheet index
	    activeSheetIndex := spreadsheetImport.GetActiveSheetIndex()

	    // Get the name of the active sheet using the index
	    activeSheetName := spreadsheetImport.GetSheetName(activeSheetIndex)

	    rows, err := spreadsheetImport.GetRows(activeSheetName)
	    helper.PanicIfError(err)

	    // Get the maximum number of rows
	    maxRows := len(rows)
    	for k := 2; k < maxRows+1; k++ {
    		serialNumber, err := spreadsheetImport.GetCellValue(activeSheetName, "A"+strconv.Itoa(k))
    		helper.PanicIfError(err)
    		macAddress, err := spreadsheetImport.GetCellValue(activeSheetName,  "B"+strconv.Itoa(k))
    		helper.PanicIfError(err)

    		fileExport.SetCellValue(activeSheetNameExport, "A"+strconv.Itoa(i), serialNumber)
		    fileExport.SetCellValue(activeSheetNameExport, "B"+strconv.Itoa(i), macAddress)
		    fileExport.SetCellValue(activeSheetNameExport, "C"+strconv.Itoa(i), penerima[j].WarehousePenerima)
		    fileExport.SetCellValue(activeSheetNameExport, "D"+strconv.Itoa(i), penerima[j].Type)
		    fileExport.SetCellValue(activeSheetNameExport, "E"+strconv.Itoa(i), batch)

		    i+=1
    	}
    }

    // Set active sheet of the workbook
    fileExport.SetActiveSheet(activeSheetIndexExport)

    bytesBuffer, err := fileExport.WriteToBuffer()

	//Export to download
	return bytesBuffer.Bytes(), filename, err
}

func (s *penerimaService) DownloadTemplatePenerima() ([]byte, string, error){
	filename := "template_data_penerima.xlsx"
	templatePath := "template/template_penerima.xlsx"
    fileExport, err := excelize.OpenFile(templatePath)
    helper.PanicIfError(err)

    bytesBuffer, err := fileExport.WriteToBuffer()

	//Export to download
	return bytesBuffer.Bytes(), filename, err
}

func (s *penerimaService) DownloadTemplateSerialNumber(jenisDelivery string) ([]byte, string, error){
	filename := "template_serial_number.xlsx"
	var templatePath string

	if(jenisDelivery == "ont"){
		templatePath = "template/template_upload_serial_number_ont.xlsx"	
	}else if(jenisDelivery == "stb"){
		templatePath = "template/template_upload_serial_number_stb.xlsx"
	}
	
    
    fileExport, err := excelize.OpenFile(templatePath)
    helper.PanicIfError(err)

    bytesBuffer, err := fileExport.WriteToBuffer()

	//Export to download
	return bytesBuffer.Bytes(), filename, err
}

func (s *penerimaService) AddPenerima(penerima domain.PenerimaPost){
	s.penerimaRepository.AddPenerima(penerima)
}

func (s *penerimaService) DeletePenerimaById(id int){
	s.penerimaRepository.DeletePenerimaById(id);
}

func (s *penerimaService) DownloadSerialNumber(id int) ([]byte, string, error){
	fileName := s.penerimaRepository.GetSNById(id);
	var filePath string

	if(fileName != "SN tidak ada!"){
		filePath = "Uploaded SN/"+fileName
	}else{
		filePath = ""
	}

	fileExport, err := excelize.OpenFile(filePath)
    helper.PanicIfError(err)

    bytesBuffer, err := fileExport.WriteToBuffer()
    //Export to download
	return bytesBuffer.Bytes(), fileName, err
}

// Generic function to check if a value exists in a slice
func Contains[T comparable](slice []T, value T) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

func (s *penerimaService) EditOnDeliveryById(penerima domain.PenerimaPost, id string) string{
	spreadsheetSN, err := excelize.OpenFile("Uploaded SN/"+penerima.SNMacBarcode)
	helper.PanicIfError(err)

	// Get the active sheet index
    activeSheetIndex := spreadsheetSN.GetActiveSheetIndex()

    // Get the name of the active sheet using the index
    activeSheetName := spreadsheetSN.GetSheetName(activeSheetIndex)

    rows, err := spreadsheetSN.GetRows(activeSheetName)
    helper.PanicIfError(err)

    // Get the maximum number of rows
    maxRows := len(rows)
    tmpMaxRow := 0
    for i:= 2; i < maxRows+1; i++{
    	cellAData, err := spreadsheetSN.GetCellValue(activeSheetName, "A"+strconv.Itoa(i))
    	helper.PanicIfError(err)
    	if(cellAData != ""){
    		tmpMaxRow += 1
    	}
    }

    maxRows = tmpMaxRow
    idInt, err := strconv.Atoi(id)
    helper.PanicIfError(err)
    penerimaTmp := s.penerimaRepository.GetDataById(idInt)
    if(maxRows != penerimaTmp.Qty){
    	return "File gagal diupload dikarenakan data serial number tidak sama dengan quantity seharusnya!"
    }else{
    	var tmpSN []string
    	for i:= 2; i < maxRows+2; i++{
    		cellAData, err := spreadsheetSN.GetCellValue(activeSheetName, "A"+strconv.Itoa(i))
    		helper.PanicIfError(err)
    		cellBData, err := spreadsheetSN.GetCellValue(activeSheetName, "B"+strconv.Itoa(i))
    		helper.PanicIfError(err)

    		if(cellAData == "" || cellBData == ""){
    			return "File gagal diupload dikarenakan data serial number tidak lengkap dengan quantity seharusnya!"
    		}

    		if(!Contains(tmpSN, cellAData)){
    			tmpSN = append(tmpSN, cellAData)
    		}else{
    			return "File gagal diupload dikarenakan ada SN yang sama (duplicate)!"
    		}
    	}

    	//insert add notif function

    	s.penerimaRepository.EditTanggalPenerimaanById(id, penerima)
    }
    return "Data berhasil diedit!"
}

func (s *penerimaService) GetFitur(namaFitur string) string{
	return s.fiturRepository.GetFitur(namaFitur);
}

func (s *penerimaService) UploadPenerimaan(){
	templatePath := "template/uploaded_penerima.xlsx"
    spreadsheetImport, err := excelize.OpenFile(templatePath)
    helper.PanicIfError(err)

    // Get the active sheet index
    activeSheetIndexUploaded := spreadsheetImport.GetActiveSheetIndex()

    // Get the name of the active sheet using the index
    activeSheetName := spreadsheetImport.GetSheetName(activeSheetIndexUploaded)

    rows, err := spreadsheetImport.GetRows(activeSheetName)
    helper.PanicIfError(err)

    var penerimaUploads []domain.PenerimaPost

    // Get the maximum number of rows
    maxRows := len(rows)
    for i := 3; i < maxRows+1; i++{
    	var penerimaUpload domain.PenerimaPost

    	penerimaUpload.Type, err = spreadsheetImport.GetCellValue(activeSheetName, "A"+strconv.Itoa(i))
    	penerimaUpload.Qty, err = spreadsheetImport.GetCellValue(activeSheetName, "B"+strconv.Itoa(i))
    	penerimaUpload.AlamatPengirim, err = spreadsheetImport.GetCellValue(activeSheetName, "C"+strconv.Itoa(i))
    	penerimaUpload.PICPengirim, err = spreadsheetImport.GetCellValue(activeSheetName, "D"+strconv.Itoa(i))
    	penerimaUpload.AlamatPenerima, err = spreadsheetImport.GetCellValue(activeSheetName, "E"+strconv.Itoa(i))
    	penerimaUpload.WarehousePenerima, err = spreadsheetImport.GetCellValue(activeSheetName, "F"+strconv.Itoa(i))
    	penerimaUpload.PICPenerima, err = spreadsheetImport.GetCellValue(activeSheetName, "G"+strconv.Itoa(i))
    	penerimaUpload.TanggalPengiriman, err = spreadsheetImport.GetCellValue(activeSheetName, "H"+strconv.Itoa(i))
    	penerimaUpload.TanggalSampai, err = spreadsheetImport.GetCellValue(activeSheetName, "I"+strconv.Itoa(i))
    	penerimaUpload.IDOGD, err = spreadsheetImport.GetCellValue(activeSheetName, "J"+strconv.Itoa(i))
    	penerimaUpload.SNMacBarcode, err = spreadsheetImport.GetCellValue(activeSheetName, "K"+strconv.Itoa(i))
    	penerimaUpload.Batch, err = spreadsheetImport.GetCellValue(activeSheetName, "L"+strconv.Itoa(i))
    	penerimaUpload.IDOGDTimeAdded, err = spreadsheetImport.GetCellValue(activeSheetName, "M"+strconv.Itoa(i))
    	penerimaUpload.SNTimeAdded, err = spreadsheetImport.GetCellValue(activeSheetName, "N"+strconv.Itoa(i))
    	penerimaUpload.TimeAdded, err = spreadsheetImport.GetCellValue(activeSheetName, "O"+strconv.Itoa(i))

    	if(penerimaUpload.IDOGD != ""){
    		if(penerimaUpload.IDOGDTimeAdded == ""){
    			// Get current date and time
				now := time.Now()

				// Format the date and time
				timeNow := now.Format("2006-01-02 15:04:05")

    			penerimaUpload.IDOGDTimeAdded = timeNow;
    		}
    	}

    	penerimaUploads = append(penerimaUploads, penerimaUpload)
    }

   	s.penerimaRepository.AddPenerimaBulk(penerimaUploads)

    // Ensure the file is closed properly
    defer func() {
        if err := spreadsheetImport.Close(); err != nil {
            fmt.Println("Error closing file:", err)
        }
    }()
}

func (s *penerimaService) DeleteAllPenerima(){
	s.penerimaRepository.DeleteAllPenerima()
}

func (s *penerimaService) EditIDOGDById(id string, penerima domain.PenerimaPost){
	s.penerimaRepository.EditIDOGDById(id, penerima)
}

func (s *penerimaService) EditTanggalOnly(id string, penerima domain.PenerimaPost){
	s.penerimaRepository.EditTanggalOnly(id, penerima)
}