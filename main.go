package main

//untuk mengelompokkan beberapa file dalam satu paket
//kenapa namanya main? di golang ada tipe package yaitu executable hanya mengekesuksi dengan command line
//lalu ada library package nya` bebas selain nama "main"
//pemanggilan import di golang untuk menjadi standard library, pemanggilan import untuk memanggil package yang berbeda
//pemanggilan import jika mendapatkan code dari sumber lain
//func main sebuah method khusus yang akan dijalankan ketika kita run untuk package main executable yang default nya dari golang

import (
	"fmt"
	"pertama/calculation"
)

func main() {
	fmt.Println("Halo, belajar golang")
	result := calculation.Add(9, 9)
	result2 := calculation.Multiply(9, 7)
	fmt.Println(result)
	fmt.Println(result2)
	var name string = "Ruby on Rails"
	// var age int
	//age = 10
	age := 20
	if age > 10 {
		fmt.Println("Boleh main game")
	} else {
		fmt.Println("Kamu belum boleh")
	}

	score := 65
	var grade string

	if score <= 50 {
		grade = "E"
	} else if score <= 60 {
		grade = "D"
	} else if score < 70 {
		grade = "C"
	} else {
		grade = "NULL"
	}
	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(grade)

}
