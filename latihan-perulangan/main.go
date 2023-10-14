package main

import "fmt"

func main() {
	// Latihan 1
	/**
	  Silahkan buat bentuk seperti ini :
	  *****
	  ****
	  ***
	  **
	  *
	*/
	var a, b, c int
	a = 1
	for b = 5; b >= a; b-- {

		for c = 1; c <= b; c++ {

			fmt.Print("*")
		}
		fmt.Println()
	}

	fmt.Println()

	// Latihan 2
	/**
	  Silahkan buat bentuk seperti ini :
	  *
	  **
	  ***
	  ****
	  *****
	*/
	var d, e, f int
	d = 5
	for e = 1; e <= d; e++ {

		for f = 1; f <= e; f++ {

			fmt.Print("*")
		}
		fmt.Println()
	}

	fmt.Println()

	// Latihan 3
	/**
	  Silahkan buat bentuk Noo Bee (panjang : 50) dengan kondisi berikut :
	  Jika habis dibagi 3 : Noo
	  Jika habis dibagi 5 : Bee
	  Jika habis dibagi 3 dan 5 : Noo Bee

	  Contoh :
	  1, 2, Noo, 4, Bee, Noo, 7, 8, Noo, Bee, 11, Noo, 13, 14, NooBee, 16, 17, Noo, 19, Bee, ... n
	*/
	for i := 1; i <= 50; i++ {
		bagiTiga := i % 3
		bagiLima := i % 5

		if bagiTiga == 0 && bagiLima == 0 {
			fmt.Printf("NooBee, ")

		} else if bagiTiga == 0 {
			fmt.Print("Noo, ")

		} else if bagiLima == 0 {
			fmt.Print("Bee, ")

		} else {
			fmt.Print(i, ", ")
		}
	}
}
