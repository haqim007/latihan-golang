package main

import (
	"fmt"
	"strconv" // untuk convert tipe data
)

var pelajaran string = "Go"

func main(){

	// nama := "Haqim"
	// umur := 22
	// fmt.Println("Hello, nama saya "+ nama)
	// fmt.Println("umur saya", umur)

	// fmt.Println("Hello, ini adalah bahasa "+ pelajaran)

	var hasil int

	a := 9
	b := 10

	const A = 100

	hasil = a-b

	fmt.Println("hasilnya adalah " + strconv.Itoa(hasil))
	fmt.Println("constant", A)

}