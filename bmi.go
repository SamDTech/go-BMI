package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/samdtech/bmi/info"
)



func main() {
	// Output Information
	fmt.Println(info.MainTitle)
	fmt.Println(info.Separator)

	// Prompt user for input
	fmt.Print(info.WeightPrompt)

	weightInput, _ := reader.ReadString('\n')

	fmt.Print(info.HeightPrompt)

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
