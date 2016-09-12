// Поменять местами содержимое переменных A и B и вывести новые
// значения A и B
package main

import (
	"fmt"

	"github.com/dreddsa5dies/1000GoExamples/ioutil"
)

func main() {
	var x, y float64
	x = ioutil.Number("Введите A")
	y = ioutil.Number("Введите B")
	x = x + y
	y = x - y
	x = x - y
	fmt.Println("замена А\t: ", x)
	fmt.Println("замена B\t: ", y)
}
