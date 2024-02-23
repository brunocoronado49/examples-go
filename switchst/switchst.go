package switchst

import (
	"fmt"
	"time"
)

func HelloSwitch() {
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

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'am a bool")
		case int:
			fmt.Println("I'am a int")
		default:
			fmt.Println("Don't knoe type %T\n", t)
		}
	}

	whatAmI(true)
	whatAmI(1)
	whatAmI("Hey")
}
