// Даны целые числа K , N и набор из N вещественных чисел: A 1 , A 2 , ...,
// A N . Вывести K -e степени чисел из данного набора:
// ( A 1 )^K , ( A 2 )^K , ..., ( A N )^K

package main

import (
	"fmt"
	"math"

	"github.com/dreddsa5dies/1000GoExamples/ioutil"
)

func main() {
	arr := ioutil.RandomInt(ioutil.Integer("число N"))
	k := ioutil.Integer("число K")
	series(arr, k)
}

func series(arr []int, k int) {
	for i := 0; i < len(arr)-1; i++ {
		fmt.Printf("( A%v = %v )^%v = %v\n", i, arr[i], k, math.Pow(float64(arr[i]), float64(k)))
	}
}
