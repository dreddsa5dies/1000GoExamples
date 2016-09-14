// Дано вещественное число — цена 1 кг конфет. Вывести стоимость 1,
// 2, ..., 10 кг конфет
package main

import (
	"fmt"

	"github.com/dreddsa5dies/1000GoExamples/ioutil"
)

func main() {
	var m float64

	m = ioutil.Number("цена 1 кг конфет")
	fmt.Println("стоимость")
	for i := 1; i <= 10; i++ {
		fmt.Printf("%v кг конфет\t %v\n", i, m*float64(i))
	}
}
