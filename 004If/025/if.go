// Для данного целого x найти значение следующей функции f , принимаю-
// щей значения целого типа:
// 				2· x , если x < –2 или x > 2,
// f ( x ) =
// 				–3· x , в противном случае.
package main

import (
	"fmt"

	"github.com/dreddsa5dies/1000GoExamples/ioutil"
)

func main() {
	var x, f float64
	x = ioutil.Number("x")

	if x < -2 || x > 2 {
		f = 2 * x
	} else {
		f = -3 * x
	}
	fmt.Printf("f ( x ) = %v\n", f)
}
