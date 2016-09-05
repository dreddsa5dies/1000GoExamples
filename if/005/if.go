// Даны три целых числа. Найти количество положительных и количество от-
// рицательных чисел в исходном наборе
package main

import (
	"fmt"

	"github.com/dreddsa5dies/1000GoExamples/util"
)

func main() {
	var n, m, l int
	var x []int
	n = util.Integer("количество чисел")
	for i := 0; i < n; i++ {
		x = append(x, util.Integer("число"))
	}
	for _, v := range x {
		if v > 0 {
			m++
		} else {
			l++
		}
	}
	fmt.Printf("положительных = %v\n", m)
	fmt.Printf("отрицательных = %v\n", l)
}
