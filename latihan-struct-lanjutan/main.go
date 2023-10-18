package main

import "fmt"

type OrangTua struct {
	Nama string
	Umur int
}

type Siswa struct {
	parent OrangTua //embedded struct
	Nama   string
	Umur   int
	Kelas  string
}

type daftarSiswa struct {
	Nama           string
	Umur           int
	Kelas          string
	Nama_Orang_Tua string
	Umur_Orang_Tua int
}
type daftarsiswa []daftarSiswa

func main() {
	//orang tua 1 - siswa1
	parent1 := OrangTua{
		Nama: "Budi",
		Umur: 40,
	}

	siswa1 := Siswa{
		parent: parent1,
		Nama:   "Ali",
		Umur:   15,
		Kelas:  "9A",
	}

	fmt.Println("Informasi Siswa1:")
	fmt.Println("Nama:", siswa1.Nama, ",", "Umur:", siswa1.Umur, ",", "Kelas:", siswa1.Kelas)
	fmt.Println("Orang Tua:", siswa1.parent.Nama, ",", "Umur:", siswa1.parent.Umur)
	fmt.Println()

	//orang tua 2 - siswa2
	parent2 := OrangTua{
		Nama: "Citra",
		Umur: 39,
	}

	siswa2 := Siswa{
		parent: parent2,
		Nama:   "David",
		Umur:   14,
		Kelas:  "8B",
	}

	fmt.Println("Informasi Siswa2:")
	fmt.Println("Nama:", siswa2.Nama, ",", "Umur:", siswa2.Umur, ",", "Kelas:", siswa2.Kelas)
	fmt.Println("Orang Tua:", siswa2.parent.Nama, ",", "Umur:", siswa2.parent.Umur)
	fmt.Println()

	//orang tua 3 - siswa 3
	parent3 := OrangTua{
		Nama: "Eva",
		Umur: 35,
	}

	siswa3 := Siswa{
		parent: parent3,
		Nama:   "Fina",
		Umur:   16,
		Kelas:  "10C",
	}

	fmt.Println("Informasi Siswa3:")
	fmt.Printf("Nama: %s, Umur: %d, Kelas: %s\n", siswa3.Nama, siswa3.Umur, siswa3.Kelas)
	fmt.Printf("Orang Tua: %s, Umur: %d \n", siswa3.parent.Nama, siswa3.parent.Umur)
	fmt.Println()

	var daftarsiswa daftarsiswa

	var dftrsiswa1 = daftarSiswa{Nama: "Eva", Umur: 12, Kelas: "6B", Nama_Orang_Tua: "Rudi", Umur_Orang_Tua: 40}
	daftarsiswa = append(daftarsiswa, dftrsiswa1)
	fmt.Println("Daftar Siswa:")
	fmt.Println("Nama Siswa:", dftrsiswa1.Nama, ",", "Umur:", dftrsiswa1.Umur, ",", "Kelas:", dftrsiswa1.Kelas)
	fmt.Println("Orang Tua:", dftrsiswa1.Nama_Orang_Tua, ",", "Umur:", dftrsiswa1.Umur_Orang_Tua)
	fmt.Println()

	var dftrsiswa2 = daftarSiswa{Nama: "Faisal", Umur: 11, Kelas: "5A", Nama_Orang_Tua: "Dewi", Umur_Orang_Tua: 38}
	daftarsiswa = append(daftarsiswa, dftrsiswa2)
	fmt.Println("Nama Siswa:", dftrsiswa2.Nama, ",", "Umur:", dftrsiswa2.Umur, ",", "Kelas:", dftrsiswa2.Kelas)
	fmt.Println("Orang Tua:", dftrsiswa2.Nama_Orang_Tua, ",", "Umur:", dftrsiswa2.Umur_Orang_Tua)
	fmt.Println()

	var dftrsiswa3 = daftarSiswa{Nama: "Grace", Umur: 10, Kelas: "4C", Nama_Orang_Tua: "Hendro", Umur_Orang_Tua: 42}
	daftarsiswa = append(daftarsiswa, dftrsiswa3)
	fmt.Println("Nama Siswa:", dftrsiswa3.Nama, ",", "Umur:", dftrsiswa3.Umur, ",", "Kelas:", dftrsiswa3.Kelas)
	fmt.Println("Orang Tua:", dftrsiswa3.Nama_Orang_Tua, ",", "Umur:", dftrsiswa3.Umur_Orang_Tua)
}
