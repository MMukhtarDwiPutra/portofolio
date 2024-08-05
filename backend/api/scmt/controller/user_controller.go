package controller

import(
	"portofolio.com/domain/scmt"
	"portofolio.com/service/scmt"
	"portofolio.com/api/helper"
	"encoding/json"
	"net/http"
	"golang.org/x/crypto/bcrypt"
	"portofolio.com/domain"
	"portofolio.com/api/config"
	"time"
	"github.com/golang-jwt/jwt/v4"
	"fmt"
)

type UserController interface{

}

type userController struct{
	service service.UserService
}

func NewUserController(service service.UserService) *userController{
	return &userController{service}
}

func (c *userController) Login(w http.ResponseWriter, r *http.Request){
	var userInput domain.User
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&userInput)
	helper.PanicIfError(err)
	defer r.Body.Close()

	user := c.service.GetUser(userInput.Username)
	fmt.Println(user)
	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userInput.Password)); err != nil{
		fmt.Println(user.Password)
		fmt.Println(userInput.Password)
	}
	helper.PanicIfError(err)

	expTime := time.Now().Add(time.Minute * 1)
	claims := &config.JWTClaim{
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer: "go-jwt-mux",
			ExpiresAt: jwt.NewNumericDate(expTime),
		},
	}

	tokenAlgo := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := tokenAlgo.SignedString(config.JWT_KEY)
	helper.PanicIfError(err)
	http.SetCookie(w, &http.Cookie{
		Name:"token",
		Path:"/",
		Value:token,
		HttpOnly: true,
	})

	webResponse := web.WebResponse{
		Code : 200,
		Status : "OK",
		Data : map[string]string{
            "message": "Login sukses!",
        },
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (c *userController) Register(w http.ResponseWriter, r *http.Request){
	var userInput domain.User
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&userInput)
	helper.PanicIfError(err)
	defer r.Body.Close()

	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(userInput.Password), bcrypt.DefaultCost)
	userInput.Password = string(hashPassword)

	c.service.Register(userInput)
	webResponse := web.WebResponse{
		Code : 200,
		Status : "OK",
		Data : userInput,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (c *userController) Logout(w http.ResponseWriter, r *http.Request){
	http.SetCookie(w, &http.Cookie{
		Name: "token",
		Path: "/",
		Value: "",
		HttpOnly: true,
		MaxAge: -1,
	})

	webResponse := web.WebResponse{
		Code : 200,
		Status : "OK",
		Data : map[string]string{
            "message": "Logout berhasil!",
        },
	}

	helper.WriteToResponseBody(w, webResponse)
}