package validator

import "github.com/naufalfmm/cryptocurrency-price-api/utils/validator"

func NewValidator() (validator.Validator, error) {
	return validator.NewV10()
}
