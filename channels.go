package main
import (
	"fmt"
	"time"
)

func display(ch chan string) {
	time.Sleep(2 * time.Second)
	fmt.Println("Normal goroutine started")
	ch <- "Hello"
}

func main() {
	ch := make(chan string)
	go display(ch)
	x := <-ch

	fmt.Println("Main function started")
	fmt.Println("Printing x in main() after taking from channel:", x)

}
