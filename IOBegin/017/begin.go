// Даны три точки A , B , C на числовой оси. Найти длины отрезков AC
// и BC и их сумму
package main

import (
	"fmt"

	"github.com/dreddsa5dies/1000GoExamples/util"
)

func main() {
	var x, y, z float64
	x = util.Number("Введите точку А")
	y = util.Number("Введите точку B")
	z = util.Number("Введите точку C")
	sum := util.ModNumber(x-z) + util.ModNumber(y-z)
	fmt.Println("длина отрезка АС:\t", util.ModNumber(x-z))
	fmt.Println("длина отрезка ВС:\t", util.ModNumber(y-z))
	fmt.Println("сумма длинн отрезков:\t", sum)
}
