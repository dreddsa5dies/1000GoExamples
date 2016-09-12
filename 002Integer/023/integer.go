// С начала суток прошло N секунд ( N — целое). Найти количество
// полных минут, прошедших с начала последнего часа
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
	x = (x % 3600) / 60
	fmt.Printf("полных минут, прошедших с начала последнего часа: %v\n", x)
}
