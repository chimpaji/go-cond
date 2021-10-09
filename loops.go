package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func main() {
	selectedChoice,err := getUserChoice()

	if err != nil {
		fmt.Println("Invalid choice, exiting")
		return
	}

	if selectedChoice == "1" {
		calculateSumUpToNumber()
	} else if selectedChoice == "2" {
		calculateFactorial()
	} else if selectedChoice == "3" {
		calculateSumManually()
	}else if selectedChoice == "4" {
		calculateListSum()
	}


	
}

func calculateSumUpToNumber(){
	fmt.Print("Please enter your number:")
	chosenNumber, err := getInputNumber()

	if err != nil {
		fmt.Println(err)
		return
	}

	sum:= 0;

	for i:=1 ; i<=chosenNumber ; i++ {
		fmt.Println(i)
		sum+=i;
	}
	fmt.Printf("Result:%v",sum)
}

func calculateFactorial(){
	fmt.Print("Please enter your number:")
	chosenNumber, err := getInputNumber()

	if err != nil {
		fmt.Println(err)
		return
	}

	factorial:= 1;

	for i:=1 ; i<=chosenNumber ; i++ {
		fmt.Println(i)
		factorial=factorial*i;
	}
	fmt.Printf("Result:%v",factorial)
}

func calculateSumManually(){
	sum:=0
	isEnteringNumber := true
	for isEnteringNumber{
		choosenNumber, err := getInputNumber()
		sum+=choosenNumber
		fmt.Printf("Result:%v\n",sum)
		isEnteringNumber = err==nil
	}
}

func calculateListSum(){
	//1,2,3,4,5,6
	fmt.Print("Enter a list of numbers, seperate by comma:")
	input,err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return 
		
	}
	//1,2,3,4,5,6
	input = strings.Replace(input,"\r\n", "",-1)
	inputNumbers := strings.Split(input, ",")  //["1","2","3","4","5"]
	//loop thru inputNumber, conver each item to int, plus to sum
	sum:=0
	//to loop array, slice, map we can use; for (index,value) := range (slice) {}
	for _,value := range inputNumbers {
		number,err :=strconv.ParseInt(value,0,64)
		if err != nil {
			//instead of print err and stop execute the code, now we will move on this entered value
			//and start a new loop waiting for the next input using "continue" or break the loop using "break" and continue execute code in funciton
			continue
		}
		sum+=int(number)
	}
	fmt.Printf("Result:%v",sum)




}

func getInputNumber() (int, error){
	fmt.Print("Enter your input number:")
	inputNumber,err := reader.ReadString('\n')
	if err != nil {
		return 0,err

	}
	inputNumber = strings.Replace(inputNumber,"\r\n", "",-1)
	chosenNumber ,err := strconv.ParseInt(inputNumber,0,0)
	if err != nil {
		return 0,err //dont have to print err here, just return it, the main fn will handle the err
	}

	return int(chosenNumber), nil
}

func getUserChoice () (string, error) {
	fmt.Println("Please make your choice:")
	fmt.Println("1) Add up all the numbers to number  X")
	fmt.Println("2) Build the factorial up to number X")
	fmt.Println("3) Sum up manually entered nubmers")
	fmt.Println("4) Sum up a list of entered numbers")

	fmt.Print("Enter choice number:")

	userInput, inputErr  := reader.ReadString('\n')

	//if input error
	if inputErr != nil { 
		return "",inputErr
	}

	//clean up input
	userInput = strings.Replace(userInput, "\r\n", "", -1)

	//loop thru condition 
	if userInput == "1" || userInput == "2" || userInput == "3" || userInput == "4" {
		return userInput,nil
	} else {
		return "",errors.New("invalid input")
	}

}