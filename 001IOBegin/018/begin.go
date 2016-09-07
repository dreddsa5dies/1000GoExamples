// Даны три точки A , B , C на числовой оси. Точка C расположена между
// точками A и B . Найти произведение длин отрезков AC и BC
package main

import (
	"fmt"

	"github.com/dreddsa5dies/1000GoExamples/util"
)

func main() {
	var x, y, z float64
	x = util.Number("Введите точку А")
	y = util.Number("Введите точку B")
	if x < y {
		z = util.Number("Введите точку C")
		for z < x || z > y {
			fmt.Println("Ввод неверных данных, введите C меньше В и больше А")
			_, err := fmt.Scanf("%f", &z)
			if err != nil {
				panic("Ввод неверных данных")
			}
		}
	} else {
		z = util.Number("Введите точку C")
		for z > x || z < y {
			fmt.Println("Ввод неверных данных, введите C меньше A и больше B")
			_, err := fmt.Scanf("%f", &z)
			if err != nil {
				panic("Ввод неверных данных")
			}
		}
	}
	sum := util.ModNumber(x-z) * util.ModNumber(y-z)
	fmt.Println("произведение длинн отрезков:\t", sum)
}
