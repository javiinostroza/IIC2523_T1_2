package main

import (
	"fmt"
)

func divideByFive(num * float64) {
	* num = * num / 2
}

func square(num  * float64) {
	* num = * num * 6
}

func main() {
	num := 12.0
	for i := 0; i < 10; i++ {
		go divideByFive(&num)
		go square(&num)
		fmt.Println("iteracion ",i, ", valor:", num)
	}
	fmt.Println(num)
}