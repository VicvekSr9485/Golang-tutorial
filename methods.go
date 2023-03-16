package main
import ("fmt")
 /*
 Note:
 - Go is not an object oriented language and it doesn’t have the concept of class.
 - Methods give a feel of what you do in object oriented programs where the functions
 - of a class are invoked using the syntax objectname.functionname()
 */
// declare the structure named new
type new struct {
	name string
	age int
	gender string
	city string
}

//Declaring a function with receiver of the type new
func(a new) display() {
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

	//prints the member properties of newData and newData2 using display method
	//Invoking the method using the receiver of the type new
    // syntax is variable.methodname()
	fmt.Println("First member of the struct:")

	newData.display()

	fmt.Println("Second member of the struct:")

	newData2.display()
}
