package FloatHelper

import (
	"math"
	"strconv"
)

func RoundDown(number float64, precision int) float64 {
	pow := math.Pow10(precision)
	return math.Floor(number*pow) / pow
}

func ParseToFloat64(v string) float64 {
	f, err := strconv.ParseFloat(v, 64)
	if err != nil {
		panic(err)
	}
	return f
}
