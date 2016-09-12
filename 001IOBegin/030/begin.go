// Дано значение угла α в радианах (0 < α < 2·π). Определить значение
// этого же угла в градусах, учитывая, что 180° = π радианов. В качестве зна-
// чения π использовать 3.14
package main

import (
	"fmt"
	"math"

	"github.com/dreddsa5dies/1000GoExamples/ioutil"
)

func main() {
	var x float64
	x = ioutil.RadNumber("Введите значение угла α в градусах (0 < α < 2·π)")
	fmt.Println("значение угла в радианах:\t", x*180/math.Pi)
}
