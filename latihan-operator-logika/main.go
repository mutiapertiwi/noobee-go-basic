package main

import "fmt"

func main() {
	gender := "female"
	age := 21

	//gender:="male"
	//age :=18

	if gender == "female" || age >= 21 {
		fmt.Println("Pelamar dapat dipekerjakan.")
	} else {
		fmt.Println("elamar tidak dapat di pekerjakan")
	}
}
