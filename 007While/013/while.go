// Дано число A (> 1). Вывести наименьшее из целых чисел K , для кото-
// рых сумма 1 + 1/2 + ... + 1/ K будет больше A , и саму эту сумму

package main

import (
	"fmt"

	"github.com/dreddsa5dies/1000GoExamples/ioutil"
)

func main() {
	var (
		a    int
		x, y float64
	)

	for a < 1 {
		a = ioutil.Integer("A")
	}
	x, y = countLoop(a)
	fmt.Printf("1 + 1/2 + ... + 1/ K будет больше A\nсумм = %v\tA = %v\n", x, y)

}

func countLoop(a int) (x, y float64) {
	x = 0
	y = 0
	for x <= float64(a) {
		y++
		x = x + 1/y
	}
	return x, y
}
