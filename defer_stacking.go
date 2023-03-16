package main
import ("fmt")

func display(x int) {
	fmt.Println(x)
}

func main() {
	// stacked functions execute in the Last In First Out(LIFO) order
	defer display(1)
    defer display(2)
    defer display(3)
    defer display(4)
    defer display(5)
	fmt.Println("This runs first")
}
