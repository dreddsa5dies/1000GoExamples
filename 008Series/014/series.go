// Дано целое число K и набор ненулевых целых чисел; признак его за-
// вершения — число 0. Вывести количество чисел в наборе, меньших K

package main

import (
	"fmt"

	"github.com/dreddsa5dies/1000GoExamples/ioutil"
)

func main() {
	k := ioutil.Integer("число K")
	count := series(k)
	fmt.Printf("количество чисел в наборе, меньших K: %v\n", count)
}

func series(k int) int {
	var (
		r     int
		count int
	)
	for {
		r = ioutil.Integer("число")
		if r == 0 {
			break
		}
		if r < k {
			count++
		}
	}
	return count
}
