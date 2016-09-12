// Дана площадь S круга. Найти его диаметр D и длину L окружности,
// ограничивающей этот круг, учитывая, что L = π · D , S = π · D^2 /4. В качестве
// значения π использовать 3.14.
package main

import (
	"fmt"
	"math"

	"github.com/dreddsa5dies/1000GoExamples/ioutil"
)

func main() {
	var x, y float64
	x = ioutil.NoNNumber("введите площадь S круга")
	y = math.Sqrt((4 * x) / math.Pi)
	fmt.Println("диаметр круга:\t\t", y)
	fmt.Println("длина L окружности:\t", y*math.Pi)
}
