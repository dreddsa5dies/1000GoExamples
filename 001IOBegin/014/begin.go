// Дана длина L окружности. Найти ее радиус R и площадь S круга,
// ограниченного этой окружностью, учитывая, что L = 2· π · R , S = π · R^2 . В ка-
// честве значения π использовать 3.14.
package main

import (
	"fmt"
	"math"

	"github.com/dreddsa5dies/1000GoExamples/util"
)

func main() {
	var x, y float64
	const pi2 = 2 * math.Pi
	x = util.NoNNumber("Введите длину окружности")
	y = x / pi2
	fmt.Println("Радиус круга:\t", y)
	fmt.Println("Площадь круга:\t", y*y*math.Pi)
}
