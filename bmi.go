package main

import (
	"fmt"
	"strconv"
	"strings"
)

const mainTitle = "BMI Calculator"
const separator = "---------------------------------------"
const weightPrompt = "Please Enter your weight (kg): "
const heightPrompt = "Please Enter your height (m): "

func main() {
	// Output Information
	fmt.Println(mainTitle)
	fmt.Println(separator)

	// Prompt user for input
	fmt.Print(weightPrompt)

	weightInput, _ := reader.ReadString('\n')

	fmt.Print(heightPrompt)

	heightInput, _ := reader.ReadString('\n')

	// remove the \n in weightInput
	weightInput = strings.Replace(weightInput, "\n", "", -1)

	// remove the \n in heightInput
	heightInput = strings.Replace(heightInput, "\n", "", -1)

	weight, _ := strconv.ParseFloat(weightInput, 64)
	height, _ := strconv.ParseFloat(heightInput, 64)

	// calculate the bmi with the formula: weight / (height * height)

	bmi := weight / (height * height)

	fmt.Printf("Your BMI is %.2f\n", bmi)
}
