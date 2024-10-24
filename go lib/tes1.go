package main
import"fmt"
func main(){
	var input,i int
	i=1
	fmt.Scanln(&input)
	faktor(input,i)
}

func faktor(input,i int){
	if i==input {
		fmt.Println(i)
	}else {
		if input % i ==0 {
			fmt.Println (i)
		}
		faktor(input,i+1)
	}
}