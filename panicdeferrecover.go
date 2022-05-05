package main

import "fmt"

func runApp(error bool) {
	defer endApp()
	fmt.Println("Aplikasi Mulai")
	if error {
		panic("Terjadi Error ")
	}
	fmt.Println("Aplikasi Berjalan")
}

func endApp() {
	message := recover()
	if message != nil {
		fmt.Println("Ups!", message)
	}

	fmt.Println("Aplikasi Selesai")
}

func panicdeferrecover() {
	runApp(false) // ganti true or false
}
