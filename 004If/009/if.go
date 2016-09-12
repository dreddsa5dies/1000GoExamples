// Даны две переменные вещественного типа: A , B . Перераспределить значе-
// ния данных переменных так, чтобы в A оказалось меньшее из значений,
// а в B — большее. Вывести новые значения переменных A и B
package main

import (
	"fmt"

	"github.com/dreddsa5dies/1000GoExamples/ioutil"
)

func main() {
	var x, y float64
	x = ioutil.Number("число A")
	y = ioutil.Number("число B")
	if x > y {
		x = x + y
		y = x - y
		x = x - y
		fmt.Printf("A: %v\n", x)
		fmt.Printf("B: %v\n", y)
	} else {
		fmt.Printf("A: %v\n", x)
		fmt.Printf("B: %v\n", y)
	}
}
