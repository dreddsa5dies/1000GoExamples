// Даны целые числа K , N и набор из N целых чисел. Если в наборе
// имеются числа, меньшие K , то вывести True; в противном случае вывести
// False

package main

import (
	"fmt"

	"github.com/dreddsa5dies/1000GoExamples/ioutil"
)

func main() {
	n := ioutil.Integer("число N")
	k := ioutil.Integer("число K")
	series(n, k)
}

func series(n, k int) {
	var (
		r int
		b bool
	)
	b = false
	for i := 0; i < n; i++ {
		r = ioutil.Integer("число")
		if r < k {
			b = true
		}
	}
	fmt.Printf("%v\n", b)
}
