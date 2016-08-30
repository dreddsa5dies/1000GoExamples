// Найти значение функции y = 4·(x–3)^6 – 7·(x–3)^3 + 2 при данном значе-
// нии x
package main

import (
	"fmt"
	"math"

	"github.com/dreddsa5dies/1000GoExamples/util"
)

func main() {
	var x float64
	x = util.Number("Введите x")
	fmt.Println("значение функции y:\t", 4*math.Pow((x-3), 6)-7*math.Pow((x-3), 2)+2)
}
