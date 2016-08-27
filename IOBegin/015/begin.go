// Дана площадь S круга. Найти его диаметр D и длину L окружности,
// ограничивающей этот круг, учитывая, что L = π · D , S = π · D^2 /4. В качестве
// значения π использовать 3.14.
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
