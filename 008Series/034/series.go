// Даны целые числа K , N , а также K наборов целых чисел по N элемен-
// тов в каждом наборе. Для каждого набора выполнить следующее действие:
// если в наборе содержится число 2, то вывести сумму его элементов; если в
// наборе нет двоек, то вывести 0

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
				check = series(arr)
				fmt.Printf("Сумма элементов массива = %v\n", check)
			}
		}
	}
}

func series(arr []int) int {
	sum := 0
	for i := 0; i < len(arr); i++ {
		sum = sum + arr[i]
	}
	return sum
}
