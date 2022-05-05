package main

import "fmt"

type Address struct {
	District, City, Country string
}

// Pointer di Method
func (address *Address) changeCountryToIndonesia() {
	address.Country = "Indonesia"
}

// Pointer di Function
func changeDistrict(name string, address *Address) {
	address.District = name
}

func pointer() {
	address1 := Address{"Wonogiri", "Jawa Tengah", "Indonesia"}

	// Di Golang , default nya menggunakan copy by Value
	copyAddress1 := address1

	/* === Operator & ==== */

	// Jika ingin copy by reference harus membuat pointer
	// Jika data di var pointer diubah, var akan menunjuk ke data address1
	// sehingga data di address1 akan ikut berubah
	var address2 *Address = &address1 // *Address adalah tipe data untuk pointer ke Address
	address2.District = "Semarang"

	fmt.Println("\n=== Percobaan 1 ===")
	fmt.Println(address1)
	fmt.Println(address2)

	/* === Operator * ==== */
	// Penggunaan tanda * sebagai OPERATOR

	// Mengubah var pointer address2 agar menunjuk ke data baru di memory
	// dan data di address1 akan dipaksa menunjuk ke data yang baru yang di buat address2
	*address2 = Address{"Jakarta Utara", "DKI Jakarta", "Indonesia"}

	// Cara lain Membuar pointer , namun data tersebut di set kosong
	addressKosong := new(Address)
	addressKosong.City = "Jawa Timur"

	fmt.Println("\n=== Percobaan 2 ===")
	fmt.Println(address1)
	fmt.Println(copyAddress1)
	fmt.Println(address2)
	fmt.Println(addressKosong)

	/* === Pointer di Function ==== */
	// agar ketika address diubah di function, data di sini juga ikut berubah
	// Jika tidak menggunakan pointer, golang akan mengcopy data parameter, sehingga data disini tidak berubah

	fmt.Println("\n=== Percobaan 3 ===")
	// ke function biasa
	alamat := Address{"", "Jawa Tengah", "Japan"}
	changeDistrict("Solo", &alamat)
	fmt.Println(alamat)

	// ke struct function (method)
	alamat.changeCountryToIndonesia()
	fmt.Println(alamat)
}
