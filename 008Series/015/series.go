// Дано целое число K и набор ненулевых целых чисел; признак его за-
// вершения — число 0. Вывести номер первого числа в наборе, большего K .
// Если таких чисел нет, то вывести 0

package main

import (
	"fmt"

	"github.com/dreddsa5dies/1000GoExamples/ioutil"
)

func main() {
	k := ioutil.Integer("число K")
	count := series(k)
	fmt.Printf("номер первого числа в наборе, большего K: %v\n", count)
}

func series(k int) int {
	var (
		r          int
		count, num int
	)
	for {
		count++
		r = ioutil.Integer("число")
		if r == 0 {
			break
		}
		if r > k {
			num = count
			break
		}
	}
	return num
}
