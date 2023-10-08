package main

import "fmt"

func main() {

	huruf := "a"
	//huruf := "b"

	switch huruf {
	case "a":
		fmt.Println("Huruf", huruf, "adalah huruf vokal")
	case "b":
		fmt.Println("Huruf", huruf, "adalah huruf konsonan")
	default:
		fmt.Println("Input huruf")
	}

}
