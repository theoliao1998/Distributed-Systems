// All files start with a package declaration
package main

// Import statements, one package on each line
import (
	"errors"
	"fmt"
)

// Main method will be called when the Go executable is run
func main() {
	fmt.Println("Hello world!")
	basic()
	add(1, 2)
	divide(3, 4)
	loops()
	// Data structures
	slices()
	maps()
	// Object oriented programming
	sharks()
	// Goroutines and channels
	goroutines()
	channels()
	bufferedChannels()
}

// Function declaration
func basic() {
	// Declare x as a variable, initialized to 0
	var x int
	// Declare y as a variable, initialized to 2
	var y int = 2
	// Declare z as a variable, initialized to 4
	// This syntax can only be used in a function
	z := 4

	// Assign values to variables
	x = 1
	y = 2
	z = x + 2 * y + 3

	// Print the variables; just use %v for most types
	fmt.Printf("x = %v, y = %v, z = %v\n", x, y, z)
}

// Function declaration; takes in 2 ints and outputs an int
func add(x, y int) int {
	return x + y
}


// Function that returns two things; error is nil if successful
func divide(x, y int) (float64, error) {
	if y == 0 {
		return 0.0, errors.New("Divide by zero")
	}
	// Cast x and y to float64 before dividing
	return float64(x) / float64(y), nil
}

func loops() {
	// For loop
	for i := 0; i < 10; i++ {
		fmt.Print(".")
	}
	// While loop
	sum := 1
	for sum < 1000 {
		sum *= 2
	}
	fmt.Printf("The sum is %v\n", sum)
}

func slices() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(slice)
	fmt.Println(slice[2:5]) // 3, 4, 5
	fmt.Println(slice[5:]) // 6, 7, 8
	fmt.Println(slice[:3]) // 1, 2, 3
	slice2 := make([]string, 3)
	slice2[0] = "tic"
	slice2[1] = "tac"
	slice2[2] = "toe"
	fmt.Println(slice2)
	slice2 = append(slice2, "tom")
	slice2 = append(slice2, "radar")
	fmt.Println(slice2)
	for index, value := range slice2 {
		fmt.Printf("%v: %v\n", index, value)
	}
	fmt.Printf("Slice length = %v\n", len(slice2))
}

func maps() {
	myMap := make(map[string]int)
	myMap["yellow"] = 1
	myMap["magic"] = 2
	myMap["amsterdam"] = 3
	fmt.Println(myMap)
	myMap["magic"] = 100
	delete(myMap, "amsterdam")
	fmt.Println(myMap)
	fmt.Printf("Map size = %v\n", len(myMap))
}
