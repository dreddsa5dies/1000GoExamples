// Дано целое число N и набор из N вещественных чисел. Вывести сумму
// и произведение чисел из данного набора

package main

import (
	"fmt"

	"github.com/dreddsa5dies/1000GoExamples/ioutil"
)

func main() {
	n := ioutil.Integer("число N")
	sum, pro := series(n)
	fmt.Printf("Сумма: %v\tПроизведение: %v\n", sum, pro)
}

func series(n int) (sum, pro float64) {
	sum = 0.0
	pro = 1.0
	r := 0.0
	for i := 0; i < n; i++ {
		r = ioutil.Number("число")
		sum = sum + r
		pro = pro * r
	}
	return sum, pro
}
