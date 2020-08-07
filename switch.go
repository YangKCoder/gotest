package main

import "fmt"


func main(){

	var grade string = "B"
	var marks int = 90


	switch marks {
		case 90: grade = "A"
		case 80: grade = "B"
		case 50,60,70: grade = "C"
		default: grade = "D"
	}

	switch {
		case grade == "A":
			fmt.Printf("优秀\n")
		case grade == "B":
			fmt.Printf("良好\n")
		case grade == "C":
			fmt.Printf("及格\n")
		case grade == "D":
			fmt.Printf("差\n")
	}

	fmt.Printf("your grade is %s\n",grade)



}
