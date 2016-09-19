// Дано вещественное число A и целое число N (> 0). Используя один
// цикл, вывести все целые степени числа A от 1 до N
package main

import (
	"fmt"
	"math"

	"github.com/dreddsa5dies/1000GoExamples/ioutil"
)

func main() {
	var (
		a float64
		n int
	)

	a = ioutil.Number("A")

	n = -1
	for n < 1 {
		n = ioutil.Integer("N")
	}
	fmt.Println("все целые степени числа A от 1 до N")
	foo1(a, n)
}

func foo1(a float64, n int) {
	var k float64
	for i := 1; i <= n; i++ {
		k = math.Pow(a, float64(i))
		fmt.Printf("%v^%v=%v\n", a, i, k)
	}
}
