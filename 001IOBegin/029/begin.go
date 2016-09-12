// Дано значение угла α в градусах (0 < α < 360). Определить значение
// этого же угла в радианах, учитывая, что 180° = π радианов. В качестве зна-
// чения π использовать 3.14
package main

import (
	"fmt"
	"math"

	"github.com/dreddsa5dies/1000GoExamples/ioutil"
)

func main() {
	var x float64
	x = ioutil.GrNumber("Введите значение угла α в градусах (0 < α < 360)")
	fmt.Println("значение угла в радианах:\t", x*math.Pi/180)
}
