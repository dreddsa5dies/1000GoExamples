// Дан набор ненулевых целых чисел; признак его завершения — чис-
// ло 0. Вывести количество чисел в наборе

package main

import (
	"fmt"

	"github.com/dreddsa5dies/1000GoExamples/ioutil"
)

func main() {
	n := ioutil.Integer("число N")
	count := series(n)
	fmt.Printf("Количесвто чисел: %v\n", count)
}

func series(n int) int {
	var (
		r     uint
		count int
	)
	for i := 0; i < n; i++ {
		count++
		r = ioutil.UInteger("число")
		if r == 0 {
			break
		}
	}
	return count
}
