package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num1 int = 1
	var num2 string = "1"

	num1Str := strconv.Itoa(num1)
	fmt.Println("Hasil perbandingan:", num1Str == num2)

	var num3 int = 1
	var num4 string = "1"
	num4Int, err := strconv.Atoi(num4)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(num3 == num4Int)
	}

	//numStr := "1234"
	//num, err := strconv.ParseInt(numStr, 10, 32)
	//if err != nil {
	//	fmt.Println(err.Error())
	//} else {
	//	fmt.Println(num)
	//}

	numStr := "12.34"
	num, err := strconv.ParseFloat(numStr, 32)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(num)
	}

	//bString := "true"
	//b, err := strconv.ParseBool(bString)
	//if err != nil {
	//	fmt.Println(err.Error())
	//} else {
	//	fmt.Println(b == true)
	//}

	b := true
	bString := strconv.FormatBool(b)
	fmt.Println(bString == "true")

	num5 := 10
	float1 := float32(num5)
	fmt.Printf("num1 :%d\n", num5)
	fmt.Printf("float1: %f\n", float1)
	fmt.Printf("Casting float to int: %d\n", int(float1))

	str1 := "Haii"
	byte1 := []byte(str1)

	fmt.Printf("str1: %v\n", str1)
	fmt.Printf("byte1: %v\n", byte1)
	fmt.Printf("Cast dari byte1 ke string %v\n", string(byte1))
	fmt.Printf("Cast dari byte ke string %v\n", string(65))
}
