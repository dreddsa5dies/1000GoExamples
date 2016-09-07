// Даны три целых числа. Найти количество положительных чисел в исход-
// ном наборе
package main

import (
	"fmt"

	"github.com/dreddsa5dies/1000GoExamples/util"
)

func main() {
	var n, m int
	var x []int
	n = util.Integer("количество чисел")
	for i := 0; i < n; i++ {
		x = append(x, util.Integer("число"))
	}
	for _, v := range x {
		if v > 0 {
			m++
		}
	}
	fmt.Println(m)
}
