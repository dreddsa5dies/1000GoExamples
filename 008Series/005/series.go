// Дано целое число N и набор из N положительных вещественных чисел.
// Вывести в том же порядке целые части всех чисел из данного набора (как
// вещественные числа с нулевой дробной частью), а также сумму всех целых
// частей

package main

import (
	"fmt"
	"math/rand"

	"github.com/dreddsa5dies/1000GoExamples/ioutil"
)

func main() {
	n := ioutil.Integer("число N")
	sum := series(n)
	fmt.Printf("Сумма: %v\n", sum)
}

// возврат суммы sum + int(r)
func series(n int) (sum int) {
	var r float64
	for i := 0; i < n; i++ {
		// генерация рандомного чисел floats в
		// диапазоне `f < 10.0`.
		r = (rand.Float64() * 5) + 5
		fmt.Printf("число %v\n", r)
		fmt.Printf("целая часть числа %v\n", int(r))
		sum = sum + int(r)
	}
	return sum
}
