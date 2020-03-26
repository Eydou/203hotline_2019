//
// EPITECH PROJECT, 2020
// 202unsold_2019
// File description:
// main
//

package main

import (
	"fmt"
	"os"
	"strconv"

	functions "./functions"
)

func help() {
	fmt.Printf("USAGE\n   ./203hotline [n k | d]\n")
	fmt.Printf("DESCRIPTION\n")
	fmt.Printf("n\tn value for the computation of C (n, k)\n")
	fmt.Printf("k\tk value for the computation of C (n, k)\n")
	fmt.Printf("d\taverage duration of calls (in seconds)\n")
	os.Exit(0)
}

func main() {
	args := os.Args

	if len(args) == 2 {
		if args[1] == "-h" || args[1] == "--help" {
			help()
		}
	}
	if _, err := functions.ErrorArgs(args); err != nil {
		fmt.Fprintf(os.Stderr, "\033[31mX\033[0m Error: %s\n", err)
		os.Exit(84)
	}
	number1, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println(err, number1)
		os.Exit(84)
	}
	if len(args) == 3 {
		number2, err := strconv.Atoi(args[2])
		if err != nil {
			fmt.Println(err, number2)
			os.Exit(84)
		}
		os.Exit(functions.ConfBinomial(int64(number1), int64(number2)))
	} else {
		os.Exit(functions.Math(float64(number1)))
	}
}
