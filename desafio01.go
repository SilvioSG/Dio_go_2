package main

import "fmt"

func main() {
	sum := 0
	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Println(i, "É divisível por 3")
			sum += i
		}

	}
	fmt.Println("Soma dos números divisíveis por 3 de 1 a 100 é:", sum)
}
