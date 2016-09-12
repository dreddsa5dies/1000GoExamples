// Даны два числа. Вывести большее из них
package main

import (
	"fmt"

	"github.com/dreddsa5dies/1000GoExamples/ioutil"
)

func main() {
	var n, m int
	var x []int
	//	n = ioutil.Integer("количество чисел")
	n = 2
	for i := 0; i < n; i++ {
		x = append(x, ioutil.Integer("число"))
	}
	for _, v := range x {
		m = v
	}
	fmt.Printf("большее число: %v\n", m)
}
