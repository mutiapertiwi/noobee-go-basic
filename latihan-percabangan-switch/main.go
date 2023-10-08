package main

import "fmt"

func main() {

	huruf := "a"
	//huruf := "b"

	switch huruf {
	case "a", "i", "u", "e", "o":
		fmt.Println("Huruf", huruf, "adalah huruf vokal")
	case "b", "c", "d", "f", "g":
		fmt.Println("Huruf", huruf, "adalah huruf konsonan")
	default:
		fmt.Println("Input huruf")
	}

}
