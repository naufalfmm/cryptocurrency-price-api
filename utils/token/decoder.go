package token

//go:generate mockgen -package=mockToken -destination=./mockToken/decoder.go -source=decoder.go
type Decoder interface {
	DecodeToken(t string, claims Claims) (Claims, error)
}
