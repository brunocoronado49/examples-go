package main

import (
	"fmt"
	"time"
	"tutorial/constants"
	"tutorial/ifelse"
	"tutorial/loopfor"
	"tutorial/variables"
)

func main() {
	// Hello World in Go ===============
	fmt.Println("Hello, World!")

	// Values ===============
	fmt.Println("go" + "lang")

	fmt.Println("1 + 1 =", 1+1)
	fmt.Println("7.0/3.0 = ", 7.0/3.0)

	fmt.Println(true && true)
	fmt.Println(true || true)
	fmt.Println(!true)

	// Variables ===============
	variables.HelloVariables()

	// Constants ===============
	constants.HelloConstants()

	// For ===============
	loopfor.HelloFor()

	// If/Else ===============
	ifelse.HelloIfElse()

	// Switch
	o := 2
	fmt.Print("Write ", o, " as ")
	switch o {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("I'ts the weekend")
	default:
		fmt.Println("I'ts a weekday")
	}
}
