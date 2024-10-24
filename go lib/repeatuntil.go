package main

import "fmt"

func main() {
	var n, iterasi int
	var berhenti bool
	var result1 bool

	fmt.Scanln(&n)
	iterasi = 0
	berhenti = false
	for !berhenti {
		iterasi = iterasi + 1
		result1 = (n % iterasi) == 0
		fmt.Println(iterasi, result1)
		berhenti = iterasi == n
	}
}
