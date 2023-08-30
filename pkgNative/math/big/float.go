package big

import (
	"math/big"
)

func SetPrec(x float64, prec uint) float64 {
	n, _ := big.NewFloat(x).SetPrec(prec).Float64()
	return n
}
