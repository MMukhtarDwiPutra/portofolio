package helper

import(
	"net/http"
)

func DownloadHandler(w http.ResponseWriter, r *http.Request, namaFile string, fileBytes []byte) {
    // Set the headers to prompt the browser to download the file
    w.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
    w.Header().Set("Content-Disposition", "attachment; filename="+namaFile)
    w.Header().Set("Content-Length", string(len(fileBytes)))

    // Write the file to the response
    w.Write(fileBytes)
}