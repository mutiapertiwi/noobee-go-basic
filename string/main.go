package main

import (
	"fmt"
	"strings"
)

func main() {
	string1 := "Nama saya Mutiara"
	cari := "Mutiara"

	isExist := strings.Contains(string1, cari)
	fmt.Printf("Apakah string :%s ada di dalam text %s? %v\n", cari, string1, isExist)

	string2 := "Nama saya Pertiwi"

	fmt.Println("Text asli :", string2)
	fmt.Println("Ubah 1 huruf a:", strings.Replace(string2, "a", "o", 1))
	fmt.Println("Ubah 2 huruf a:", strings.Replace(string2, "a", "o", 2))
	fmt.Println("Ubah 3 huruf a:", strings.Replace(string2, "a", "o", 3))
	fmt.Println("Ubah semua huruf a:", strings.Replace(string2, "a", "o", -1))

	stringSplit := strings.Split(string2, " ")
	fmt.Println(string2)
	fmt.Println(stringSplit)
	stringJoin := strings.Join(stringSplit, "-")
	fmt.Println(string2)
	fmt.Println(stringSplit)
	fmt.Println(stringJoin)
	fmt.Println(strings.ToUpper(string2))
	fmt.Println(strings.ToLower(string2))
}
