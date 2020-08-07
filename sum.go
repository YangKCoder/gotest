package main

import(

	"fmt"
	"os"
)


func main(){

	var count string
	
	for _,args := range os.Args[1:] {
	fmt.Println(args)
	
	count = count + args

}

	fmt.Println(count)
}
