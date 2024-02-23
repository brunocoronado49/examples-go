package main

import (
	"fmt"
	"tutorial/arrays"
	"tutorial/constants"
	"tutorial/ifelse"
	"tutorial/loopfor"
	"tutorial/switchst"
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
	switchst.HelloSwitch()

	// Arrays
	arrays.HelloArrays()
}
