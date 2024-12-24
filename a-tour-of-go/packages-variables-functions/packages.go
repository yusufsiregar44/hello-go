package main

// import packages
import (
	"fmt"
	"math/cmplx"
)

// Basic Go Types
var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

// constants
const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)

// iota
const (
	Sunday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

// function declaration
func add(x, y int) int {
	return x + y
}

// named return values
func swap(x, y string) (string, string) {
	return y, x
}

// naked return, ideally use in short functions
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum % 2
	return
}

// variable declaration without initialization
var q, w, e int
var ruby, diamond bool

// variable declaration with initialization
var c, python, java = true, false, "no!"

// short variable declaration can only be used inside a function
func shortVar() {
	i := 1
	j := 2
	k := "hello"
	a, b, c := 1, 2, "hey"
	fmt.Println(i, j, k)
	fmt.Println(a, b, c)
}

// zero values
func zeroValues() {
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}

// type conversion
func typeConversion() {
	i := 42.9999
	f := int(i)
	fmt.Println(i, f)
}

// type inference
func typeInference() {
	i := 42           // int
	f := 3.142        // float64
	g := 0.867 + 0.5i // complex128
	fmt.Printf("i is of type %T\n", i)
	fmt.Printf("f is of type %T\n", f)
	fmt.Printf("g is of type %T\n", g)
}

/*
on another note, some common go format identifiers are:
%T - prints the type of the value
%v - prints the value in its default format
%d - prints decimal integers
%f - prints floating-point numbers
%s - prints strings
%t - prints boolean values
%q - prints quoted strings
*/

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

// main function
func main() {
	// // fmt.Println("My favorite number is", rand.Intn(10))
	// // fmt.Println("This is Pi:", math.Pi)
	// // fmt.Println("This is addition:", add(1, 2))
	// // a, b := swap("hello", "world")
	// // fmt.Println("This is swap:", a, b)
	// // x, y := split(17)
	// // fmt.Println("This is split:", x, y)
	// // fmt.Println(q, w, e)
	// // fmt.Println(ruby, diamond)
	// // fmt.Println(c, python, java)
	// // fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	// // fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	// // fmt.Printf("Type: %T Value: %v\n", z, z)
	// // shortVar()
	// // zeroValues()
	// // typeConversion()
	// // typeInference()
	// fmt.Println(needInt(Small))
	// fmt.Println(needFloat(Small))
	// fmt.Println(needFloat(Big))
}
