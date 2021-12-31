package main

import (
	"github.com/samdtech/bmi/info"
)

func main() {
	// Output Information
	info.PrintWelcome()

	weight, height := getUserMetrics()

	bmi := calculateBmi(weight, height)

	outputResult(bmi)
}

func calculateBmi(weight float64, height float64) float64 {
	return weight / (height * height)
}
