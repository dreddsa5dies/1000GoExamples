// Дано целое число N (> 0). Найти произведение
// N! = 1·2·...· N
// ( N–факториал ). Чтобы избежать целочисленного переполнения, вычис-
// лять это произведение с помощью вещественной переменной и вывести
// его как вещественное число
package main

import (
	"fmt"

	"github.com/dreddsa5dies/1000GoExamples/ioutil"
)

func main() {
	var (
		n int
	)

	n = -1
	for n < 0 {
		n = ioutil.Integer("N")
	}
	fmt.Println("N–факториал")
	fmt.Printf("%v\n", foo1(n))
}

func foo1(n int) float64 {
	k := 1.0
	for i := 1; i <= n; i++ {
		k = k * float64(i)
	}
	return k
}
