// Описать процедуру Mean( X , Y , AMean , GMean ), вычисляющую среднее
// арифметическое AMean = ( X + Y )/2 и среднее геометрическое GMean =
// = Sqrt(X ⋅ Y) двух положительных чисел X и Y ( X и Y — входные, AMean
// и GMean — выходные параметры вещественного типа). С помощью этой
// процедуры найти среднее арифметическое и среднее геометрическое для
// пар ( A , B ), ( A , C ), ( A , D ), если даны A , B , C , D

package main

import (
	"fmt"
	"math"

	"github.com/dreddsa5dies/1000GoExamples/ioutil"
)

func main() {
	arr := ioutil.RandomFlt(4)
	fmt.Printf("%v\n", arr)
	for i := 1; i < len(arr); i++ {
		a, b := mean(arr[0], arr[i])
		fmt.Printf("Число1 = %v, Число2 = %v\t", arr[0], arr[i])
		fmt.Printf("AMean = %v, GMean = %v\n", a, b)
	}
}

func mean(x, y float64) (AMean, GMean float64) {
	AMean = (x + y) / 2
	GMean = math.Sqrt(x * y)
	return AMean, GMean
}
