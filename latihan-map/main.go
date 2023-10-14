package main

import "fmt"

func main() {
	nomor := make(map[string]string)

	nomor["Alice"] = "1234567890"
	nomor["Bob"] = "9876543210"

	fmt.Println("Semua Kontak", nomor)
	fmt.Println("Nomor Telepon Alice:", nomor["Alice"])

	nomor["Charlie"] = "5555555555"

	fmt.Println("Setelah menambahkan kontak Charlie:", nomor)

	delete(nomor, "Bob")

	fmt.Println("Setelah hapus kontak Bob", nomor)

	fmt.Println("Data Kontak")

	for key, value := range nomor {
		fmt.Printf("%s:%s\n", key, value)
	}
}
