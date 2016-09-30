// Даны положительные числа A , B , C . На прямоугольнике размера A × B
// размещено максимально возможное количество квадратов со стороной C
// (без наложений). Найти количество квадратов, размещенных на прямо-
// угольнике. Операции умножения и деления не использовать

package main

import (
	"fmt"

	"github.com/dreddsa5dies/1000GoExamples/ioutil"
)

func main() {
	var (
		a, b, c float64
		k       int
	)

	for a <= 0 {
		a = ioutil.Number("a")
	}
	for b <= 0 {
		b = ioutil.Number("b")
	}
	for c <= 0 {
		c = ioutil.Number("c")
	}

	k = countLoop(a, b, c)
	fmt.Printf("количество квадратов\n")
	fmt.Printf("%v\n", k)
}

func countLoop(a, b, c float64) int {
	var bTemp float64
	var k int
	for a-c >= 0 {
		a = a - c
		bTemp = b
		for bTemp-c >= 0 {
			bTemp = bTemp - c
			k++
		}
	}
	return k
}
