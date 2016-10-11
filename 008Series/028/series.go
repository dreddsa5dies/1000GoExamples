// Дано целое число N и набор из N вещественных чисел: A 1 , A 2 , ..., A N .
// Вывести следующие числа:
// ( A 1 )^N , ( A 2 )^N–1 , ..., ( A N–1 )^2 , A N

package main

import (
	"fmt"
	"math"

	"github.com/dreddsa5dies/1000GoExamples/ioutil"
)

func main() {
	arr := ioutil.RandomInt(ioutil.Integer("число N"))
	series(arr)
}

func series(arr []int) {
	i := 0
	j := len(arr) - 1
	for i < len(arr) && j >= 0 {
		fmt.Printf("( A%v = %v )^%v = %v\n", i+1, arr[i], j+1, math.Pow(float64(arr[i]), float64(j+1)))
		i++
		j--
	}
}
