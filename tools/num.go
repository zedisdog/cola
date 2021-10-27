package tools

import (
	"errors"
	"github.com/shopspring/decimal"
	"strconv"
)

func Mul(num interface{}, mul int) (int, error) {
	var tmp decimal.Decimal
	switch num.(type) {
	case string:
		t, _ := strconv.ParseFloat(num.(string), 64)
		tmp = decimal.NewFromFloat(t)
	case float32:
		tmp = decimal.NewFromFloat32(num.(float32))
	case float64:
		tmp = decimal.NewFromFloat(num.(float64))
	case int:
		tmp = decimal.NewFromInt(int64(num.(int)))
	default:
		return 0, errors.New("type is invalid")
	}
	tmp = tmp.Mul(decimal.NewFromInt(int64(mul)))
	tt, _ := tmp.Float64()
	return int(tt), nil
}
