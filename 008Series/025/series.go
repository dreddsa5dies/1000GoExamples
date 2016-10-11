// Дано целое число N и набор из N целых чисел, содержащий по край-
// ней мере два нуля. Вывести сумму чисел из данного набора, расположен-
// ных между первым и последним нулем (если первый и последний нули
// идут подряд, то вывести 0)

package main

import (
	"fmt"

	"github.com/dreddsa5dies/1000GoExamples/ioutil"
)

func main() {
	fmt.Println(series(ioutil.RandomInt(ioutil.Integer("число N"))))
}

func series(arr []int) int {
	arr = append(arr, 0)
	fmt.Println(arr)
	sum := 0
	check := 0
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] == 0 && check < 2 {
			check++
			for j := i; j < len(arr)-1; j++ {
				sum = sum + arr[j]
				check = 0
			}
		}
	}
	return sum
}
