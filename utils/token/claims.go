package token

import "time"

//go:generate mockgen -package=mockToken -destination=./mockToken/claims.go -source=claims.go
type Claims interface {
	Valid() error
	SetExp(exp time.Duration) Claims
}
