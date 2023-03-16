package main
import ("fmt")

// declare the structure named new
type new struct {
	name string
	age int
	gender string
	city string
}

//function which accepts variable of "new" type and prints new properties
func display(a new) {
	fmt.Println("name: ", a.name)
	fmt.Println("age: ", a.age)
	fmt.Println("gender: ", a.gender)
	fmt.Println("city: ", a.city)
}

// main function
func main() {
	// declares a variable newData of the type new
	var newData new
	// //assign values to members of newData
	newData.name = "Ronnie Coleman"
	newData.age = 45
	newData.gender = "Male"
	newData.city = "California"

	//declares and assign values to variable newData2 of type new
	newData2 := new{"Andra", 24, "Female", "Zurich"}

	//prints the member properties of newData and newData2 using display function
	fmt.Println("First member of the struct:")
	display(newData)
	fmt.Println("Second member of the struct:")
	display(newData2)
}
