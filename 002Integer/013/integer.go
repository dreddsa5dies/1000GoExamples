// Дано трехзначное число. В нем зачеркнули первую слева цифру и
// приписали ее справа. Вывести полученное число
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
	fmt.Printf("полученное число: %v\n", (x%10)*10+100*((x-(100*(x/100)))/10)+x/100)
}
