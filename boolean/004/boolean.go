// Даны два целых числа: A , B . Проверить истинность высказывания:
// «Справедливы неравенства A > 2 и B ≤ 3»
package main

import (
	"fmt"

	"github.com/dreddsa5dies/1000GoExamples/util"
)

func main() {
	var a, b int
	a = util.Integer("целое число A")
	b = util.Integer("целое число B")
	if a > 2 || b <= 3 {
		fmt.Printf("неравенства A > 2 и B ≤ 3 справедливы\n")
	} else {
		fmt.Printf("неравенства A > 2 и B ≤ 3 несправедливы\n")
	}
}
