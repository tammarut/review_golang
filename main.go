package main

import (
	"fmt"

	"github.com/tammarut/review_golang/search"
)

func main() {
	fmt.Println("Hello, Go")

	//=> if-else
	if i := 1; i%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}

	//=> Loop
	for i := 1; i <= 10; i++ {
		fmt.Print(i, " ")
	}

	//=> Switch
	i := 3
	switch i {
	case 0:
		fmt.Println("\nZero")
	case 1:
		fmt.Println("\nOne")
	case 2:
		fmt.Println("\nTwo")
	case 3:
		fmt.Println("\nThree")
	case 4:
		fmt.Println("\nFour")
	default:
		fmt.Println("\nUnknown number")
	}

	//=> Type Declaration
	var message string
	message = "khunpleum"
	println(message)

	//=> Arrays
	var groupA [3]string
	groupA[0] = "Pleum"
	groupA[1] = "Tankwa"
	groupA[2] = "Dew"
	groupB := [3]string{"Chinya", "Nuna", "Tai"}
	fmt.Printf("\n %q \n %q \n", groupA, groupB)

	//=> Slice
	var groupC []string
	groupC = append(groupC, "Daw")
	groupC = append(groupC, "Ammi")
	groupC = append(groupC, "Not")
	for i, name := range groupC {
		fmt.Printf("Index: %d name: %q \n", i, name)
	}

	//=> Call function
	printFullName("Tom", "Holland")
	fmt.Println(getOwl())

	//=> Goroutines
	search.Run("dog")
}

func printFullName(fname, lname string) {
	fmt.Println(fname + " " + lname)
}

func getOwl() string {
	return "Black owl"
}
