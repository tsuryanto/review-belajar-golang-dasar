package main

import "fmt"

// Variadic Function
func sumAll(numbers ...int) int {
	var hasil int
	for _, number := range numbers {
		hasil += number
	}
	return hasil
}

// Return multiple values
func bio(idmhs int) (string, uint8) {
	if idmhs == 123 {
		return "Taufiq", 23
	} else {
		return "None", 0
	}
}

// named return value
func hitung(panjang int, lebar int) (luas int, keliling int) {
	luas = panjang * lebar
	keliling = 2 * (panjang + lebar)
	return
}

// Function as Parameter
func tulisGelar(nama string, filtered func(string) bool) string {
	if !filtered("Hahaha") {
		return nama + " S.Kom"
	} else {
		return nama + "(Isian Anda Gak Bener)"
	}
}

func function() {
	fmt.Println(sumAll(1, 2, 3))

	// slice to variadic function
	angka := []int{1, 2, 3, 4}
	fmt.Println(sumAll(angka...))

	fmt.Println(bio(123))

	var luas, keliling int = hitung(4, 5)
	fmt.Println("Luas =", luas, "Keliling =", keliling)

	// Anonimous Function
	// Function as Value
	filter := func(name string) bool {
		return name != "Hahaha"
	}

	fmt.Println(tulisGelar("Taufiq Suryanto", filter))
}
