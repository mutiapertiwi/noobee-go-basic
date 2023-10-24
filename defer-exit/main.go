package main

import (
	"fmt"
	"os"
)

// defer
func cetak(text string) {
	fmt.Println(text)
}

func main() {

	//defer cetak("Text 1")
	//defer cetak("Text 2")
	//cetak("Text 3")
	//defer cetak("Text 4")
	//defer cetak("Text 5")

	num1 := 5

	if num1%2 > 0 {
		cetak("ini adalah bilangan ganjil")
		//defer cetak("akan berakhir setelah proses di atas selesai")
		func() {
			defer cetak("akan berakhir setelah proses di atas selesai")
		}()

		cetak("Oh tentik tidak, defer di eksekusi setelah ini")

		//exit
		names := []string{"Noob", "Bee", "Go", "NodeJS"}
		for _, name := range names {
			if name == "Go" {
				os.Exit(1)
			}
			cetak(name)
		}
	}

}
