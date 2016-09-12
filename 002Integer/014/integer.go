// Дано трехзначное число. В нем зачеркнули первую справа цифру и
// приписали ее слева. Вывести полученное число
package main

import (
	"fmt"

	"github.com/dreddsa5dies/1000GoExamples/ioutil"
)

func main() {
	var x uint
	for x < 100 || x > 1000 {
		x = ioutil.UInteger("Введите трехзначное число A")
	}
	fmt.Printf("полученное число: %v\n", (x%10)*100+10*(x/100)+(x-(100*(x/100)))/10)
}
