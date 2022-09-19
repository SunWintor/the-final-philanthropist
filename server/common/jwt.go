package common

import (
	"github.com/SunWintor/tfp/configs"
	"github.com/SunWintor/tfp/server/ecode"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type JWT struct {
	signingKey []byte
}

type TFPClaims struct {
	Key string `json:"key"`
	jwt.RegisteredClaims
}

func NewJWT() *JWT {
	return &JWT{
		[]byte(configs.GetConf().JWTSalt),
	}
}

func (j *JWT) CreateClaims(key string) TFPClaims {
	claims := TFPClaims{
		Key: key,
		RegisteredClaims: jwt.RegisteredClaims{
			NotBefore: jwt.NewNumericDate(time.Unix(time.Now().Unix()-600, 0)),
			ExpiresAt: jwt.NewNumericDate(time.Unix(time.Now().Unix()+86400*30, 0)),
			Issuer:    "SunWintor",
		},
	}
	return claims
}

func (j *JWT) CreateToken(claims TFPClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.signingKey)
}

func (j *JWT) ParseToken(tokenString string) (*TFPClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &TFPClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return j.signingKey, nil
	})
	if err != nil {
		return nil, convertJwtError(err)
	}
	if token == nil {
		return nil, ecode.TokenInvalid
	}
	if claims, ok := token.Claims.(*TFPClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, ecode.TokenInvalid
}

func convertJwtError(err error) error {
	if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			return ecode.TokenMalformed
		} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
			return ecode.TokenExpired
		} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
			return ecode.TokenNotValidYet
		} else {
			return ecode.TokenInvalid
		}
	}
	return nil
}
