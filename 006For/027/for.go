// Дано вещественное число X (| X | < 1) и целое число N (> 0). Найти значение выражения
// X + 1· X^3 /(2·3) + 1·3· X^5 /(2·4·5) + ... +
// + 1·3·...·(2· N –1)· X^2·N+1 /(2·4·...·(2· N )·(2· N +1)).
// Полученное число является приближенным значением функции arcsin в
// точке X

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
	fmt.Println(`X + 1· X^3 /(2·3) + 1·3· X^5 /(2·4·5) + ... + 
	+ 1·3·...·(2· N –1)· X^2·N+1 /(2·4·...·(2· N )·(2· N +1))`)
	fmt.Printf("%v\n", foo1(n, x))
}

func foo1(n int, x float64) float64 {
	var a float64
	k := x
	for i := 1.0; i <= float64(n); i++ {
		a = 2*i + 1
		k = k + (2.0*i-1.0)*(math.Pow(x, a))/foo2(i)
	}
	return k
}

func foo2(x float64) float64 {
	k := (2 * x) * (2*x + 1)
	for i := 1; i < int(x); i++ {
		if i%2 == 0 {
			x = x * float64(i)
		}
	}
	return x * k
}
