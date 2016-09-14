// Даны целые числа K и N ( N > 0). Вывести N раз число K
package main

import (
	"fmt"

	"github.com/dreddsa5dies/1000GoExamples/ioutil"
)

func main() {
	var m, n int

	m = ioutil.Integer("число")
	n = -1
	for n < 0 {
		n = ioutil.Integer("число")
	}

	loop(m, n)
}

func loop(m, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%v\n", m)
	}
}
