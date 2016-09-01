// С начала суток прошло N секунд ( N — целое). Найти количество
// полных минут, прошедших с начала суток.
package main

import (
	"fmt"

	"github.com/dreddsa5dies/1000GoExamples/util"
)

func main() {
	var x uint
	for x == 0 {
		x = util.UInteger("прошло N секунд ( N — целое)")
	}
	x = x / 60
	fmt.Printf("полных минут, прошедших с начала суток: %v\n", x)
}
