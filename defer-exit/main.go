package main

import "fmt"

//defer
func cetak(text string) {
	fmt.Println(text)
}

func main() {

	defer cetak("Text 1")
	defer cetak("Text 2")
	cetak("Text 3")
	defer cetak("Text 4")
	defer cetak("Text 5")

}
