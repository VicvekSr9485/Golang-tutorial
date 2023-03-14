package main
import ("fmt")

func main() {
	var a int
	a=5
	fmt.Println("a: ", a)

	var b int= 44
	fmt.Println("b: ", b)

	// declaring a variable without mentioning Type explicitly
	var c=42
	fmt.Println("c: ", c)

	// declaring multiple variables
	var i, j = 100, "Hello"
	fmt.Println("i and j: ", i, j)

	// declaring a variable without the var keyword
	k := "World"
	fmt.Println("k: ", k)

	// declaring a constant
	const d = 65
	fmt.Println("d: ", d)
}