package main

import "fmt"

func cetak(text string) {
	fmt.Println(text)
}

func main() {
	num1 := 5
	num2 := 5

	if num1+num2 == 10 {
		defer cetak("Program Selesai")
	}
	fmt.Println("Masukkan angka pertama:", num1)
	fmt.Println("Masukkan angka kedua:", num2)
	fmt.Println("Hasil penjumlahan:", num1+num2)
}
