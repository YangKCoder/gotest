package main

import "fmt"


func compare( x,y int) int{

	if(x>y){
	  return x;
	}else {
	 return y;
	}
}


func main(){
	var x,y = 10 ,20

	fmt.Println(compare(x,y));

}
