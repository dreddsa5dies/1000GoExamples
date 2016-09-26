// Дано целое число N (> 1). Найти наибольшее целое число K , при ко-
// тором выполняется неравенство 3^K < N

package main

import (
	"fmt"
	"math"

	"github.com/dreddsa5dies/1000GoExamples/ioutil"
)

func main() {
	var (
		n int
		x int
	)

	for n < 1 {
		n = ioutil.Integer("N")
	}
	x = countLoop(n)
	fmt.Printf("3^K < N, K = %v\n", x)

}

func countLoop(n int) int {
	x := 1
	for math.Pow(3, float64(x)) < float64(n) {
		x++
	}
	return x - 1
}
