package math

import "fmt"

func Math(method string, num1 float64, num2 float64) string {

	fmt.Println("Calculator started")

	switch method {
	case "Multiply":
		fmt.Println("Multiplying")
		return fmt.Sprintf("%.2f * %.2f = %.2f", num1, num2, num1*num2)
	case "Add":
		fmt.Println("Adding")
		return fmt.Sprintf("%.2f + %.2f = %.2f", num1, num2, num1+num2)
	case "Subtract":
		fmt.Println("Subtracting")
		return fmt.Sprintf("%.2f - %.2f = %.2f", num1, num2, num1-num2)
	case "Divide":
		if num2 == 0 {
			return "Cannot divide by zero"
		}
		fmt.Println("Dividing")
		return fmt.Sprintf("%.2f / %.2f = %.2f", num1, num2, num1/num2)
	default:
		fmt.Println("Invalid input")
		return "Wrong input"
	}

}
