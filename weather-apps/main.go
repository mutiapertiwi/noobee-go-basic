package main

import (
	"log"
	"net/http"
	weather "weather-apps/app"
)

func main() {
	port := ":8000"

	http.HandleFunc("/weather", weather.GetCurrentWeather)
	log.Println("server running at port", port)
	http.ListenAndServe(port, nil)
}
