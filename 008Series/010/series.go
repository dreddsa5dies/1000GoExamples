// Дано целое число N и набор из N целых чисел. Если в наборе имеют-
// ся положительные числа, то вывести True; в противном случае вывести
// False

package main

import (
	"fmt"

	"github.com/dreddsa5dies/1000GoExamples/ioutil"
)

func main() {
	n := ioutil.Integer("число N")
	series(n)
}

// возврат суммы sum + int(r)
func series(n int) {
	var (
		r int
		b bool
	)
	b = false
	for i := 0; i < n; i++ {
		r = ioutil.Integer("число")
		if r > 0 {
			b = true
		}
	}
	fmt.Printf("%v\n", b)
}
