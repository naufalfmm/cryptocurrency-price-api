package token

//go:generate mockgen -package=mockToken -destination=./mockToken/data.go -source=data.go
type Data interface {
	CreatedBy() string
}
