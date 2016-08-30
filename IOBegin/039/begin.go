// Найти корни квадратного уравнения A·x^2 + B·x + C = 0, заданного
// своими коэффициентами A, B, C (коэффициент A не равен 0), если извест-
// но, что дискриминант уравнения положителен. Вывести вначале меньший,
// а затем больший из найденных корней. Корни квадратного уравнения на-
// ходятся по формуле x 1,2 = ( − B ± D ) / (2· A ), где D — дискриминант , рав-
// ный B^2 – 4· A · C
package main

import (
	"fmt"
	"math"

	"github.com/dreddsa5dies/1000GoExamples/util"
)

func main() {
	var x, y, z, x1, x2, d float64
	fmt.Println("A·x^2 + B·x + C = 0")
	x = util.NotNullNumber("Введите A")
	y = util.Number("Введите B")
	z = util.Number("Введите C")
	d = math.Pow(y, 2) - 4*x*z
	x1 = (-y + d) / (2 * x)
	x2 = (-y - d) / (2 * x)
	if x1 > x2 {
		fmt.Println("x2 = ", x2)
		fmt.Println("x1 = ", x1)
	} else {
		fmt.Println("x1 = ", x1)
		fmt.Println("x2 = ", x2)
	}
}
