package main

import (
	"fmt"
	"os"
	"log"
)

func main(){
	cmd := os.Args[1:]
	fmt.Println(cmd);
	if(len(cmd)==0){
		log.Fatal("error ....");	
	}else if(cmd[0] == "start"){
		fmt.Println("server start ...")
}



}
