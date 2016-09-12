// Для данного вещественного x найти значение следующей функции f ,
// принимающей значения целого типа:
// 				0, если x < 0,
// f ( x ) =	1, если x принадлежит [0, 1), [2, 3), ...,
// 				–1, если x принадлежит [1, 2), [3, 4), ... .
package main

import (
	"fmt"

	"github.com/dreddsa5dies/1000GoExamples/ioutil"
)

func main() {
	var x, f float64
	x = ioutil.Number("x")

	if x < 0 {
		f = 0
	} else {
		if x >= float64(int(x)) && x < x+1 && int(x)%2 == 0 {
			f = 1
		} else if x >= float64(int(x)) && x < x+1 && int(x)%2 != 0 {
			f = -1
		}
	}

	fmt.Printf("f ( x ) = %v\n", f)
}
