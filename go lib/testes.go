package main
import "fmt"
func main(){
	var input int
	var luasling,luastab float64
	fmt.Scanln(&input)
	luaslingkaran(input,&luasling)
	luastabung(luasling,&luastab)
	fmt.Println(luasling,luastab)
	
}
func luaslingkaran(input int, luasling *float64){
	*luasling=3.14*float64(input)*float64(input)
}
func luastabung(luasling float64, luastab *float64){
	*luastab=luasling*5
}

