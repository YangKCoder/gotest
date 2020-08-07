package main

import "fmt"

func main(){
	stringArr := []string{"java","golang"}
	
	for i,e := range stringArr {
		fmt.Printf("index: %d \t item: %s \n",i,e)
	}
}
