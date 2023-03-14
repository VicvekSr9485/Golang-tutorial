package main
import ("fmt")

func main() {
	var i = 25
	for i = 1; i <=  25; i++ {
		fmt.Println("i: ", i)
	}

	if i > 20 {
		fmt.Println("i is greater than 20 ")
	} else if i < 10 {
		fmt.Println("i is less than 10 ")
	}
}