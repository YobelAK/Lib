package main

import "fmt"

func main() {
	s := "Bbc_Abc"
	masuk :=0
	i:=0
	for(i<len(s)){
		if(i==0){
			fmt.Println(int(s[i]))
		}
		if(masuk==1){
			fmt.Println(int(s[i]))
			masuk=0
		}
		if(int(s[i])==95){
			masuk=1
		}
		i++
		
	}
}