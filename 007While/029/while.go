// Дано вещественное число ε (> 0). Последовательность вещественных
// чисел A K определяется следующим образом:
// A 1 = 1,
// A 2 = 2,
// A K = ( A K–2 + 2· A K–1 )/3, K = 3, 4, ... .
// Найти первый из номеров K , для которых выполняется условие
// | A K – A K–1 | < ε ,
// и вывести этот номер, а также числа A K–1 и A K

package main

import (
	"fmt"

	"github.com/dreddsa5dies/1000GoExamples/ioutil"
)

func main() {
	var (
		e          float64
		k, ak, ak1 float64
	)

	for e <= 0 {
		e = ioutil.Number("ε")
	}

	k, ak, ak1 = countLoop(e)
	fmt.Printf("| K\t| Ak–1\t| Ak\t|\n")
	fmt.Printf("| %v\t| %v\t| %v\t|\n", k, ak, ak1)
}

func countLoop(e float64) (k, ak1, ak float64) {
	ak1 = 1.0
	ak = 2.0
	k = 2.0
	for ioutil.ModNumber(ak-ak1) >= e {
		ak2 := ak1
		ak1 = ak
		k++
		ak = (ak2 + 2*ak1) / 3
	}
	return k, ak, ak1
}
