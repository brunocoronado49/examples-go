package main

import (
	"fmt"
	"math"
)

// Constants ===============
const s string = "constant"

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
	var a string = "Initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	f := "apple"
	fmt.Println(f)

	// Constants ===============
	fmt.Println(s)

	const n = 500000
	const z = 3e20 / n
	fmt.Println(z)

	fmt.Println(int64(z))
	fmt.Println(math.Sin(n))

	// For ===============
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	for j := 0; j < 3; j++ {
		fmt.Println(j)
	}

	for i := range 3 {
		fmt.Println("range: ", i)
	}

	for {
		fmt.Println("loop")
		break
	}

	for n := range 6 {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}

	// If/Else ===============

}
