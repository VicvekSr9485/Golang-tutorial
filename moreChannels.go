package main
import (
	"fmt"
	"time"
)

//This subroutine pushes numbers 0 to 9 to the channel and closes the channel
func addToChannel(ch chan int) {
	fmt.Println("Send data")

	for i := 0; i < 10; i++ {
		ch <- i // pushing data to channel
	}
	close(ch)
}

//This subroutine fetches data from the channel and prints it.
func fetchFromChannel(ch chan int) {
	fmt.Println("Read data")

	for {
		x, flag := <-ch	//fetch data from channel
		//flag is true if data is received from the channel
		//flag is false when the channel is closed

		if flag == true {
			fmt.Println(x)
		} else {
			fmt.Println("Empty channel")
			break
		}
	}
}

func main() {
	//creating a channel variable to transport integer values
	ch := make(chan int)

	//invoking the subroutines to add and fetch from the channel
	//These routines execute simultaneously
	go addToChannel(ch)
	go fetchFromChannel(ch)

	//delay is to prevent the exiting of main() before goroutines finish
	time.Sleep(5 * time.Second)
	fmt.Println("Main block running")
}
