// Даны два числа. Вывести вначале большее, а затем меньшее из них
package main

import (
	"fmt"

	"github.com/dreddsa5dies/1000GoExamples/ioutil"
)

func main() {
	var x, y int
	x = ioutil.Integer("число 1")
	y = ioutil.Integer("число 2")
	if x > y {
		fmt.Printf("большее: %v\n", x)
		fmt.Printf("меньшее: %v\n", y)
	} else {
		fmt.Printf("большее: %v\n", y)
		fmt.Printf("меньшее: %v\n", x)
	}
}
