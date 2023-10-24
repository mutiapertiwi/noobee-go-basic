package main

import (
	"errors"
	"fmt"
	"strconv"
)

type Student struct {
	Name   string
	Heigth int
}

func (s *Student) Validate() error {
	if s.Name == "" {
		return errors.New("Nama tidak boleh kosong")
	}

	if len(s.Name) <= 3 {
		return errors.New("Nama harus lebih dari 3")
	}

	if s.Heigth == 0 {
		return errors.New("Tinggi tidak boleh kosong")
	}

	return nil
}

// panic
func division(number1, number2 int) {
	defer catchError()
	if number2 == 0 {
		panic("Tidak dapat membagi angka dengan 0")
	} else {
		result := number1 / number2
		fmt.Println(result)
	}
}

// recover
func catchError() {
	//panic akan masuk ke recover
	err := recover()
	if err != nil {
		fmt.Println("RECOVER", err)
	}
}

func main() {
	//Error
	numStr := "2"

	num, err := strconv.Atoi(numStr)
	if err != nil {
		fmt.Println(numStr, "bukan sebuah angka")
		fmt.Println(err.Error())
	} else {
		fmt.Println(num, "adalah sebuah angka")
	}

	//membuat custom error
	student := Student{Name: "NooBee"}

	err = student.Validate()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Helo", student.Name)
	}

	division(3, 0)
}
