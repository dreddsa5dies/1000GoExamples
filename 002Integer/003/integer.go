// Дан размер файла в байтах. Используя операцию деления нацело,
// найти количество полных килобайтов, которые занимает данный файл
// (1 килобайт = 1024 байта)
package main

import (
	"fmt"

	"github.com/dreddsa5dies/1000GoExamples/ioutil"
)

func main() {
	var x int
	x = ioutil.Integer("Введите размер файла в байтах")
	fmt.Printf("количество полных килобайтов: %v (kb)\n", x/1024)
}
