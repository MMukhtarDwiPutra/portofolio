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
	"strconv"
)

type UserController interface{

}

type userController struct{
	userService service.UserService
}

func NewUserController(userService service.UserService) *userController{
	return &userController{userService}
}

func (c *userController) Login(w http.ResponseWriter, r *http.Request){
	var userInput domain.User
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&userInput)
	helper.PanicIfError(err)
	defer r.Body.Close()

	user := c.userService.GetUserByUsername(userInput.Username)
	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userInput.Password)); err != nil{
		webResponse := web.WebResponse{
			Code : 200,
			Status : "OK",
			Data : map[string]string{
	            "message": "Password Salah",
	        },
		}

		helper.WriteToResponseBody(w, webResponse)
		return
	}
	helper.PanicIfError(err)

	expTime := time.Now().Add(time.Minute * 60)
	claims := &config.JWTClaim{
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer: strconv.Itoa(int(user.ID)),
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

	c.userService.Register(userInput)
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

func (c *userController) GetUser(w http.ResponseWriter, r *http.Request){
	cookie, err := r.Cookie("token")

	claims := &config.JWTClaim{}
	token, err := jwt.ParseWithClaims(cookie.Value, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
        }
		return []byte(config.JWT_KEY), nil
	})

	if err != nil{
		webResponse := web.WebResponse{
			Code : 200,
			Status : "OK",
			Data : map[string]string{
	            "message": "Unauthenticated",
	        },
		}

		helper.WriteToResponseBody(w, webResponse)
	}

	if !token.Valid{
		webResponse := web.WebResponse{
			Code : 200,
			Status : "OK",
			Data : map[string]string{
	            "message": "Token invalid",
	        },
		}

		helper.WriteToResponseBody(w, webResponse)
	}

	id, err := strconv.Atoi(claims.Issuer)
	helper.PanicIfError(err)

	users := c.userService.GetUserById(id)
	
	webResponse := web.WebResponse{
		Code : 200,
		Status : "OK",
		Data : users,
	}

	helper.WriteToResponseBody(w, webResponse)
}