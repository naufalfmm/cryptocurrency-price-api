package jwt

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/naufalfmm/cryptocurrency-price-api/utils/token"
)

type encoder struct {
	privateKey string
	alg        string
	exp        time.Duration
}

func NewEncoder(privateKey, alg string, exp time.Duration) token.Encoder {
	return &encoder{
		privateKey: privateKey,
		alg:        alg,
		exp:        exp,
	}
}

func (e *encoder) EncodeToken(claims token.Claims) (string, error) {
	if e.exp > 0 {
		claims = claims.SetExp(e.exp)
	}

	privateKey := []byte(e.privateKey)

	newToken := jwt.New(jwt.GetSigningMethod(e.alg))
	newToken.Claims = claims

	signedToken, err := newToken.SignedString(privateKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
