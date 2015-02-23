package main

import (
	"fmt"
)

// Return both quotient and modulus at once
func DivMod(dividend, divisor int) (quotient, modulus int) {
	modulus = dividend % divisor
	quotient = dividend / divisor
	return
}

// Recursively convert numbers to text (up to 1000)
func word(n int) string {
	if n == 1000 {
		return "onethousand"
	}

	// Hundreds
	quo, mod := DivMod(n, 100)
	if quo > 0 && mod > 0 {
		return word(quo) + "hundredand" + word(mod)
	}
	if quo > 0 {
		return word(quo) + "hundred"
	}

	// Tens
	quo, mod = DivMod(mod, 10)
	if quo > 1 {
		switch quo {
		case 9:
			return "ninety" + word(mod)
		case 8:
			return "eighty" + word(mod)
		case 7:
			return "seventy" + word(mod)
		case 6:
			return "sixty" + word(mod)
		case 5:
			return "fifty" + word(mod)
		case 4:
			return "forty" + word(mod)
		case 3:
			return "thirty" + word(mod)
		case 2:
			return "twenty" + word(mod)
		}
	}
	if quo == 1 {
		switch mod {
		case 9:
			return "nineteen"
		case 8:
			return "eighteen"
		case 7:
			return "seventeen"
		case 6:
			return "sixteen"
		case 5:
			return "fifteen"
		case 4:
			return "fourteen"
		case 3:
			return "thirteen"
		case 2:
			return "twelve"
		case 1:
			return "eleven"
		case 0:
			return "ten"
		}
	}
	switch mod {
	case 9:
		return "nine"
	case 8:
		return "eight"
	case 7:
		return "seven"
	case 6:
		return "six"
	case 5:
		return "five"
	case 4:
		return "four"
	case 3:
		return "three"
	case 2:
		return "two"
	case 1:
		return "one"
	default:
		return ""
	}
}

func main() {
	letters := ""
	for i := 1; i <= 1000; i++ {
		w := word(i)
		fmt.Printf("%v: %v\n", i, w)
		letters += w
	}
	fmt.Printf("Letter count: %v\n", len(letters))
}
