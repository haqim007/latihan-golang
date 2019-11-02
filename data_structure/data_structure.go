package main

import (
	"fmt"
	"strconv"
	"data_structure/model"
)


type Nilai struct{
	Matkul string
	Nilai int
}

// method 
func (m Nilai) getNilai (){
	fmt.Println("ini hasil dari method getNilai")
	fmt.Println("nilai matkul", m.Matkul, "adalah", m.Nilai)
}

//method pointer
func (nilai *Nilai) changeNilai (newMatkul string, newNilai int){
	nilai.Matkul = newMatkul
	nilai.Nilai = newNilai
}

func main(){

	//array
	var member[3] string
	member[0] = "Haqim 007"
	member[1] = "Novi"
	member[2] = "Putri"

	array2 := [...] int {4,3,2,3,4}
	fmt.Println(member)
	fmt.Println(array2)

	fmt.Println("\n")

	//looping array 
	fmt.Println("ini menggunakan for loop")
	for i := 0; i < len(member); i++ {
		fmt.Println("data member ke-"+ strconv.Itoa(i) +" adalah "+member[i])
	}

	fmt.Println("\n")

	//range
	fmt.Println("ini menggunakan range")
	for key, value := range member {
		// fmt.Println("Member ke-"+strconv.Itoa(key)+" adalah "+value)
		fmt.Printf("Member ke-%d adalah %s\n", key, value)
	}

	fmt.Println("\n")

	// slice - sifatnya dapat mereplace data asli melalui slice
	fmt.Println("ini adalah slice")
	memberSlice := member[1:3]
	newMember := memberSlice
	newMember[0] = "haqim aja"
	fmt.Println("hasil slice dari key 1 sampai data ke 3(bukan key ke-3) adalah",memberSlice)
	fmt.Println("isi newMember adalah",newMember)

	fmt.Println("\n")

	//append -> menambah data array hasil slice
	fmt.Println("ini append")
	newMember = append(newMember, "rooney")
	fmt.Println(newMember)

	fmt.Println("\n")

	// menggunakan fungsi copy -> untuk mengcopy hasil slice kedalam variable baru
	fmt.Println("ini copy")
	member2 := make([] string, len(memberSlice))

	copy(member2, memberSlice)
	member2[0] = "Haqim member 2"
	fmt.Println(newMember)
	fmt.Println(member2)

	fmt.Println("\n")

	//map -> untuk membuat kumpulan data yang keynya bisa ditentukan sendiri dan harus unik
	fmt.Println("Ini map")
	var biodata = map[string]string{
		"nama" : "Haqim",
		"alamat" : "jakarta",
		"nohp" : "08151368668",
	} 

	fmt.Println(biodata)
	//delete map elemen
	fmt.Println("hapus elemen dengan key nohp")
	delete(biodata, "nohp")
	for key, value := range biodata{
		fmt.Println(key + " : "+ value)
	}
	// fmt.Println(biodata["nama"])


	fmt.Println("\n")


	biodata3 := models.StructBiodata{
		Nama : "Haqim",
		Alamat: "Jakarta",
		Umur : 22,
		Nohp : "08151368668",

	}

	fmt.Println(biodata3)


	//methods
	nilai := Nilai{
		Matkul : "MTSI",
		Nilai : 90,
	}
	
	nilai.getNilai()

	nilai.changeNilai("PDDP", 90)

	nilai.getNilai()

}