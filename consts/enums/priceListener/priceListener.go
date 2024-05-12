package priceListener

import (
	"database/sql/driver"
	"encoding/json"

	"github.com/naufalfmm/cryptocurrency-price-api/consts"
)

type (
	PriceListener int

	PriceListenerClass struct {
		Code string `json:"code"`
		Name string `json:"name"`
	}
)

const (
	ReasonListen PriceListener = iota + 1
	ReasonShutdown
)

var ReasonConstants = []PriceListenerClass{
	{"listening", "Listening"},
	{"shutdown", "Shutdown"},
}

func (cs PriceListener) Value() (driver.Value, error) {
	if cs < 1 || int(cs) > len(ReasonConstants) {
		return "", nil
	}
	return ReasonConstants[cs-1].Code, nil
}

func (cs PriceListener) Name() string {
	if cs < 1 || int(cs) > len(ReasonConstants) {
		return ""
	}
	return ReasonConstants[cs-1].Name
}

func (cs PriceListener) Code() string {
	if cs < 1 || int(cs) > len(ReasonConstants) {
		return ""
	}
	return ReasonConstants[cs-1].Code
}

func (cs PriceListener) MarshalJSON() ([]byte, error) {
	if cs < 1 || int(cs) > len(ReasonConstants) {
		return nil, consts.ErrUnknownConstant
	}

	return json.Marshal(ReasonConstants[cs-1])
}
