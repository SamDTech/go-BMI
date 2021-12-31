package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/samdtech/bmi/info"
)

var reader = bufio.NewReader(os.Stdin)

func getUserMetrics() (float64, float64) {
	weight := getUserInput(info.WeightPrompt)

	height := getUserInput(info.HeightPrompt)

	return weight, height
}

func getUserInput(promptText string) float64 {
	fmt.Print(promptText)

	userInput, _ := reader.ReadString('\n')
	// remove the \n in weightInput
	userInput = strings.Replace(userInput, "\n", "", -1)
	input, _ := strconv.ParseFloat(userInput, 64)

	return input
}
