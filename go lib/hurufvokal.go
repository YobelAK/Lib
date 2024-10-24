package main  
  
import (  
 "fmt"  
)  
  

func main() {
var kata string

fmt.Print("Masukkan kata: ")
fmt.Scan(&kata)

jumlahVokal := 0

for i:=0; i<len(kata); i++ {
	if kata[i] == 'a' || kata[i] == 'i' || kata[i] == 'u' || kata[i] == 'e' || kata[i] == 'o' || 
	   kata[i] == 'A' || kata[i] == 'I' || kata[i] == 'U' || kata[i] == 'E' || kata[i] == 'O' {
		jumlahVokal++
	}else if kata[i] == " " {
		fmt.Scanln(&kata)
	}
}

fmt.Println("Jumlah huruf vokal dalam kata", kata, "adalah", jumlahVokal)
}
