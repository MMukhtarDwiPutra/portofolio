package controller

import(
	"net/http"
	"portofolio.com/service/scmt"
	"portofolio.com/api/helper"
	"portofolio.com/domain"
	"portofolio.com/domain/scmt"
	"github.com/gorilla/mux"
	"os"
	"io"
	// "fmt"
	// "encoding/json"
	// "strconv"
)

type DataSCMTController interface{
	GetAllDataTmp(w http.ResponseWriter, r *http.Request)
	InsertDataTmp(w http.ResponseWriter, r *http.Request)
	CountRetailPerWitel(w http.ResponseWriter, r *http.Request)
	// GetDataTMPById(w http.ResponseWriter, r *http.Request)
}

type dataTmpController struct{
	dataTmpService service.DataTmpService
}

func NewDataTmpController(dataTmpService service.DataTmpService) *dataTmpController{
	return &dataTmpController{dataTmpService}
}

func (c *dataTmpController) GetAllDataTmp(w http.ResponseWriter, r *http.Request){
	dataTmpResponse := c.dataTmpService.GetAllDataTmp()
	webResponse := web.WebResponse{
		Code : 200,
		Status : "OK",
		Data : dataTmpResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (c *dataTmpController) InsertDataTmp(w http.ResponseWriter, r *http.Request){
	dataTmp := domain.DataTmp{}
	helper.ReadFromRequestBody(r, &dataTmp)

	c.dataTmpService.InsertDataTmp(dataTmp)
	dataTmpResponse := c.dataTmpService.GetLastDataTmp()
	webResponse := web.WebResponse{
		Code : 200,
		Status : "OK",
		Data : dataTmpResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (c *dataTmpController) CountRetailPerWitel(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	merk, _ := params["merk"]

	countResponse := c.dataTmpService.CountRetailPerWitel(merk)

	webResponse := web.WebResponse{
		Code : 200,
		Status : "OK",
		Data : countResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (c *dataTmpController) CountPremiumPerWitel(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	merk, _ := params["merk"]

	countResponse := c.dataTmpService.CountPremiumPerWitel(merk)

	webResponse := web.WebResponse{
		Code : 200,
		Status : "OK",
		Data : countResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (c *dataTmpController) CountSTBPerWitel(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	merk, _ := params["merk"]

	countResponse := c.dataTmpService.CountSTBPerWitel(merk)

	webResponse := web.WebResponse{
		Code : 200,
		Status : "OK",
		Data : countResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (c *dataTmpController) CountAPPerWitel(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	merk, _ := params["merk"]

	countResponse := c.dataTmpService.CountAPPerWitel(merk)

	webResponse := web.WebResponse{
		Code : 200,
		Status : "OK",
		Data : countResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}


func (c *dataTmpController) GetRekapDeliveryTREG(w http.ResponseWriter, r *http.Request){
	data := c.dataTmpService.RekapDeliveryTREG();

	webResponse := web.WebResponse{
		Code : 200,
		Status : "OK",
		Data : data,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (c *dataTmpController) GetRekapDeliveryWitel(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	lokasiWH, _ := params["lokasi_wh"]
	data := c.dataTmpService.RekapDeliveryWitel(lokasiWH);

	webResponse := web.WebResponse{
		Code : 200,
		Status : "OK",
		Data : data,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (c *dataTmpController) ExportDataTmp(w http.ResponseWriter, r *http.Request){
	fileBytes, fileName, err := c.dataTmpService.GetExportDataTmp()
	helper.PanicIfError(err)

	helper.DownloadHandler(w, r, fileName, fileBytes)
}

func (c *dataTmpController) ExportMinimumStockDatabase(w http.ResponseWriter, r *http.Request){
	fileBytes, fileName, err := c.dataTmpService.GetExportMinimumStockDatabase()
	helper.PanicIfError(err)

	helper.DownloadHandler(w, r, fileName, fileBytes)
}

func (c *dataTmpController) DownloadTemplateMinimumStock(w http.ResponseWriter, r *http.Request){
	fileBytes, fileName, err := c.dataTmpService.DownloadTemplateMinimumStock()
	helper.PanicIfError(err)

	helper.DownloadHandler(w, r, fileName, fileBytes)
}

func (c *dataTmpController) DownloadTemplateDataTmp(w http.ResponseWriter, r *http.Request){
	fileBytes, fileName, err := c.dataTmpService.DownloadTemplateDataTmp()
	helper.PanicIfError(err)

	helper.DownloadHandler(w, r, fileName, fileBytes)
}

func (c *dataTmpController) UploadDataTmp(w http.ResponseWriter, r *http.Request){
	// Retrieve the file from form data
	file, _, err := r.FormFile("file")
	helper.PanicIfError(err)
	defer file.Close()

	// Create a destination file
	destPath := "template/uploaded_data_tmp.xlsx"
	destFile, err := os.Create(destPath)
	helper.PanicIfError(err)

	// Copy the uploaded file data to the destination file
	_, err = io.Copy(destFile, file)
	helper.PanicIfError(err)

	c.dataTmpService.UploadDataTmp()

	// Attempt to delete the file
	defer func() {
		destFile.Close()
		
		// Path to the file you want to delete
		filePath := "template/uploaded_data_tmp.xlsx"

		// Attempt to delete the file
		err = os.Remove(filePath)
		helper.PanicIfError(err)
	}()

	webResponse := web.WebResponse{
		Code : 200,
		Status : "OK",
		Data : map[string]string{
            "message": "Upload file sukses!",
        },
	}

	defer helper.WriteToResponseBody(w, webResponse)
}

func (c *dataTmpController) ExportDataTmpRekapPage(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	jenisWarna := params["jenis_warna"]
	jenisExport := params["jenis_export"]

	fileBytes, fileName, err := c.dataTmpService.ExportDataTmp(jenisWarna, jenisExport)
	helper.PanicIfError(err)

	helper.DownloadHandler(w, r, fileName, fileBytes)
}
// testing insert
// {
// 	"region": "region testing",
// 	"lokasi_wh" : "lokasi test",
// 	"status" : "status tes",
// 	"jumlah" : 0,
// 	"deskripsi" : "deskripsi tes"
// }