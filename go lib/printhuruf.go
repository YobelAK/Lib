package main  
  
import (  
 "fmt"  
)  
  

func main() {
var kata string  
fmt.Print("Masukkan kata: ")
fmt.Scanln(&kata)
fmt.Printf("%c",kata[1])
}