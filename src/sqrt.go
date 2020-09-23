package main

import (
	"fmt"
	"math"
)

// ニュートン法を用いて平方根を求める
func Sqrt(x float64) float64 {
	z := float64(1)

	for i := 0; i < 10; i++ {
		z -= ((z*z - x) / (2*z))
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(math.Sqrt(2))
}
