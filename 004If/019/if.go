// Даны четыре целых числа, одно из которых отлично от трех других, рав-
// ных между собой. Определить порядковый номер числа, отличного от ос-
// тальных
package main

import (
	"fmt"

	"github.com/dreddsa5dies/1000GoExamples/ioutil"
)

func main() {
	var x, y, z, d int
	fmt.Println("четыре целых числа, одно из которых отлично от трех других, равных между собой")
	x = ioutil.Integer("число A")
	y = ioutil.Integer("число B")
	z = ioutil.Integer("число C")
	d = ioutil.Integer("число D")

	if x == y && y == z && z == d {
		panic("все равны\n")
	}

	if x == y && x == z {
		fmt.Printf("порядковый номер числа: 4\n")
	} else if x == y && x == d {
		fmt.Printf("порядковый номер числа: 3\n")
	} else {
		if x == z && x == d {
			fmt.Printf("порядковый номер числа: 2\n")
		} else if y == z && y == d {
			fmt.Printf("порядковый номер числа: 1\n")
		} else {
			fmt.Printf("нет равных чисел\n")
		}
	}
}
