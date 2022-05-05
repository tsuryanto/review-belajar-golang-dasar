package main

import (
	"fmt"
	"reflect"
)

func tipedata() {
	fmt.Println("Hello World")

	// variable
	var name string = "Taufiq Suryanto"
	fmt.Println("Nama Saya :", name)
	var usia int16 = 23
	fmt.Printf("\nUsia %d tahun", usia)

	// Convert ke tipe data lainnya
	hurufAngka := 10.12 // float
	fmt.Println(int8(hurufAngka))

	// Array
	mataAngin := [4]string{"Utara", "Selatan", "Timur", "Barat"}
	fmt.Println("\n", reflect.TypeOf(mataAngin))
	for _, arah := range mataAngin {
		fmt.Println(arah)
	}

	// slice
	slice1 := mataAngin[1:3]
	fmt.Println(slice1)

	slice2 := append(slice1, "BaratDaya")
	fmt.Println(slice2)

	slice1[0] = "Timur Laut"
	fmt.Println(slice2)

	numbers := []int{1, 2, 3, 4, 5}
	for _, number := range numbers {
		fmt.Println(number)
	}

	// map
	var biodata = make(map[string]string)
	biodata["name"] = "Taufiq"
	biodata["usia"] = "23"

	fmt.Println(biodata)

}
