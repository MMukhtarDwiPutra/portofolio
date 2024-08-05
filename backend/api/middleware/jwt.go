package middleware

import (
	"net/http"
	"portofolio.com/api/helper"
	"github.com/golang-jwt/jwt/v4"
	"portofolio.com/domain"
	"portofolio.com/api/config"
)

func JWTMiddleware(next http.Handler) http.Handler{
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request){
		c, err := r.Cookie("token")
		if err != nil{
			if err == http.ErrNoCookie{
				webResponse := web.WebResponse{
					Code : 200,
					Status : "OK",
					Data : map[string]string{
			            "message": "Unauthorized",
			        },
				}

				helper.WriteToResponseBody(w, webResponse)
				return
			}
		}

		tokenString := c.Value
		claim := &config.JWTClaim{}
		token, err := jwt.ParseWithClaims(tokenString, claim, func(t *jwt.Token) (interface{}, error) {
			return config.JWT_KEY, nil
		})

		if err != nil{
			v, _ := err.(*jwt.ValidationError)
			switch v.Errors{
			case jwt.ValidationErrorSignatureInvalid:
				webResponse := web.WebResponse{
					Code : 200,
					Status : "OK",
					Data : map[string]string{
			            "message": "Unauthorized",
			        },
				}

				helper.WriteToResponseBody(w, webResponse)
				return
			case jwt.ValidationErrorExpired:
				webResponse := web.WebResponse{
					Code : 200,
					Status : "OK",
					Data : map[string]string{
			            "message": "Unauthorized, Token Expired",
			        },
				}

				helper.WriteToResponseBody(w, webResponse)
				return
			default:
				webResponse := web.WebResponse{
					Code : 200,
					Status : "OK",
					Data : map[string]string{
			            "message": "Unauthorized",
			        },
				}

				helper.WriteToResponseBody(w, webResponse)
				return
			}
		}

		if !token.Valid{
			webResponse := web.WebResponse{
					Code : 200,
					Status : "OK",
					Data : map[string]string{
			            "message": "Unauthorized, Token Expired",
			        },
				}

				helper.WriteToResponseBody(w, webResponse)
			return
		}

		next.ServeHTTP(w, r)
	})
}