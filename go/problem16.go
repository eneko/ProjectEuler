package main

import (
	"fmt"
	"math/big"
)

func main() {
	result := big.NewInt(2)
	result.Exp(result, big.NewInt(1000), big.NewInt(0))
	fmt.Printf("Result: %v\n", result)

	var total int64 = 0
	for result.Cmp(big.NewInt(0)) == 1 {
		digit := new(big.Int)
		result.DivMod(result, big.NewInt(10), digit)
		total += digit.Int64()
	}
	fmt.Printf("Digit sum: %v\n", total)
}
