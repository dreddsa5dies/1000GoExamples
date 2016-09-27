// Дано целое число N (> 1). Вывести наименьшее из целых чисел K ,
// для которых сумма 1 + 2 + ... + K будет больше или равна N , и саму эту
// сумму

package main

import (
	"fmt"

	"github.com/dreddsa5dies/1000GoExamples/ioutil"
)

func main() {
	var (
		n    int
		x, y int
	)

	for n < 1 {
		n = ioutil.Integer("N")
	}
	x, y = countLoop(n)
	fmt.Printf("1 + 2 + ... + K будет больше или равна N\nсумм = %v\tK = %v\n", x, y)

}

func countLoop(n int) (x, y int) {
	x = 1
	y = 1
	for x < n {
		y++
		x = x + y
	}
	return x, y
}
