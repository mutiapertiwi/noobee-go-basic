package main

import "fmt"

//parameter pointer
func changeName(name *string, newName string) {
	fmt.Println("Change name from", *name, "to", newName)
	*name = newName
}

//pointer pada method
type Car struct {
	Name  string
	Color string
}

func (c *Car) setName(newName string) {
	fmt.Println("Try to change from", c.Name, "name to =>", newName)
	c.Name = newName
}

func (c *Car) changeName2(newName string) {
	fmt.Println("Try to change from", c.Name, "name to =>", newName)
	c.Name = newName
	//fmt.Println("didalam fungsi changeName2", c.Name)
}

func main() {
	nama := "NooBee"
	namaPointer := &nama

	fmt.Println(nama)
	fmt.Println(namaPointer)  //mengambil alamat memory
	fmt.Println(*namaPointer) //mengambil value

	var ptr *int
	fmt.Println(ptr)

	//contoh penerapan pointer
	var num1 int = 5
	var num2 *int = &num1

	fmt.Println("===== Nilai Awal ====")
	fmt.Println(num1)
	fmt.Println(num2)
	fmt.Println(*num2)

	fmt.Println("===== Nilai Setelah num1 diubah ====")
	num1 = 6
	fmt.Println(num1)
	fmt.Println(num2)
	fmt.Println(*num2)

	fmt.Println("===== Nilai Setelah num2 diubah ====")
	*num2 = 10
	fmt.Println(num1)
	fmt.Println(num2)
	fmt.Println(*num2)

	//parameter pointer
	myName := "NooBee"
	fmt.Println("Nama awal:", myName)

	changeName(&myName, "NooBeeID")
	fmt.Println("Nama setelah diubah", myName)

	//pointer pada method
	car := Car{Name: "BMW", Color: "White"}
	fmt.Println(car)

	car.setName("Civic")
	fmt.Println(car)

	car.changeName2("Chevrolet")
	fmt.Println(car)
}
