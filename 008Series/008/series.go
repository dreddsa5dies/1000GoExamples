// Дано целое число N и набор из N целых чисел. Вывести в том же по-
// рядке все четные числа из данного набора и количество K таких чисел

package main

import (
	"fmt"
	"math/rand"

	"github.com/dreddsa5dies/1000GoExamples/ioutil"
)

func main() {
	n := ioutil.Integer("число N")
	count := series(n)
	fmt.Printf("Количество чисел: %v\n", count)
}

// возврат суммы sum + int(r)
func series(n int) int {
	var r, count int
	for i := 0; i < n; i++ {
		r = rand.Intn(n)
		if r%2 == 0 {
			count++
			fmt.Printf("четное число: %v\n", r)
		}
	}
	return count
}
