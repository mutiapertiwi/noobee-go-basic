package main

import "fmt"

func main() {
	num := 4
	//num :=5

	if num%4 == 0 {
		fmt.Println(num, "Merupakan bilangan genap")
	} else if num%4 == 1 {
		fmt.Println(num, "Merupakan bilangan ganjil")
	} else {
		fmt.Println("Silahkan input angka")
	}
}
