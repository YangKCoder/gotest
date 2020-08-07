package main

import (
	"fmt"
	"os/exec"
)

func main(){
	cmd := exec.Command("whoami")

	str,_ := cmd.Output()
	fmt.Println(string(str))

}
