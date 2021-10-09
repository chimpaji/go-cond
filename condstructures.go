package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your age:")
	userAgeInput, _ := reader.ReadString('\n')
	userAgeInput = strings.Replace(userAgeInput,"\r\n","",-1)
	userAge, err := strconv.ParseInt(userAgeInput,0,64) // 0 will infered base as the input, 64 is bit size

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