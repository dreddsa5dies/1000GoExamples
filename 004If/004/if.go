// Даны три целых числа. Найти количество положительных чисел в исход-
// ном наборе
package main

import (
	"fmt"

	"github.com/dreddsa5dies/1000GoExamples/ioutil"
)

func main() {
	var n, m int
	var x []int
	n = ioutil.Integer("количество чисел")
	// ввод чисел в слайс
	for i := 0; i < n; i++ {
		x = append(x, ioutil.Integer("число"))
	}
	// ранжирование слайса
	for _, v := range x {
		if v > 0 {
			m++
		}
	}
	fmt.Println(m)
}
