// Даны два целых числа A и B ( A < B ). Вывести в порядке убывания все це-
// лые числа, расположенные между A и B (не включая числа A и B ), а также
// количество N этих чисел
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
	x = reserse(x)
	for k := range x {
		fmt.Printf("%v ", x[k])
	}

	fmt.Printf("\n")
	fmt.Printf("чисел %v\n", (m - n - 1))
}

func loop(m, n int) []int {
	var i []int
	m--
	n++
	for n <= m {
		i = append(i, n)
		n++
	}
	return i
}

func reserse(s []int) []int {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}
