package main
import "fmt"
func main () {
var berat,totala,totalb,totalc,totald float64
var grade string
fmt.Scanln(&berat,&grade)
	for berat>0 && (grade=="A" || grade=="B" ||grade=="C"||grade=="D") {
		if grade=="A" {
		totala=totala+berat
		}else if grade=="B" {
		totalb=totalb+berat
		}else if grade=="C" {
		totalc=totalc+berat
		}else if grade=="D" {
		totald=totald+berat
		}
		fmt.Scanln(&berat,&grade)
	}
	fmt.Println("berat yg disalurkan:", totala+totalb+totalc)
	fmt.Println("berat grade a:", totala)
	fmt.Println("berat grade b:", totalb)
	fmt.Println("berat grade c:", totalc)
	fmt.Println("berat yg tidak disalurkan:", totald)
	
}