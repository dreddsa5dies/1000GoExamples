// Даны три целых числа, одно из которых отлично от двух других, равных
// между собой. Определить порядковый номер числа, отличного от осталь-
// ных
package main

import (
	"fmt"

	"github.com/dreddsa5dies/1000GoExamples/ioutil"
)

func main() {
	var x, y, z int
	fmt.Println("три целых числа, одно из которых отлично от двух других, равных между собой")
	x = ioutil.Integer("число A")
	y = ioutil.Integer("число B")
	z = ioutil.Integer("число C")

	if x == y && y == z {
		panic("все равны\n")
	}

	if x == y {
		fmt.Printf("порядковый номер числа: 3\n")
	} else if x == z {
		fmt.Printf("порядковый номер числа: 2\n")
	} else {
		if y == z {
			fmt.Printf("порядковый номер числа: 1\n")
		} else {
			fmt.Printf("нет равных чисел\n")
		}
	}
}
