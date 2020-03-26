//
// EPITECH PROJECT, 2020
// 201yams_2019
// File description:
// error
//

package functions

import (
	"fmt"
	"math"
	"math/big"
	"strconv"
	"time"
)

//ConfBinomial algo binom = factorial(a) / (factorial(b) * factorial(a-b))
func ConfBinomial(a int64, b int64) int {
	if b > a {
		fmt.Printf("\033[31mX\033[0m Error: the second number : %v is greater than %v\n", b, a)
		return 84
	}
	res := new(big.Int)
	fmt.Printf("%v-combination of a set of size %v:\n%v\n", uint64(b), uint64(a), res.Binomial(a, b))
	return 0
}

//PoissonAlgo algo poisson
func PoissonAlgo(k float64, x float64) float64 {
	fact := new(big.Int)
	fact = fact.MulRange(1, int64(k))
	sFact := fmt.Sprintf("%v", fact)
	resFact, _ := strconv.ParseFloat(sFact, 64)

	return (math.Exp(x*(-1)) * math.Pow(x, k)) / resFact
}

//BinomialAlgo algo binomial
func BinomialAlgo(k float64, p float64) *big.Float {
	binom := new(big.Int)
	ope1 := math.Pow(p, k)
	ope2 := math.Pow((1 - p), (3500 - k))
	fOpe1 := fmt.Sprintf("%v", ope1)
	fOpe2 := fmt.Sprintf("%v", ope2)
	var result = new(big.Float)
	var bigOpe2 = new(big.Float)

	fmt.Sscan(fOpe1, result)
	fmt.Sscan(fOpe2, bigOpe2)
	binom.Binomial(int64(3500), int64(k))
	convBinom := new(big.Float).SetInt(binom)
	result.Mul(convBinom, result)
	result.Mul(result, bigOpe2)

	return result // = binomial * pow(p, k) * pow((1 - p), (n - k))
}

//PoissonDist algo poisson distribution
func PoissonDist(nb float64) int {
	fmt.Printf("Poisson distribution:\n")
	start := time.Now()
	//
	overload := float64(1.0)
	p := float64(nb / (8 * 60 * 60) * 3500)

	for k := float64(0); k != 51; k++ {
		fmt.Printf("%.0f -> %.3f", k, PoissonAlgo(k, p))
		if (int(k+1) % 5) == 0 {
			fmt.Printf("\n")
		} else if k != 50 {
			fmt.Printf("\t")
		}
		if k <= 25 {
			overload -= PoissonAlgo(k, p)
		}
	}
	//
	duration := time.Since(start)
	fmt.Printf("\nOverload: %.1f%%", overload*100)
	fmt.Printf("\nComputation time: %.2f ms\n",
		float64(duration.Nanoseconds())/1000000)
	return 0
}

//BinomialDist algo binomial distribution
func BinomialDist(nb float64) int {
	fmt.Printf("Binomial distribution:\n")
	start := time.Now()
	//
	overload := new(big.Float)
	overload.SetFloat64(1)
	hundred := new(big.Float).SetFloat64(100)
	p := float64(nb / (8 * 60 * 60))

	for k := float64(0); k != 51; k++ {
		fmt.Printf("%.0f -> %.3f", k, BinomialAlgo(k, p))

		if (int(k+1) % 5) == 0 {
			fmt.Printf("\n")
		} else if k != 50 {
			fmt.Printf("\t")
		}
		if k <= 25 {
			overload.Sub(overload, BinomialAlgo(k, p))
		}
	}
	//
	duration := time.Since(start)
	fmt.Printf("\nOverload: %.1f%%", overload.Mul(overload, hundred))
	fmt.Printf("\nComputation time: %.2f ms\n\n",
		float64(duration.Nanoseconds())/1000000)
	return 0
}

//Math binom and poisson distibution
func Math(a float64) int {
	BinomialDist(a)
	PoissonDist(a)
	return 0
}
