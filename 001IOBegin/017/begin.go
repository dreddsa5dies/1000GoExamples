// Даны три точки A , B , C на числовой оси. Найти длины отрезков AC
// и BC и их сумму
package main

import (
	"fmt"

	"github.com/dreddsa5dies/1000GoExamples/ioutil"
)

func main() {
	var x, y, z float64
	x = ioutil.Number("Введите точку А")
	y = ioutil.Number("Введите точку B")
	z = ioutil.Number("Введите точку C")
	sum := ioutil.ModNumber(x-z) + ioutil.ModNumber(y-z)
	fmt.Println("длина отрезка АС:\t", ioutil.ModNumber(x-z))
	fmt.Println("длина отрезка ВС:\t", ioutil.ModNumber(y-z))
	fmt.Println("сумма длинн отрезков:\t", sum)
}
