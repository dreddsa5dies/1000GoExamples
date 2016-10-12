// Даны целые числа K , N , а также K наборов целых чисел по N элемен-
// тов в каждом наборе. Найти количество наборов, содержащих число 2. Ес-
// ли таких наборов нет, то вывести 0

package main

import (
	"fmt"

	"github.com/dreddsa5dies/1000GoExamples/ioutil"
)

func main() {
	k := ioutil.Integer("число K")
	n := ioutil.Integer("число N")
	check := 0
	for i := 0; i < k; i++ {
		arr := ioutil.RandomInt(n)
		for j := 0; j < len(arr); j++ {
			if arr[j] == 2 {
				check++
			}
		}
	}
	fmt.Printf("Количество массивов с числом 2 = %v\n", check)
}
