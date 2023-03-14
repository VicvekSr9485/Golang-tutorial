package main
import ("fmt")

func main() {
	a, b := 2, 3

	switch a * b {
		case 5:
            fmt.Println("multiple is 5")
        case 6:
            fmt.Println("multiple is 6")
        case 24:
            fmt.Println("multiple is 24")
        default:
            fmt.Println("Default Case")
	}
}