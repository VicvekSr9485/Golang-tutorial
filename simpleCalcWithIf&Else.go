package main
import ("fmt")

func main() {
	var num1, num2 float64
	var operator string

	fmt.Print("Enter first number: ", num1)
	fmt.Scanln(&num1)

	fmt.Print("Enter second number: ", num2)
	fmt.Scanln(&num2)

	fmt.Print("Enter an operator of any of these types (+, -, *, /): ", operator)
	fmt.Scanln(&operator)

	if operator == "+" {
		fmt.Printf("%.2f + %.2f = %.2f\n", num1, num2, num1+num2)
	} else if operator == "-" {
		fmt.Printf("%.2f - %.2f = %.2f\n", num1, num2, num1-num2)
	} else if operator == "*" {
		fmt.Printf("%.2f * %.2f = %.2f\n", num1, num2, num1*num2)
	} else if operator == "/" {
		if num2 == 0 {
			fmt.Println("Error: You cannot divide by 0")
		} else {
			fmt.Printf("%.2f / %.2f = %.2f\n", num1, num2, num1/num2)
		}
	} else {
		fmt.Println("Error: Invalid operator")
	}
}