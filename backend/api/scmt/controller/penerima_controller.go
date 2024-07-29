package controller

import(
	"net/http"
	"portofolio.com/service/scmt"
	"portofolio.com/api/helper"
	"portofolio.com/domain"
	// "portofolio.com/domain/scmt"
	// "github.com/gorilla/mux"
	// "fmt"
	// "encoding/json"
	// "strconv"
)

type PenerimaController interface{

}

type penerimaController struct{
	penerimaService service.PenerimaService
}

func NewPenerimaController(penerimaService service.PenerimaService) *penerimaController{
	return &penerimaController{penerimaService}
}

func (c *penerimaController) GetPengirimanONT(w http.ResponseWriter, r *http.Request){
	penerimaResponse := c.penerimaService.GetPengirimanONT()
	webResponse := web.WebResponse{
		Code : 200,
		Status : "OK",
		Data : penerimaResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}