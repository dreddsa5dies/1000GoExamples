// Для данного вещественного x найти значение следующей функции f , при-
// нимающей вещественные значения:
// 					2·sin( x ), если x > 0,
// f ( x ) =
// 					6 – x , если x ≤ 0
package main

import (
	"fmt"
	"math"

	"github.com/dreddsa5dies/1000GoExamples/ioutil"
)

func main() {
	var x, f float64
	x = ioutil.Number("x")

	if x > 0 {
		f = 2 * math.Sin(x)
	} else {
		f = 6 - x
	}
	fmt.Printf("f ( x ) = %v\n", f)
}
