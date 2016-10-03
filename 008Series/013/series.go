// Дан набор ненулевых целых чисел; признак его завершения — чис-
// ло 0. Вывести сумму всех положительных четных чисел из данного набора.
// Если требуемые числа в наборе отсутствуют, то вывести 0

package main

import (
	"fmt"

	"github.com/dreddsa5dies/1000GoExamples/ioutil"
)

func main() {
	sum := series()
	fmt.Printf("Сумма: %v\n", sum)
}

func series() int {
	var (
		r   int
		sum int
	)
	for {
		r = ioutil.Integer("число")
		if r == 0 {
			break
		}
		if r%2 == 0 {
			sum = sum + r
		}
	}
	return sum
}
