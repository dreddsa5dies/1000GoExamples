// Даны десять вещественных чисел. Найти их среднее арифметическое

package main

import (
	"fmt"

	"github.com/dreddsa5dies/1000GoExamples/ioutil"
)

func main() {
	rezult := series()
	fmt.Printf("Сумма: %v\n", rezult)
}

func series() float64 {
	rezult := 0.0
	r := 0.0
	for i := 0; i < 10; i++ {
		r = ioutil.Number("число")
		rezult = rezult + r
	}
	return rezult / 10
}
