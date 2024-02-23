package constants

import (
	"fmt"
	"math"
)

// Constants ===============
const s string = "constant"

func HelloConstants() {
	fmt.Println(s)

	const n = 500000
	const z = 3e20 / n
	fmt.Println(z)

	fmt.Println(int64(z))
	fmt.Println(math.Sin(n))
}
