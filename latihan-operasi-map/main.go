package main

import "fmt"

func main() {
	fruits := make(map[string]int)
	fruits["Apple"] = 10
	fruits["Banana"] = 15
	fruits["Orange"] = 8
	fruits["Grape"] = 20

	fmt.Println("Sebelum di tambah/dihapus:")
	fmt.Println(fruits)

	//menambahkan mango dan strawberry
	fruits["Mango"] = 12
	fruits["Strawberry"] = 7
	fmt.Println("Setelah menambahkan mango dan strawberry")
	fmt.Println(fruits)

	//menghapus buah orange
	fmt.Println("Setelah menghapus buah Orange")
	delete(fruits, "Orange")
	fmt.Println(fruits)

	//jumlah apel
	key := "Apple"
	val, isExist := fruits[key]

	if isExist {
		fmt.Println(key, "Jumlah apel yang tersedia adalah", val)
	} else {
		fmt.Println(key, "Jumlah apel tidak tersedia")
	}

	//daftar dan jumlah buah
	fmt.Println("Daftar buah-buahan beserta jumlahnya:")
	for key, value := range fruits {
		fmt.Println(key, ":", value)
	}

}
