package main

import "fmt"

func main() {
	//membuat map
	var person = map[string]string{
		"name": "Mutiara",
		"job":  "Budak anabul",
	}

	fmt.Println(person)

	computer := map[string]int{
		"stock": 50,
		"harga": 2000000,
	}
	fmt.Println(computer)

	//Membuat map dengan function make
	laptop := make(map[string]string)
	laptop["merk"] = "Apple"
	laptop["model"] = "Macbook Pro"

	fmt.Println(laptop)

	//membuat map kosong
	mapKosong1 := make(map[string]string)
	fmt.Println(mapKosong1)

	//mengakses elemen map
	fmt.Println(person["name"])
	fmt.Println(person["job"])

	//menambah dan mengubah elemen map
	laptop["tahun"] = "2020"
	laptop["model"] = "Macbook Air"

	fmt.Println(laptop)
	fmt.Println(len(laptop))
}
