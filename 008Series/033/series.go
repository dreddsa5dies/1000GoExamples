// Даны целые числа K , N , а также K наборов целых чисел по N элемен-
// тов в каждом наборе. Для каждого набора вывести номер его последнего
// элемента, равного 2, или число 0, если в данном наборе нет двоек

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
				check = j + 1
				fmt.Printf("Номер элемента Массива_%v с числом 2: %v\n", i+1, check)
			}
		}
	}
}
