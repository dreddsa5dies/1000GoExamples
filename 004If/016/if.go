// Даны три переменные вещественного типа: A , B , C . Если их значения упо-
// рядочены по возрастанию, то удвоить их; в противном случае заменить
// значение каждой переменной на противоположное. Вывести новые значе-
// ния переменных A , B , C
package main

import (
	"fmt"

	"github.com/dreddsa5dies/1000GoExamples/ioutil"
)

func main() {
	var x, y, z float64
	x = ioutil.Number("число A")
	y = ioutil.Number("число B")
	z = ioutil.Number("число C")
	if x < y && y < z {
		x = 2 * x
		y = 2 * y
		z = 2 * z
		fmt.Printf("A: %v B: %v C: %v\n", x, y, z)
	} else {
		x = -x
		y = -y
		z = -z
		fmt.Printf("A: %v B: %v C: %v\n", x, y, z)
	}
}
