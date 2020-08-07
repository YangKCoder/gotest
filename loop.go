package main

import (
	"fmt"
	"os"
)


func main(){
	for line,args :=range os.Args[1:] {
		fmt.Printf("%d \t %s \n",line,args)

}
}
