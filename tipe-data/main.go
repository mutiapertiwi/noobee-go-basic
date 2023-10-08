package main

import "fmt"

func main() {
	//Contoh tipe data string
	var txt string = "Mutiara sayang anabul"
	//fmt.Println(txt)
	fmt.Printf("Value: %v, tipe data: %T\n", txt, txt)

	var txt2 string = `
	Cui nakal suka pipis "sembarangan"
	Mollie baik 'tapi nakal'
	Mantap..
	`
	fmt.Println(txt2)

	//Contoh tipe data integer
	var a int16 = 30000
	var age uint8 = 55
	var bigNumber int64 = -4657687980908978675

	fmt.Println(a)
	fmt.Println(age)
	fmt.Println(bigNumber)

	//contoh tipe data float
	var x float64 = 314
	fmt.Printf("%f\n", x)
	fmt.Printf("%.1f\n", x)
	fmt.Printf("%.2f\n", x)
	fmt.Println(x)

	//contoh tipe data bool
	var isComplete bool
	var isMarried bool = true

	fmt.Println(isComplete)
	fmt.Println(isMarried)
}
