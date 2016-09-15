// Дано целое число N (> 0). Найти значение выражения
// 1.1 – 1.2 + 1.3 – ...
// ( N слагаемых, знаки чередуются). Условный оператор не использовать
package main

import (
	"fmt"

	"github.com/dreddsa5dies/1000GoExamples/ioutil"
)

func main() {
	var m int
	m = -1
	for m < 0 {
		m = ioutil.Integer("N")
	}

	fmt.Println("1.1 – 1.2 + 1.3 – ...")
	fmt.Printf("%5.1f\n", foo1(m))
}

func foo1(m int) float64 {
	k := 1.1
	a := 1.0
	for i := 2.0; i <= float64(m); i++ {
		a = -a
		k = k + a*(1+i/10)
	}
	return k
}
