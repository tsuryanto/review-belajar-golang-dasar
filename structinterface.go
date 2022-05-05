package main

import "fmt"

// Interface People
type people interface {
	getName() string
	getAge() int
}

func SayHello(people people) {
	fmt.Println("Hello", people.getName())
	fmt.Println("I am", people.getAge(), "Years Old")
}

// struct
type Customer struct {
	Name string
	Age  int
}

// Methods milik struct Customer
// Golang otomatis menjadikan Struct ini implementasi Poople karena memiliki atribut function yang sama
func (customer Customer) getName() string {
	return customer.Name
}

func (customer Customer) getAge() int {
	return customer.Age
}

// End Methods milik struct Customer

func structinterface() {
	rafi := Customer{
		Name: "Ahmad Arafi",
		Age:  10,
	}
	SayHello(rafi)
}
