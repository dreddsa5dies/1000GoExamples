// Дано целое число N (> 1) и набор из N вещественных чисел. Прове-
// рить, образует ли данный набор возрастающую последовательность. Если
// образует, то вывести True, если нет — вывести False

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

// проверка на возрастающую последовательность
// считаем количество последовательно меньших чисел
// должно быть равно длине среза
func checkSeries(count []int) bool {
	check := 0
	for i := 0; i < len(count)-1; i++ {
		if count[i] < count[i+1] {
			check++
		}
	}
	nB := check == len(count)-1
	return nB
}
