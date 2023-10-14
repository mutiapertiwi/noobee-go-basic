package main

import (
	"fmt"
)

func main() {
	//contoh for loop increment
	for counter := 0; counter < 5; counter++ {
		fmt.Println("Counter", counter)
	}
	fmt.Println("Counter berhenti")

	//contoh for loop decrement
	for i := 10; i >= 1; i-- {
		fmt.Println(i)
	}

	//contoh nested loop
	a := [2]string{"buah", "jus"}
	b := [3]string{"cui", "maxie", "mollie"}
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(b); j++ {
			fmt.Println(a[i], b[j])
		}
	}

	//contoh for range
	//slice
	fruits := []string{"Grape", "Banana", "Apple"}
	for index, fruit := range fruits {
		fmt.Println(index, fruit)
	}

	//map
	persons := map[string]string{
		"name": "Mutiara",
		"job":  "Programmer",
	}

	for _, person := range persons {
		fmt.Println(person)
	}

	//infinite loop
	//for {
	//	fmt.Println("ini tidak berhenti")
	//}

	for j := 0; j < 10; j++ {
		if j%2 == 0 {
			continue
		}
		fmt.Println(j)
	}
	fmt.Println("Selesai")

}
