package main
import ("fmt")

func main() {
	// create an empty array with a length of 12
	var firstArray [100] int 

	for i := 0; i < len(firstArray); i++ {
		firstArray[i] = i + 1 // assign a value to each element of the array
	}
	fmt.Println(firstArray)

	var firstSlice [] int = firstArray[3:10] // create a slice of firstArray
	fmt.Println(firstSlice)
	firstSliceAppend := append(firstSlice, 55)
	fmt.Println(firstSliceAppend)

	// create an empty array with a length of 52
	secondArray := make([]string, 52)

	// define variables for the ASCII codes of capital and small letters
	capitalStart := int('A')
	smallStart := int('a')

	for i := 0; i < 52; i++ {
		if i < 26 {
			secondArray[i] = string(capitalStart + i)
		} else {
			secondArray[i] = string(smallStart + i - 26)
		}
	}
	fmt.Println(secondArray)

	secondSlice := secondArray[4:15]
	fmt.Println(secondSlice)
	secondSliceAppend := append(secondSlice, "appended")
	fmt.Println(secondSliceAppend)
}
