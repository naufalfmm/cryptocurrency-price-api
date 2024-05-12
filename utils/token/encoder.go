package token

//go:generate mockgen -package=mockToken -destination=./mockToken/encoder.go -source=encoder.go
type Encoder interface {
	EncodeToken(claims Claims) (string, error)
}
