package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)

func main(){
	// Output Information
	fmt.Println("BMI Calculator")
	fmt.Println("---------------------------------------")

	// Prompt user for input
	fmt.Print("Please Enter your weight (kg): ")

	weightInput, _ := reader.ReadString('\n')

	fmt.Print("Please Enter your height (m): ")
	

	heightInput, _ := reader.ReadString('\n')

	fmt.Print(weightInput)
	fmt.Print(heightInput)
}