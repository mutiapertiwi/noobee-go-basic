package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name  string
	Class string
}

func main() {
	users := []User{
		User{Name: "NooB", Class: "A"},
		User{Name: "Java", Class: "B"},
		User{Name: "Docker", Class: "C"},
	}

	fmt.Println("Bentuk dasar", users)

	//encode object ke JSON
	userJSON, err := json.Marshal(users)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Bentuk JSON", string(userJSON))

	//decode json ke object
	var users2 []User
	fmt.Println("User saat ini", len(users2))

	//data yang di kirim dari frontend
	dataDariFrontend := `{"Name": "NooBee", "Class":"B"}`
	//menguba string menjadi byte
	data := []byte(dataDariFrontend)
	var user User

	//proses passing
	err = json.Unmarshal(data, &user)
	if err != nil {
		fmt.Println("Tamba data gagal! \n", err.Error())
	}
	users2 = append(users2, user)
	fmt.Println("User telah berhasil di tambah dari fronted", len(users2))
	fmt.Println(users2)
}
