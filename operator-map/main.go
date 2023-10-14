package main

import "fmt"

func main() {
	nilai := map[string]int{
		"Agama":     80,
		"Matematik": 90,
		"Olahraga":  70,
		"Design":    90,
	}

	fmt.Println("Sebelum dihapus", nilai)
	fmt.Println("Length", len(nilai))

	//menghapus item map
	delete(nilai, "Olahraga")

	fmt.Println("Setelah dihapus", nilai)
	fmt.Println("Lengt", len(nilai))

	//mencari item map
	key := "Agama"
	val, isExist := nilai[key]

	if isExist {
		fmt.Println(key, "is exist with value", val)
	} else {
		fmt.Println(key, "is not exist")
	}

	//slice of map
	students := []map[string]string{
		map[string]string{"name": "Mutiara", "major": "Teknik komputer"},
		map[string]string{"name": "Pertiwi", "major": "Teknik informasi"},
		map[string]string{"name": "Mollie", "major": "Teknik informatika"},
	}

	for _, student := range students {
		fmt.Println(student["name"], "berada di jurusan", student["major"])
	}
}
