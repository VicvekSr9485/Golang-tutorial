package main
import ("fmt")

func calc(num1 int, num2 int) (int, int, int, int) {
	sum := num1 + num2
	diff := num1  - num2
	mul := num1 * num2
	div := num1 / num2
	return sum, diff, mul, div
}

func main() {
	a, b := 25, 5

	sum, diff, mul, div := calc(a, b)
	fmt.Println("sum", sum)
	fmt.Println("diff", diff)
	fmt.Println("mul", mul)
	fmt.Println("div", div)
}