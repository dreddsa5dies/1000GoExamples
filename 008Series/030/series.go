// Даны целые числа K , N , а также K наборов целых чисел по N элемен-
// тов в каждом наборе. Для каждого набора вывести сумму его элементов

package main

import (
	"fmt"

	"github.com/dreddsa5dies/1000GoExamples/ioutil"
)

func main() {
	k := ioutil.Integer("число K")
	n := ioutil.Integer("число N")
	for i := 0; i < k; i++ {
		arr := ioutil.RandomInt(n)
		fmt.Printf("Массив_%v: %v\n", i+1, arr)
		fmt.Printf("Сумма элементов в Массив_%v = %v\n", i+1, series(arr))
	}
}

func series(arr []int) int {
	sumArr := 0
	for i := 0; i < len(arr); i++ {
		sumArr = sumArr + arr[i]
	}
	return sumArr
}
