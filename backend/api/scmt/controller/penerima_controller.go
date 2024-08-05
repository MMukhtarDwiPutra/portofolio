package controller

import(
	"net/http"
	"portofolio.com/service/scmt"
	"portofolio.com/api/helper"
	"portofolio.com/domain"
	"portofolio.com/domain/scmt"
	"github.com/gorilla/mux"
	"fmt"
	"encoding/json"
	"strconv"
	"os"
	"io"
	"time"
	// "io/ioutil"
    // "log"
)

type PenerimaController interface{
	GetPengirimanONT(w http.ResponseWriter, r *http.Request)
	ExportAllPenerima(w http.ResponseWriter, r *http.Request)
	DownloadAllSNONT(w http.ResponseWriter, r *http.Request)
	DownloadAllSNONTExist(w http.ResponseWriter, r *http.Request)
	DownloadTemplatePenerima(w http.ResponseWriter, r *http.Request)
	DownloadTemplateSerialNumberONT(w http.ResponseWriter, r *http.Request)
	AddPenerima(w http.ResponseWriter, r *http.Request)
	DeletePenerimaById(w http.ResponseWriter, r *http.Request)
	DownloadSerialNumber(w http.ResponseWriter, r *http.Request)
	UploadPenerimaan(w http.ResponseWriter, r *http.Request)
	DeleteAllPenerima(w http.ResponseWriter, r *http.Request)
	EditOnDeliveryById(w http.ResponseWriter, r *http.Request)
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

func (c *penerimaController) ExportAllPenerima(w http.ResponseWriter, r *http.Request){
	fileBytes, fileName, err := c.penerimaService.ExportPenerima("All")
	helper.PanicIfError(err)

	helper.DownloadHandler(w, r, fileName, fileBytes)
}

func (c *penerimaController) DownloadAllSNONT(w http.ResponseWriter, r *http.Request){
	fileBytes, fileName, err := c.penerimaService.DownloadAllSN("ont", "all")
	helper.PanicIfError(err)

	helper.DownloadHandler(w, r, fileName, fileBytes)
}

func (c *penerimaController) DownloadAllSNONTExist(w http.ResponseWriter, r *http.Request){
	fileBytes, fileName, err := c.penerimaService.DownloadAllSN("ont", "exist")
	helper.PanicIfError(err)

	helper.DownloadHandler(w, r, fileName, fileBytes)
}

func (c *penerimaController) DownloadTemplatePenerima(w http.ResponseWriter, r *http.Request){
	fileBytes, fileName, err := c.penerimaService.DownloadTemplatePenerima()
	helper.PanicIfError(err)

	helper.DownloadHandler(w, r, fileName, fileBytes)
}

func (c *penerimaController) DownloadTemplateSerialNumberONT(w http.ResponseWriter, r *http.Request){
	fileBytes, fileName, err := c.penerimaService.DownloadTemplateSerialNumber("ont")
	helper.PanicIfError(err)

	helper.DownloadHandler(w, r, fileName, fileBytes)
}

func (c *penerimaController) AddPenerima(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodPost {
        http.Error(w, "Only POST requests are allowed", http.StatusMethodNotAllowed)
        return
    }

    var penerimaTmp domain.PenerimaPost

    decoder := json.NewDecoder(r.Body)
    err := decoder.Decode(&penerimaTmp)
    helper.PanicIfError(err)

    c.penerimaService.AddPenerima(penerimaTmp)

    webResponse := web.WebResponse{
		Code : 200,
		Status : "OK",
		Data : map[string]string{
            "message": "Data penerima berhasil ditambah!",
        },
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (c *penerimaController) DeletePenerimaById(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	helper.PanicIfError(err)

	c.penerimaService.DeletePenerimaById(id)

    webResponse := web.WebResponse{
		Code : 200,
		Status : "OK",
		Data : map[string]string{
            "message": "Data penerima berhasil dihapus!",
        },
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (c *penerimaController) DownloadSerialNumber(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	helper.PanicIfError(err)

	fileBytes, fileName, err := c.penerimaService.DownloadSerialNumber(id)
	helper.PanicIfError(err)

	if(fileName == "SN tidak ada!"){
		webResponse := web.WebResponse{
			Code : 200,
			Status : "OK",
			Data : map[string]string{
	            "message": "File serial number tidak ditemukan!",
	        },
		}

		helper.WriteToResponseBody(w, webResponse)
	}else{
		helper.DownloadHandler(w, r, fileName, fileBytes)
	}
}

func (c *penerimaController) UploadPenerimaan(w http.ResponseWriter, r *http.Request){
	// Retrieve the file from form data
	file, _, err := r.FormFile("file")
	helper.PanicIfError(err)
	defer file.Close()

	// Create a destination file
	destPath := "template/uploaded_penerima.xlsx"
	destFile, err := os.Create(destPath)
	helper.PanicIfError(err)

	// Copy the uploaded file data to the destination file
	_, err = io.Copy(destFile, file)
	helper.PanicIfError(err)

	c.penerimaService.UploadPenerimaan()

	// Attempt to delete the file
	defer func() {
		destFile.Close()
		
		// Path to the file you want to delete
		filePath := "template/uploaded_penerima.xlsx"

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

func (c *penerimaController) DeleteAllPenerima(w http.ResponseWriter, r *http.Request){
	c.penerimaService.DeleteAllPenerima()

    webResponse := web.WebResponse{
		Code : 200,
		Status : "OK",
		Data : map[string]string{
            "message": "Semua data penerima berhasil dihapus!",
        },
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (c *penerimaController) EditOnDeliveryById(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	jenisDelivery := params["jenis_delivery"]
	id := params["id"]

	var namaFitur string
	if(jenisDelivery == "ont"){
		namaFitur = "report_delivery_ont"
	}else if(jenisDelivery == "stb"){
		namaFitur = "report_delivery_stb"
	}

	var penerima domain.PenerimaPost
    
    // body, err := ioutil.ReadAll(r.Body)
    // if err != nil {
    //     log.Println("Error reading body:", err)
    //     http.Error(w, err.Error(), http.StatusBadRequest)
    //     return
    // }
    // log.Println("Request Body:", string(body))

    // err = json.NewDecoder(r.Body).Decode(&penerima)
    // if err != nil {
    //     log.Println("Error decoding JSON:", err)
    //     http.Error(w, err.Error(), http.StatusBadRequest)
    //     return
    // }

    // err = json.NewDecoder(r.Body).Decode(&penerima)
    // helper.PanicIfError(err)

    penerima.TanggalPengiriman = r.FormValue("tanggal_pengiriman")
    penerima.TanggalSampai = r.FormValue("tanggal_sampai")
    penerima.IDOGD = r.FormValue("ido_gd")

	var message string
	statusTableDisable := c.penerimaService.GetFitur(namaFitur)
	fmt.Println(statusTableDisable)
	if(statusTableDisable == "OFF"){
	// 	fmt.Println("statusTableDisable off")
		file, handler, err := r.FormFile("sn_mac_barcode_file")

		if(err == nil){
			fmt.Println("Ada file")
			now := time.Now()
			timeNow := now.Format("2006-01-02 15_04_05")
	   	 	fileName := timeNow+" "+handler.Filename
	   	 	penerima.SNMacBarcode = fileName
	   	 	
	   	 	// Create a file on the server where the uploaded file will be stored.
		    dst, err := os.Create("Uploaded SN/" + fileName)
		    helper.PanicIfError(err)
		    if err != nil {
		        http.Error(w, "Unable to create the file on server", http.StatusInternalServerError)
		        return
		    }

		    // Copy the uploaded file data to the newly created file.
		    if _, err := io.Copy(dst, file); err != nil {
		        http.Error(w, "Error saving the file", http.StatusInternalServerError)
		        return
		    }

			message = c.penerimaService.EditOnDeliveryById(penerima, id)
			if(message != "Data berhasil diedit!"){
				// Attempt to delete the file
				defer func() {
		    		dst.Close()
					
					// Path to the file you want to delete
					filePath := "Uploaded SN/" + fileName

					// Attempt to delete the file
					err = os.Remove(filePath)
					helper.PanicIfError(err)
				}()
			}
		}else if(penerima.IDOGD != ""){
			// penerimaTmp := s.penerimaService.GetDataById(id)

			//Add Notif Function here

			c.penerimaService.EditIDOGDById(id, penerima)
			message = "Data berhasil diedit!"
		}else{
			c.penerimaService.EditTanggalOnly(id, penerima)
			message = "Data berhasil diedit!"
		}
	}else{
		message = "Data tidak berhasil diubah dikarenakan sedang ada maintance data oleh admin! Silahkan coba beberapa saat lagi!"
	}

	webResponse := web.WebResponse{
		Code : 200,
		Status : "OK",
		Data : map[string]string{
	        "message": message,
    	},
    }

	helper.WriteToResponseBody(w, webResponse)
}

//Penerima testing
// {
//     "type" : "Nokia",
//     "qty" : 20,
//     "pic_pengirim" : "putra",
//     "alamat_pengirim" : "Jambi",
//     "pic_penerima" : "Fajar Fadli",
//     "warehouse_penerima" : "WH Jakarta",
//     "alamat_penerima" : "Jakarta Raya",
//     "tanggal_pengiriman" : "2024/07/30",
//     "tanggal_sampai" : "2024/08/02",
//     "batch" : "Bacth 10"
// }