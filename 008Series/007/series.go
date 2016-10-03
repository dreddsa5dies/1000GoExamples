// Дано целое число N и набор из N вещественных чисел. Вывести в том
// же порядке округленные значения всех чисел из данного набора (как целые
// числа), а также сумму всех округленных значений

package main

import (
	"fmt"
	"math"
	"math/rand"

	"github.com/dreddsa5dies/1000GoExamples/ioutil"
)

func main() {
	n := ioutil.Integer("число N")
	sum := series(n)
	fmt.Printf("Сумма: %v\n", sum)
}

func series(n int) (sum int) {
	var r float64
	for i := 0; i < n; i++ {
		// генерация рандомного чисел floats в
		// диапазоне `f < 10.0`.
		r = (rand.Float64() * 5) + 2
		fmt.Printf("число %v\n", r)
		x, y := math.Modf(r)
		if y >= 0.5 {
			r = x + 1
			fmt.Printf("округленное число: %v\n", r)
		} else {
			r = x
			fmt.Printf("округленное число: %v\n", r)
		}
		sum = sum + int(r)
	}
	return sum
}
