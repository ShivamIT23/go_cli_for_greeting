package main

import (
	"fmt"
	"go_cli_for_greeting/math"
	"strconv"

	"github.com/AlecAivazis/survey/v2"
)

func main() {
	prompt := &survey.Select{
		Message: "Choose a math operation:",
		Options: []string{"Add", "Subtract", "Multiply", "Divide", "Exit"},
	}

	var selectedOption string

	// 2. Ask the question.
	// The library handles the up/down arrows and the colors automatically.
	err := survey.AskOne(prompt, &selectedOption)

	if err != nil {
		fmt.Println("Selection failed:", err)
		return
	}

	if selectedOption == "Exit" {
		fmt.Println("Thank you for using us")
		return
	}

	var num1, num2 string

	survey.AskOne(&survey.Input{Message: "First number:"}, &num1)
	survey.AskOne(&survey.Input{Message: "Second number:"}, &num2)

	// 2. The Conversion (The "Real" Work)
	val1, err1 := strconv.ParseFloat(num1, 64)
	val2, err2 := strconv.ParseFloat(num2, 64)

	// 3. Defensive Programming: Check if the user typed "hello" instead of "10"
	if err1 != nil || err2 != nil {
		fmt.Println("\033[31mError: Please enter valid numbers!\033[0m")
		return
	}

	message := math.Math(selectedOption, val1, val2)
	fmt.Println(message)

}
