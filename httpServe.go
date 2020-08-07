package main


import (
   
	"fmt"
	"net/http"
	"io/ioutil"
)

func main(){

	resp,_ := http.Get("http://www.baidu.com")

	bytes,_ := ioutil.ReadAll(resp.Body);

	fmt.Println(string(bytes))
}
