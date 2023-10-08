package main

import "fmt"

func main() {
	//contoh single switch case statement
	time := "am"

	switch time {
	case "pm":
		fmt.Println("Malam")
	case "am":
		fmt.Println("Pagi")
	default:
		fmt.Println("Maaf, waktunya salah :)")
	}

	//multi case switch statement
	hari := "minggu"

	switch hari {
	case "senin", "selasa", "rabu", "kamis", "jumat":
		fmt.Println("Weekday")
	case "sabtu", "minggu":
		fmt.Println("Weekned")
	default:
		fmt.Println("Hari tidak valid")
	}
}
