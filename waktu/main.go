package main

import (
	"fmt"
	"time"
)

func main() {
	var timeDateLocal time.Time = time.Date(2023, 1, 16, 10, 15, 0, 0, time.Local)
	var timeDateUTC time.Time = time.Date(2023, 1, 16, 10, 15, 0, 0, time.UTC)
	var timeNow time.Time = time.Now()

	fmt.Printf("Time Date Local: %v\n", timeDateLocal)
	fmt.Printf("Time Date UTC: %v\n", timeDateUTC)
	fmt.Printf("Time Date Now: %v\n", timeNow)

	layout := "2006-01-02 15:04:05"
	myTimeStr := "2023-01-16 10:00:00"
	myTime, err := time.Parse(layout, myTimeStr)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(myTime.String())
	}

	myTimeStr2 := "2024-01-16T15:04:05Z"
	myTime2, err2 := time.Parse(time.RFC3339, myTimeStr2)
	if err2 != nil {
		fmt.Println("err:", err2)
	} else {
		fmt.Println("Oke:", myTime2)
	}

	now := time.Now()

	fmt.Println("Waktu sekarang:", now.String())
	fmt.Println("Tahun sekarang:", now.Year())
	fmt.Println("Bulan sekarang:", now.Month())
	fmt.Println("Tanggal sekarang:", now.Day())

}
