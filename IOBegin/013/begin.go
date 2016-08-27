// Даны два круга с общим центром и радиусами R 1 и R 2 ( R 1 > R 2 ). Найти
// площади этих кругов S 1 и S 2 , а также площадь S 3 кольца, внешний радиус
// которого равен R 1 , а внутренний радиус равен R 2 :
// S 1 = π ·( R 1 )^2 ,
// S 2 = π ·( R 2 )^2 ,
// S 3 = S 1 – S 2 .
// В качестве значения π использовать 3.14.
package main

import (
	"fmt"
	"math"

	"github.com/dreddsa5dies/1000GoExamples/util"
)

func main() {
	var x, y float64
	fmt.Println("Необходимо найти площади кругов, кольца")
	x = util.NotNullNumber("Введите радиус 1")
	x = util.ModNumber(x)
	y = util.NotNullNumber("Введите радиус 2")
	y = util.ModNumber(y)
	z1 := math.Pi * x * x
	z2 := math.Pi * y * y
	fmt.Println("Площадь круга 1:\t", z1)
	fmt.Println("Площадь круга 2:\t", z2)
	fmt.Println("Площадь кольца:\t\t", util.ModNumber(z1-z2))
}
