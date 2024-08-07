package controller

import(
	"portofolio.com/service/scmt"
	"net/http"
	"os"
	"io"
	"portofolio.com/api/helper"
	"portofolio.com/domain"
	// "portofolio.com/domain/scmt"
)

type GudangController interface{

}

type gudangController struct{
	gudangService service.GudangService
}

func NewGudangController(gudangService service.GudangService) *gudangController{
	return &gudangController{gudangService}
}

func (c *gudangController) UploadGudangBulk(w http.ResponseWriter, r *http.Request){
	// Retrieve the file from form data
	file, _, err := r.FormFile("file")
	helper.PanicIfError(err)
	defer file.Close()

	// Create a destination file
	destPath := "template/uploaded_gudang.xlsx"
	destFile, err := os.Create(destPath)
	helper.PanicIfError(err)

	// Copy the uploaded file data to the destination file
	_, err = io.Copy(destFile, file)
	helper.PanicIfError(err)

	c.gudangService.UploadNewGudang()

	// Attempt to delete the file
	defer func() {
		destFile.Close()
		
		// Path to the file you want to delete
		filePath := "template/uploaded_gudang.xlsx"

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