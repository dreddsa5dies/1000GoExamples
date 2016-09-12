// Известно, что X кг шоколадных конфет стоит A рублей, а Y кг ирисок
// стоит B рублей. Определить, сколько стоит 1 кг шоколадных конфет, 1 кг
// ирисок, а также во сколько раз шоколадные конфеты дороже ирисок
package main

import (
	"fmt"

	"github.com/dreddsa5dies/1000GoExamples/ioutil"
)

func main() {
	var x, y, z, q float64
	x = ioutil.NoNNumber("Введите X кол-во кг конфет")
	y = ioutil.NoNNumber("Введите их стоимость")
	fmt.Println("1 кг стоит:\t", y/x)
	z = ioutil.NoNNumber("Введите Y кол-во кг ирисок")
	q = ioutil.NoNNumber("Введите их стоимость")
	fmt.Println("1 кг стоит:\t", z/q)
	fmt.Printf("шоколадные конфеты дороже ирисок в %v раз\n", y/x/z/q)
}
