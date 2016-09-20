// Дано вещественное число X и целое число N (> 0). Найти значение выражения
// X – X^3 /(3!) + X^5 /(5!) – ... + (–1) N · X^2·N+1 /((2· N +1)!)
// ( N ! = 1·2·...· N ). Полученное число является приближенным значением
// функции sin в точке X

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
	fmt.Println("X – X^3 /(3!) + X^5 /(5!) – ... + (–1) N · X^2·N+1 /((2· N +1)!")
	fmt.Printf("%v\n", foo1(n, x))
}

func foo1(n int, x float64) float64 {
	k := x
	a := 1.0
	for i := 3; i <= n; i++ {
		a = -a
		if i%2 != 0 {
			k = k + a*(math.Pow(x, float64(i))/foo2(i))
		}
	}
	return k
}

func foo2(n int) float64 {
	k := 1.0
	for i := 1; i <= n; i++ {
		k = k * float64(i)
	}
	return k
}
