// Дано вещественное число X и целое число N (> 0). Найти значение выражения
// 1 + X + X^2 /(2!) + ... + X^N /( N !)
// ( N ! = 1·2·...· N ). Полученное число является приближенным значением
// функции exp в точке X

package main

import (
	"fmt"
	"math"

	"github.com/dreddsa5dies/1000GoExamples/ioutil"
)

func main() {
	var (
		x float64
		n int
	)

	x = ioutil.Number("X")
	n = -1
	for n < 0 {
		n = ioutil.Integer("N")
	}
	fmt.Println("1 + X + X^2 /(2!) + ... + X^N /( N !)")
	fmt.Printf("%v\n", foo1(n, x))
}

func foo1(n int, x float64) float64 {
	k := 1.0
	for i := 2; i <= n; i++ {
		k = k + (math.Pow(x, float64(i)) / foo2(i))
	}
	return k + x
}

func foo2(n int) float64 {
	k := 1.0
	for i := 1; i <= n; i++ {
		k = k * float64(i)
	}
	return k
}
