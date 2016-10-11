// Дано целое число N и набор из N вещественных чисел: A 1 , A 2 , ..., A N .
// Вывести следующие числа:
// A 1 , ( A 2 ) 2 , ..., ( A N–1 ) N–1 , ( A N ) N

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
	for i := 0; i < len(arr)-1; i++ {
		fmt.Printf("( A%v = %v )^%v = %v\n", i+1, arr[i], i+1, math.Pow(float64(arr[i]), float64(i+1)))
	}
}
