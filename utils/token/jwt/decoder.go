package jwt

import (
	"github.com/golang-jwt/jwt"
	"github.com/naufalfmm/cryptocurrency-price-api/utils/token"
)

type decoder struct {
	publicKey string
}

func NewDecoder(publicKey string) token.Decoder {
	return &decoder{
		publicKey: publicKey,
	}
}

func (d *decoder) DecodeToken(t string, claims token.Claims) (token.Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(t, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(d.publicKey), nil
	})
	if err != nil || tokenClaims == nil || tokenClaims.Claims == nil {
		return nil, ErrUnclaimedToken
	}

	tokClaims := tokenClaims.Claims.(token.Claims)
	return tokClaims, nil
}
