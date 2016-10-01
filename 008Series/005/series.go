// Дано целое число N и набор из N положительных вещественных чисел.
// Вывести в том же порядке целые части всех чисел из данного набора (как
// вещественные числа с нулевой дробной частью), а также сумму всех целых
// частей

package main

import (
	"fmt"

	"github.com/dreddsa5dies/1000GoExamples/ioutil"
)

func main() {
	var (
		sum int
		r   float64
	)
	n := ioutil.Integer("число N")

	for i := 0; i < n; i++ {
		r = ioutil.NoNNumber("число")
		fmt.Printf("целая часть числа %v\n", int(r))
		sum = sum + series(r)
	}

	fmt.Printf("Сумма: %v\n", sum)
}

// возврат суммы sum + int(r)
func series(r float64) (sum int) {
	sum = sum + int(r)
	return sum
}
