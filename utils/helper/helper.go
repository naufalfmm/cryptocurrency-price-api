package helper

import "strconv"

func SafelyReference[T any](val T) *T {
	return &val
}

func SafelyDereference[T any](val *T) T {
	if val == nil {
		var data T
		return data
	}

	return *val
}

func DefaultConvertFloat64(data string, def float64) float64 {
	conv, err := strconv.ParseFloat(data, 64)
	if err != nil {
		return def
	}

	return conv
}

func DefaultIfEmpty[T comparable](data T, def T) T {
	var emp T
	if data == emp {
		return def
	}

	return data
}
