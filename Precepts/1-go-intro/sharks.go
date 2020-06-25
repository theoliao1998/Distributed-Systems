package main

import "fmt"

// Object oriented programming
// Convention: capitalize first letter of public fields
type Shark struct {
	Name string
	Age int
}

// Declare a public method
// This is called a receiver method
func (s *Shark) Bite() {
	fmt.Printf("%v says CHOMP!\n", s.Name)
}

// Because functions in Go are pass by value
// (as opposed to pass by reference), receiver
// methods generally take in pointers to the
// object instead of the object itself.
func (s *Shark) ChangeName(newName string) {
	s.Name = newName
}

// Receiver methods can take in other objects as well
func (s *Shark) Greet(s2 *Shark) {
	if (s.Age < s2.Age) {
		fmt.Printf("%v says your majesty\n", s.Name)
	} else {
		fmt.Printf("%v says yo what's up %v\n",
			s.Name, s2.Name)
	}
}

func sharks() {
	shark1 := Shark{"Bruce", 32}
	shark2 := Shark{"Sharkira", 40}
	shark1.Bite()
	shark1.ChangeName("Lee")
	shark1.Greet(&shark2) // pass in pointer
	shark2.Greet(&shark1)
}
