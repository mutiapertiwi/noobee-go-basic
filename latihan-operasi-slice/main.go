package main

import "fmt"

func main() {
	//isi dari slice
	iniSlice := []int{10, 20, 30, 40, 50}

	//mencetak slice
	fmt.Println("====[mySlice]====")
	fmt.Println("Data:", iniSlice)
	fmt.Println("Len:", len(iniSlice))
	fmt.Println("Cap:", cap(iniSlice))
	fmt.Println()

	//len slice 1
	iniSlice1 := iniSlice[0:3]
	//len slice 2
	iniSlice2 := iniSlice[3:]

	//mencetak slice1
	fmt.Println("====[Slice1]====")
	fmt.Println("Data:", iniSlice1)
	fmt.Println("Len:", len(iniSlice1))
	fmt.Println("Cap:", cap(iniSlice1))
	fmt.Println()

	//mencetak slice2
	fmt.Println("====[Slice2]====")
	fmt.Println("Data:", iniSlice2)
	fmt.Println("Len:", len(iniSlice2))
	fmt.Println("Cap:", cap(iniSlice2))
	fmt.Println()

	//menambahkan elemen 60 ke dalam slice1
	iniSlice3 := append(iniSlice1, 60)
	fmt.Println("====[Setelah Append ke Slice1]====")
	fmt.Println("Data:", iniSlice3)
	fmt.Println("Len:", len(iniSlice3))
	fmt.Println("Cap:", cap(iniSlice3))
	fmt.Println("Data:", iniSlice2)
	fmt.Println("Len:", len(iniSlice2))
	fmt.Println("Cap:", cap(iniSlice2))
	fmt.Println()

	//menambahkan elemen 70 kedalam slice2
	iniSlice4 := append(iniSlice2, 70)
	fmt.Println("====[Setelah Append ke Slice2]====")
	fmt.Println("Data:", iniSlice3)
	fmt.Println("Len:", len(iniSlice3))
	fmt.Println("Cap:", cap(iniSlice3))
	fmt.Println("Data:", iniSlice4)
	fmt.Println("Len:", len(iniSlice4))
	fmt.Println("Cap:", cap(iniSlice4))
	fmt.Println()

}
