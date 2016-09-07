// Для данного вещественного x найти значение следующей функции f , при-
// нимающей вещественные значения:
// 					– x , если x ≤ 0,
// f ( x ) =		x^2 , если 0 < x < 2,
// 					4, если x ≥ 2.
package main

import (
	"fmt"

	"math"

	"github.com/dreddsa5dies/1000GoExamples/util"
)

func main() {
	var x, f float64
	x = util.Number("x")

	if x <= 0 {
		f = -x
	} else if x > 0 && x < 2 {
		f = math.Pow(x, 2)
	} else {
		f = 4
	}
	fmt.Printf("f ( x ) = %v\n", f)
}
