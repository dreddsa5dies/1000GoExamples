// С начала суток прошло N секунд ( N — целое). Найти количество се-
// кунд, прошедших с начала последней минуты
package main

import (
	"fmt"

	"github.com/dreddsa5dies/1000GoExamples/ioutil"
)

func main() {
	var x uint
	for x == 0 {
		x = ioutil.UInteger("прошло N секунд ( N — целое)")
	}
	x = x % 60
	fmt.Printf("количество секунд, прошедших с начала последней минуты: %v\n", x)
}
