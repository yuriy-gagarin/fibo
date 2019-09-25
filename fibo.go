package fibo

import (
	"math"
)

// FiboRecursive returns nth fibo number
func FiboRecursive(n uint) uint {
	if n < 2 {
		return n
	}

	return FiboRecursive(n-1) + FiboRecursive(n-2)
}

// FiboIterative returns nth fibo number
func FiboIterative(n uint) uint {
	if n < 2 {
		return n
	}

	var a, b, i uint = 0, 1, 2
	for ; i <= n; i++ {
		a, b = b, a+b
	}

	return b
}

var sqrt5 = math.Sqrt(5)

// FiboBinet returns nth fibo number
// https://en.wikipedia.org/wiki/Fibonacci_number#Closed-form_expression
func FiboBinet(n uint) uint {
	if n < 2 {
		return n
	}

	phn := math.Pow(math.Phi, float64(n))
	psn := 1 / phn
	r := (phn - psn) / sqrt5
	return uint(math.Round(r))
}

// FiboBinet2 returns nth fibo number
// https://en.wikipedia.org/wiki/Fibonacci_number#Computation_by_rounding
func FiboBinet2(n uint) uint {
	if n < 2 {
		return n
	}

	phn := math.Pow(math.Phi, float64(n))
	r := phn / sqrt5
	return uint(math.Round(r))
}

// FiboBinet3 returns nth fibo number
// F_n = round(phi*F_n-1)
func FiboBinet3(n uint) uint {
	if n < 2 {
		return n
	}

	r := 1.0
	for i := uint(3); i < n+1; i++ {
		r = math.Round(math.Phi * r)
	}

	return uint(r)
}

// FiboRec2 returns nth fibo number
// this uses the identity F_2n+1 = (F_n+1)^2 + (F_n)^2
func FiboRec2(n uint) uint {
	return fiboRec2(n, make(map[uint]uint))
}

func fiboRec2(n uint, memo map[uint]uint) uint {
	if _, ok := memo[n]; !ok {

		switch {
		case n == 0:
			memo[n] = 0
		case n <= 2:
			memo[n] = 1
		default:
			k := n >> 1
			f1 := fiboRec2(k, memo)
			f2 := fiboRec2(k+1, memo)
			if n%2 == 0 {
				memo[n] = 2*f1*f2 - f1*f1
			} else {
				memo[n] = f2*f2 + f1*f1
			}
		}

	}

	return memo[n]
}
