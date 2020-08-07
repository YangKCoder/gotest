package main

import "fmt"


type Phone interface{
	call()
}


type Huawei struct {
}

func (huawei Huawei) call(){
	fmt.Println("I am Huawei")
}


func main(){
	var phone Phone
	
	phone = new(Huawei)

	phone.call()

}
