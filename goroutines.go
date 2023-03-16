package main
import (
	"fmt"
	"time"
)

func display() {
	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println("display goroutine running")
	}
}

func main() {
	//invoking the goroutine display()
	go display()

	for j := 0; j < 10; j++ {
		time.Sleep(2 * time.Second)
		fmt.Println("Main program running")
	}
}
