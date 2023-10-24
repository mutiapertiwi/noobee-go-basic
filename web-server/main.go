package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Index called...")
	hello := "Helo World!"
	w.Write([]byte(hello)) //konversi dari string ke slice of byte
}

// untuk melakukan routing
func handleRoute() {
	http.HandleFunc("/", index)
}

// membuat server
func startServer(port string) {
	handleRoute()
	fmt.Println("server running at port", port)
	http.ListenAndServe(port, nil)
}

func main() {
	port := ":2008"
	startServer(port)

}
