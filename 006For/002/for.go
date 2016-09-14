// Даны два целых числа A и B ( A < B ). Вывести в порядке возрастания все
// целые числа, расположенные между A и B (включая сами числа A и B ),
// а также количество N этих чисел
package main

import (
	"fmt"

	"github.com/dreddsa5dies/1000GoExamples/ioutil"
)

func main() {
	var m, n int

	m = ioutil.Integer("число B") // B
	n = m + 1
	for n > m {
		n = ioutil.Integer("число A") // A
	}

	x := loop(m, n)
	for k := range x {
		fmt.Printf("%v ", x[k])
	}
	fmt.Printf("\n")
	fmt.Printf("чисел %v\n", (m - n + 1))
}

func loop(m, n int) []int {
	var i []int
	for n <= m {
		i = append(i, n)
		n++
	}
	return i
}
