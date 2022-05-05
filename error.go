package main

import (
	"errors"
	"fmt"
)

func pembagian(angka float32, pembagi float32) (float32, error) {
	if pembagi != 0 {
		return angka / pembagi, nil
	} else {
		return 0, errors.New("Angka Pembagi Tidak Boleh 0")
	}
}

func errorInterface() {
	hasil, err := pembagian(7, 0)
	if err == nil {
		fmt.Println(hasil)
	} else {
		fmt.Println("Ups!", err)
	}
}
