// Дано целое число N (> 2) и набор из N вещественных чисел. Набор
// называется пилообразным , если каждый его внутренний элемент либо
// больше, либо меньше обоих своих соседей (то есть является «зубцом»).
// Если данный набор является пилообразным, то вывести 0; в противном
// случае вывести номер первого элемента, не являющегося зубцом

package main

import (
	"fmt"
	"math/rand"

	"github.com/dreddsa5dies/1000GoExamples/ioutil"
)

func main() {
	fmt.Println(checkSeries(series(ioutil.UInteger("число N"))))
}

// создания среза int
func series(n uint) (x []int) {
	r := 0
	for i := 0; i < int(n); i++ {
		// генерация рандомного int
		r = rand.Int()
		// ручной ввод
		// r = ioutil.Integer("число")
		x = append(x, r)
	}
	return x
}

// Если данный набор является пилообразным, то вывести 0; в противном
// случае вывести номер первого элемента, не являющегося зубцом
func checkSeries(count []int) int {
	k := 0
	for i := len(count) - 2; i > 0; i-- {
		if !((count[i-1] < count[i]) && (count[i] > count[i+1])) &&
			!((count[i-1] > count[i]) && (count[i] < count[i+1])) {
			k = i + 1
		}
	}
	return k
}
