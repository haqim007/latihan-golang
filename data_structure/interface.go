package main

import (
	"fmt"
	"math"
)

type Bentuk interface{
	Keliling() float64
	Luas() float64
}

type Kotak struct{
	Panjang, Lebar float64
}

type Lingkaran struct{
	Radius float64
}

func main(){
	kotak1 := Kotak{10, 15}
	lingkaran1 := Lingkaran{33}
	hitungBangun(kotak1)
	hitungBangun(lingkaran1)
}

func hitungBangun(b Bentuk){
	fmt.Println("Keliling :", b.Keliling())
	fmt.Println("Luas :", b.Luas())

}

func (k Kotak) Keliling() float64{
	return 2 * k.Panjang + 2 * k.Lebar
}

func (k Kotak) Luas() float64{
	return 2 * k.Panjang * k.Lebar
}

func (l Lingkaran) Keliling() float64{
	return 2 * math.Pi * l.Radius 
}

func (l Lingkaran) Luas() float64{
	return math.Pi * l.Radius * l.Radius
}