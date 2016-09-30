// Дано целое число N (> 1). Если оно является простым , то есть не
// имеет положительных делителей, кроме 1 и самого себя, то вывести True,
// иначе вывести False

package main

import (
	"fmt"

	"github.com/dreddsa5dies/1000GoExamples/ioutil"
)

func main() {
	var (
		n int
		k bool
	)

	for n <= 0 {
		n = ioutil.Integer("N")
	}
	k = countLoop(n)
	fmt.Printf("число простое?\n")
	fmt.Printf("%v\n", k)
}

func countLoop(n int) bool {
	y := true
	x := 1
	for x < n-1 {
		x++
		if n%x == 0 {
			y = false
		}
	}
	return y
}
