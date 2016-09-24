// Дано целое число N (> 0). Если оно является степенью числа 3, то вы-
// вести True, если не является — вывести False

package main

import (
	"fmt"

	"github.com/dreddsa5dies/1000GoExamples/ioutil"
)

func main() {
	var (
		n int
		x bool
	)

	for n <= 0 {
		n = ioutil.Integer("N")
	}
	x = countLoop(n)
	fmt.Printf("является степенью числа 3 : %v\n", x)
}

func countLoop(n int) bool {
	x := 1
	for x < n {
		x = x * 3
	}
	return (x == n)
}
