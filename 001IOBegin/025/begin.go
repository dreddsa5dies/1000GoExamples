// Найти значение функции y = 3·x^6 – 6·x^2 – 7 при данном значении x
package main

import (
	"fmt"
	"math"

	"github.com/dreddsa5dies/1000GoExamples/ioutil"
)

func main() {
	var x float64
	x = ioutil.Number("Введите x")
	fmt.Println("значение функции y:\t", 3*math.Pow(x, 6)-6*math.Pow(x, 2)-7)
}
