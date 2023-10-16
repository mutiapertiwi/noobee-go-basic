package main

import "fmt"

//membuat struct
type Fruit struct {
	Name   string
	Weigth int
}

func main() {
	//penerapan
	var fruit1 Fruit
	fruit1.Name = "Apple"
	fruit1.Weigth = 1

	fmt.Println(fruit1)
	fmt.Println(fruit1.Name)
	fmt.Println(fruit1.Weigth)

	//peneraan struct 2
	fruit2 := Fruit{
		Name:   "Mango",
		Weigth: 2,
	}

	fmt.Println(fruit2)
	fmt.Println(fruit2.Name)
	fmt.Println(fruit2.Weigth)

	//penerapan struct 3
	var fruit3 = Fruit{
		Name:   "Banana",
		Weigth: 3,
	}

	fmt.Println(fruit3)
	fmt.Println(fruit3.Name)
	fmt.Println(fruit3.Weigth)

	fruit4 := Fruit{"Lemon", 4}

	fmt.Println(fruit4)
	fmt.Println(fruit4.Name)
	fmt.Println(fruit4.Weigth)

}
