package main

import (
	"fmt"

	"golang.org/x/exp/slices"
)

func main() {
	fmt.Println("Slice samples in golang")

	// fixed size collection of item
	// can't change the size after the declaration

	// Eg 1

	var arr1 []string //initial value will be empty string

	arr1 = []string{"Coffee", "Tea", "Other"}
	fmt.Println("arr1: ", arr1)

	// Eg 2

	arr2 := []string{"car", "Bus", "Tram"}
	fmt.Println("arr2: ", arr2)

	// Eg 3

	arr3 := arr2 //this is a copy operation
	fmt.Println("arr3: ", arr3)

	arr2[1] = "train"
	fmt.Println("arr2: ", arr2)
	fmt.Println("arr3: ", arr3)

	s := []string{"a", "b", "c"}
	fmt.Println(s)
	s = append(s, "d", "e")
	fmt.Println(s)

	s = slices.Delete(s, 1, 2)
	fmt.Println(s)
}
