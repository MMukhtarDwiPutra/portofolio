package controller

import(
	"net/http"
	"portofolio.com/service/scmt"
	"portofolio.com/api/helper"
	"portofolio.com/domain"
	"portofolio.com/domain/scmt"
	"github.com/gorilla/mux"
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

func NewDataTmpController(service service.DataTmpService) *dataTmpController{
	return &dataTmpController{service}
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

	fmt.Println(webResponse)
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
	fmt.Println(webResponse)

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
	fmt.Println(webResponse)

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


// testing insert
// {
// 	"region": "region testing",
// 	"lokasi_wh" : "lokasi test",
// 	"status" : "status tes",
// 	"jumlah" : 0,
// 	"deskripsi" : "deskripsi tes"
// }