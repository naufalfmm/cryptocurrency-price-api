package validator

import "github.com/go-playground/validator/v10"

//go:generate mockgen -package=mockValidator -destination=./mockValidator/mock.go -source=validator.go
type Validator interface {
	ValidateStruct(i interface{}) error
	Engine() interface{}
}

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) ValidateStruct(i interface{}) error {
	return cv.validator.Struct(i)
}

func (cv *CustomValidator) Engine() interface{} {
	return cv.validator
}
