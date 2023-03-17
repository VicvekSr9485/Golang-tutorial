package main
import (
	"fmt"
	"os"
	"errors"
)

//function accepts a filename and tries to open it.
func fileOpen(filename string) (string, error) {
	f, err := os.Open(filename)

	//er will be nil if the file exists else it returns an error object
	if err != nil {
		//created a new error object and returns it
		return "", errors.New("Custom error message: File name is wrong")
	} else {
		return f.Name(), nil
	}
}

func main() {
	//receives custom error or nil after trying to open the file
	name, error := fileOpen("invalid.txt")
	if error != nil {
        fmt.Println(error)
    }else{
    	fmt.Println("file opened", name)
    }  
}
