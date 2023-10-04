package main

import (
	"errors"
	"fmt"
)

// Varialbes
// Numbers
var var1 int

var var2 uint

var var3, var4 int //Multi declaration

var var5 float32
var var6 float64

// complex numbers a+bi
var var7 complex64
var var8 complex128

// string
var var9 string
var var10 string = "Helloo"

// boolean
var var11 bool

// Constants
const RED = 0

const (
	CAR = iota
	BUS
	TRAM
)

func main() {
	// constants

	fmt.Println(CAR)
	fmt.Println(BUS)
	fmt.Println(TRAM)

	// short declaration

	c := 10
	fmt.Println(c)

	// Type Conversion
	var3 = 10
	var5 = float32(var3)

	// multivariable
	a, b := 1, 2
	fmt.Printf("%v + %v = %v\n", a, b, a+b)
	fmt.Printf("%v * %v = %v\n", a, b, a*b)

	// Errors
	err := errors.New("hello")
	fmt.Println(err)
}
