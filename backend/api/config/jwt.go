package config

import (
	"github.com/golang-jwt/jwt/v4"
)

var JWT_KEY = []byte("qweh1ui2hiuoohaohe129oe11")

type JWTClaim struct{
	Username string
	jwt.RegisteredClaims
}