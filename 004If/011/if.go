// Даны две переменные целого типа: A и B . Если их значения не равны, то
// присвоить каждой переменной большее из этих значений, а если равны, то
// присвоить переменным нулевые значения. Вывести новые значения пере-
// менных A и B
package main

import (
	"fmt"

	"github.com/dreddsa5dies/1000GoExamples/ioutil"
)

func main() {
	var x, y int
	x = ioutil.Integer("число A")
	y = ioutil.Integer("число B")
	if x == y {
		x = 0
		y = x
		fmt.Printf("A: %v\n", x)
		fmt.Printf("B: %v\n", y)
	} else if x > y {
		y = x
		fmt.Printf("A: %v\n", x)
		fmt.Printf("B: %v\n", y)
	} else {
		x = y
		fmt.Printf("A: %v\n", x)
		fmt.Printf("B: %v\n", y)
	}
}
