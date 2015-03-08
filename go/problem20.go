package main

import (
	"fmt"
	"math/big"
)

func main() {
	result := big.NewInt(1)
	for i := 2; i <= 100; i++ {
		result.Mul(result, big.NewInt(int64(i)))
	}
	fmt.Println(result)

	var sum int64 = 0
	for result.Cmp(big.NewInt(0)) > 0 {
		mod := new(big.Int)
		result.DivMod(result, big.NewInt(10), mod)
		sum += mod.Int64()
	}
	fmt.Println(sum)
}
