package main

import (
	"fmt"
	"strconv"
)


func printNumber() int{

	return 10
}

func multiply(angka1 float32, angka2 float32) float32{
	return angka1 * angka2
}

// function dengan lebih dari 1 return
func getBiography(age int, name string, status string) (string, string){
	ageNow := strconv.Itoa(age)

	return name + " adalah seorang "+ status,
		   "umurnya "+ ageNow


}

// menamai variable yang akan direturn
func luasSegitiga(alas float32, tinggi float32)(luas float32){
	luas = 0.5 * alas * tinggi

	return
}

//switchcase
func lampuLantas(warna string) string{

	var alert string

	switch{
		case warna == "merah":
			alert = "berhenti!"
		case warna == "kuning":
			alert = "pelan-pelan"
		default:
			alert = "jalan"
	}

	return alert
}


//pointer (biasanya untuk mengubah nilai variable antar scope)
func ubahPoint(point *int){
	*point = 200
}

func main(){
	fmt.Println("function printNumber mengembalikan", printNumber())
	fmt.Println("hasil perkalian ", multiply(1.45, 5.33))

	var basicInfo, detailInfo string=  getBiography(22, "haqim", "developer")

	fmt.Println(basicInfo, detailInfo)

	fmt.Println("Luas Segitiga =", luasSegitiga(3.4, 4))

	//loop for
	for i := 0; i < 5; i++ {
		fmt.Println("test loop ke", i)	
	}

	//condition
	var num int = 10
	if num < 100 {
		fmt.Println("nilai num kurang dari 100")
	}else if num == 100{
		fmt.Println("Nilai num sama dengan 100")
	}else{
		fmt.Println("Nilai num lebih dari 100")
	}

	//switch case
	var warnaLantas string = "hijau"
	fmt.Println("warna lampu lantas "+warnaLantas+", silahkan "+ lampuLantas(warnaLantas))

	//pointer
	var point int= 100
	ubahPoint(&point)
	fmt.Println(point)
}