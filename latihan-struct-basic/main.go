package main

import "fmt"

type Book struct {
	Title  string
	Author string
}

//untuk book1
func main() {
	book1 := Book{
		Title:  "Pemograman Go",
		Author: "John Doe",
	}
	fmt.Println("Informasi tentang Book 1:")
	fmt.Println("Judul:", book1.Title)
	fmt.Println("Penulis:", book1.Author)
	fmt.Println()

	//untuk book2
	var book2 Book
	book2.Title = "Algoritma dan Struktur Data"
	book2.Title = "Jane Smith"
	fmt.Println("Informasi tentang Book 2:")
	fmt.Println("Judul:", book2.Title)
	fmt.Println("Penulis:", book2.Author)
	fmt.Println()

	//untuk book3
	book3 := Book{"Machine Learning for Beginners", "David Johnshon"}
	fmt.Println("Informasi tentang Book 3:")
	fmt.Println("Judul:", book3.Title)
	fmt.Println("Penulis:", book3.Author)
}
