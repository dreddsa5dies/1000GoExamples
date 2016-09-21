// Дано вещественное число X (| X | < 1) и целое число N (> 0). Найти значение выражения
// 1 + X /2 – 1· X^2 /(2·4) + 1·3· X^3 /(2·4·6) – ... +
// + (–1) N–1 ·1·3·...·(2· N –3)· X^N /(2·4·...·(2· N )).
// Полученное число является приближенным значением функции math.Sqrt((1 + X), 2)

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

	x = 2
	for x >= 1 {
		x = ioutil.NotNullNumber("X")
	}

	ioutil.ModNumber(x)

	n = -1
	for n < 0 {
		n = ioutil.Integer("N")
	}
	fmt.Println(`1 + X /2 – 1· X^2 /(2·4) + 1·3· X^3 /(2·4·6) – ... +
	+ (–1) N–1 ·1·3·...·(2· N –3)· X^N /(2·4·...·(2· N ))`)
	fmt.Printf("%v\n", foo1(n, x))
}

func foo1(n int, x float64) float64 {
	k := 1.0 + x/2
	a := 1.0
	for i := 2.0; i <= float64(n); i++ {
		a = -a
		k = k + (a*(2.0*i-3.0)*(math.Pow(x, i)))/foo2(i)
	}
	return k
}

func foo2(x float64) float64 {
	k := 2 * x
	a := 2.0
	for i := 1; i < int(x); i++ {
		if i%2 == 0 {
			a = a * 2 * float64(i)
		}
	}
	return a * k
}
