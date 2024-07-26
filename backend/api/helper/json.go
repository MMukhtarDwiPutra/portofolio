package helper

import(
	"encoding/json"
	"net/http"
	"fmt"
)

func ReadFromRequestBody(r *http.Request, result interface{}){
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(result)
	fmt.Println(result)
	PanicIfError(err)
}

func WriteToResponseBody(w http.ResponseWriter, response interface{}){
	w.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	err := encoder.Encode(response)
	PanicIfError(err)
}