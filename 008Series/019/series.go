// Дано целое число N (> 1) и набор из N целых чисел. Вывести те эле-
// менты в наборе, которые меньше своего левого соседа, и количество K та-
// ких элементов

package main

import (
	"fmt"

	"math/rand"

	"github.com/dreddsa5dies/1000GoExamples/ioutil"
)

func main() {
	outSeries(newSeries(series(ioutil.UInteger("число N"))))
}

// создания среза int
func series(n uint) (x []int) {
	r := 0
	for i := 0; i < int(n); i++ {
		// генерация рандомного int
		r = rand.Int()
		x = append(x, r)
	}
	return x
}

// выборка чисел в новый срез и подсчет их
func newSeries(count []int) (int, []int) {
	var nI []int
	for i := 0; i < len(count)-1; i++ {
		if count[i] > count[i+1] {
			nI = append(nI, count[i+1])
		}
	}
	return len(nI), nI
}

// вывод результата
func outSeries(i int, nI []int) {
	fmt.Printf("Количестсво чисел: %v\n", i)
	fmt.Printf("Числа:\n")
	for _, x := range nI {
		fmt.Printf("%v\n", x)
	}
}
