package main

import "fmt"

var x int = 1

func main() {
	//Membuat variable
	var myName string = "Mutiara"

	var urName string
	urName = "Cui"

	var nickName = "Mollie"

	myAge := 25

	fmt.Println(myName)
	fmt.Println(urName)
	fmt.Println(nickName)
	fmt.Println(myAge)

	//Variable di luar fungsi
	fmt.Println(x)

	// Deklarasi variable tanpa nilai awal
	var a string
	var b int
	var c bool

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	// Deklarasi multiple variable
	var d, e, f, g = 1, 2, 3, "Hello"

	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)

	// Deklarasi multiple variable dalam blok
	var (
		firstName string = "Mutiara"
		lastName         = "Pertiwi"
		height    int    = 162
	)

	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(height)

}
