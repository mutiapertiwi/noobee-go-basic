package main

import "fmt"

func main() {
	iniArray1 := [4]string{"cat", "dog", "chicken", "bird"}        //array
	iniArray2 := [...]string{"BWM", "Renault", "Audi", "Mercedes"} //array
	iniSlice := []string{"Apple", "Banana", "Mango", "Avocado"}    //slice

	fmt.Println(iniArray1)
	fmt.Println(iniArray2)
	fmt.Println(iniSlice)

	//membuat slice dari array
	iniSlice2 := iniArray1[0:2]
	iniSlice3 := iniArray1[1:]
	iniSlice4 := iniArray1[:3]
	iniSlice5 := iniArray1[:]

	fmt.Println(iniSlice2)
	fmt.Println(iniSlice3)
	fmt.Println(iniSlice4)
	fmt.Println(iniSlice5)

	//slice tipe data reference
	iniSlice2[1] = "Cow"
	fmt.Println(iniSlice2)
	fmt.Println(iniSlice3)
	fmt.Println(iniSlice4)
	fmt.Println(iniSlice5)

	//membuat slice dengan fungsi make
	numbers := make([]int, 5, 5)
	numbers[0] = 10
	fmt.Println(numbers)

	iniSlice2[1] = "dog"
	fmt.Println("iniSlice2:", iniSlice2)
	iniSlice2[1] = "fish"
	fmt.Println("iniSlice2:", iniSlice2)

}
