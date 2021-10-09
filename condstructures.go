package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	userAge,err := getUserAge() //this custom fn return 2 value, we'd like to keep err in main fn
	//check if text can be convert to int
	if err!=nil {
		fmt.Println(err)
		return //this will exit fn main execution, no more code execute
	}

	if (userAge >= 30 && userAge < 50)  || userAge > 60 {    // == != > >= < <=  
		fmt.Println("You are eligible for senior position")
	} else if userAge >= 20{
		fmt.Println("Welcome to the club")
	} else {
		fmt.Println("Too young?")
	}

	fmt.Println("I'll always be printed")

}

func getUserAge()(int64 , error){

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your age:")
	userAgeInput, _ := reader.ReadString('\n')
	userAgeInput = strings.Replace(userAgeInput,"\r\n","",-1)
	userAge, err := strconv.ParseInt(userAgeInput,0,64) // 0 will infered base as the input, 64 is bit size

	return userAge ,err //sent err to main fn to do condition check
}