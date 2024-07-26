package exception

import(
	"github.com/go-playground/validator/v10"
	"portofolio.com/api/helper"
	"portofolio.com/domain"
	"net/http"
)

func ErrorHandler(h http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		defer func(){
			if err := recover(); err != nil {
				if notFoundError(w, r, err){
					return
				}

				if validationErrors(w, r, err){
					return
				}

				internalServerError(w, r, err)
			}
		}()
		h.ServeHTTP(w, r)
	})
}

func notFoundError(w http.ResponseWriter, r *http.Request, err interface{}) bool{
	exception, exist := err.(NotFoundError)
	if exist {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)

		webResponse := web.WebResponse{
			Code : http.StatusNotFound,
			Status : "NOT FOUND",
			Data : exception.Error,
		}

		helper.WriteToResponseBody(w, webResponse)
		return true
	}else {
		return false
	}
}

func validationErrors(w http.ResponseWriter, r *http.Request, err interface{}) bool{
	exception, exist := err.(validator.ValidationErrors)
	if exist {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)

		webResponse := web.WebResponse{
			Code : http.StatusBadRequest,
			Status : "BAD REQUEST",
			Data : exception.Error(),
		}

		helper.WriteToResponseBody(w, webResponse)
		return true
	} else {
		return false
	}
}

func internalServerError(w http.ResponseWriter, r *http.Request, err interface{}){
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)

	webResponse := web.WebResponse{
		Code : http.StatusInternalServerError,
		Status : "INTERNAL SERVER ERROR",
		Data : err,
	}

	helper.WriteToResponseBody(w, webResponse)
}