// Даны две переменные целого типа: A и B . Если их значения не равны, то
// присвоить каждой переменной сумму этих значений, а если равны, то при-
// своить переменным нулевые значения. Вывести новые значения перемен-
// ных A и B
package main

import (
	"fmt"

	"github.com/dreddsa5dies/1000GoExamples/util"
)

func main() {
	var x, y int
	x = util.Integer("число A")
	y = util.Integer("число B")
	if x == y {
		x = 0
		y = x
		fmt.Printf("A: %v\n", x)
		fmt.Printf("B: %v\n", y)
	} else {
		x = x + y
		y = x
		fmt.Printf("A: %v\n", x)
		fmt.Printf("B: %v\n", y)
	}
}
