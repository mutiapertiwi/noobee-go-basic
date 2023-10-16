package main

import "fmt"

type Parent struct {
	Nama string
	Umur int
}

type Student struct {
	orangTua Parent // embedded struct
	Nama     string
	Umur     int
	Kelas    string
}

//deklarasi slice of struct
type Employee struct {
	Name        string
	Departement string
}
type employees []Employee

func main() {
	parent1 := Parent{
		Nama: "Andi",
		Umur: 50,
	}

	student1 := Student{
		orangTua: parent1,
		Nama:     "Mutiara",
		Umur:     12,
		Kelas:    "6E",
	}
	fmt.Println(student1)
	fmt.Printf("Orang tua %s Bernama %s\n", student1.Nama, student1.orangTua.Nama)

	student2 := Student{
		orangTua: Parent{
			Nama: "Jojo",
			Umur: 51,
		},
		Nama:  "Pertiwi",
		Umur:  10,
		Kelas: "5B",
	}
	fmt.Println(student2)
	fmt.Printf("Orang tua %s Bernama %s\n", student2.Nama, student2.orangTua.Nama)

	//penerapan slice of struct
	var employees employees
	fmt.Println(employees)

	var emp1 = Employee{Name: "Emp1", Departement: "Tech"}
	employees = append(employees, emp1)
	fmt.Println(employees)

	employees = append(employees, Employee{Name: "Emp2", Departement: "Finance"})
	fmt.Println(employees)
}
