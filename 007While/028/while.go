// Дано вещественное число ε (> 0). Последовательность вещественных
// чисел A K определяется следующим образом:
// A 1 = 2,
// A K = 2 + 1/ A K–1 , K = 2, 3, ... .
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
	fmt.Printf("-------------------------\n")
	fmt.Printf("| K\t| Ak–1\t| Ak\t|\n")
	fmt.Printf("| %v\t| %v\t| %v\t|\n", k, ak, ak1)
	fmt.Printf("-------------------------\n")
}

func countLoop(e float64) (k, ak, ak1 float64) {
	ak1 = 0.0
	ak = 2.0
	k = 1.0
	for ioutil.ModNumber(ak-ak1) >= e {
		k++
		ak1 = ak
		ak = 2 + 1/ak1
	}
	return k, ak, ak1
}
