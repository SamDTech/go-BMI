package info

import "fmt"

const mainTitle = "BMI Calculator"
const separator = "---------------------------------------"
const WeightPrompt = "Please Enter your weight (kg): "
const HeightPrompt = "Please Enter your height (m): "

func PrintWelcome (){
	fmt.Println(mainTitle)
	fmt.Println(separator)
}

