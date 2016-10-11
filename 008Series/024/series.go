// Дано целое число N и набор из N целых чисел, содержащий по край-
// ней мере два нуля. Вывести сумму чисел из данного набора, расположен-
// ных между последними двумя нулями (если последние нули идут подряд,
// то вывести 0)

package main

import (
	"fmt"

	"github.com/dreddsa5dies/1000GoExamples/ioutil"
)

func main() {
	fmt.Println(series(ioutil.RandomInt(ioutil.Integer("число N"))))
}

// создания среза int
func series(arr []int) []int {
	arr = append(arr, 0)
	return arr
}
