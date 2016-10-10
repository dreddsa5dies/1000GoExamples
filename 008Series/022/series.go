// Дано целое число N (> 1) и набор из N вещественных чисел. Если
// данный набор образует убывающую последовательность, то вывести 0;
// в противном случае вывести номер первого числа, нарушающего
// закономерность.

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

// проверка на убывающую последовательность
// считаем количество последовательно меньших чисел
// должно быть равно длине среза, тогда выводим 0
// иначе выводим счетчик
func checkSeries(count []int) int {
	check := 1
	for i := 0; i < len(count)-1; i++ {
		if count[i] > count[i+1] {
			check++
		} else {
			break
		}
	}
	if check == len(count) {
		return 0
	} else {
		return (check + 1)
	}
}
