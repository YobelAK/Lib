package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("##  Program Go Menghitung Jumlah Huruf Vokal  ##")
	fmt.Println("=================================================")

	var x string
	vokal := 0

	fmt.Print("Input kata / kalimat: ")
	fmt.Scanln(&x)
	fmt.Println()

	// Hitung jumlah huruf vokal
	for _, huruf := range strings.ToLower(x) {
		if huruf == 'a' || huruf == 'i' || huruf == 'u' || huruf == 'e' || huruf == 'o' {
			vokal++
		}
	}

	// Tampilkan total huruf vokal jika ditemukan
	if vokal > 0 {
		fmt.Printf("Jumlah huruf vokal = %d\n", vokal)
	} else {
		fmt.Println("Huruf vokal tidak ditemukan")
	}
}