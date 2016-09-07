// Даны два числа. Вывести большее из них
package main

import (
	"fmt"

	"github.com/dreddsa5dies/1000GoExamples/util"
)

func main() {
	var n, m int
	var x []int
	//	n = util.Integer("количество чисел")
	n = 2
	for i := 0; i < n; i++ {
		x = append(x, util.Integer("число"))
	}
	for _, v := range x {
		m = v
	}
	fmt.Printf("большее число: %v\n", m)
}
