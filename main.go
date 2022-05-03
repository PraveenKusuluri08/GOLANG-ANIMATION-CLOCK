package main

import (
	"fmt"
	"time"

	"github.com/inancgumus/screen"
)

func main() {
	type Placeholder [5]string

	var zero = Placeholder{
		"███",
		"█ █",
		"█ █",
		"█ █",
		"███",
	}

	var one = Placeholder{
		"██ ",
		" █ ",
		" █ ",
		" █ ",
		"███",
	}

	var two = Placeholder{
		"███",
		"  █",
		"███",
		"█  ",
		"███",
	}

	var three = Placeholder{
		"███",
		"  █",
		"███",
		"  █",
		"███",
	}

	var four = Placeholder{
		"█ █",
		"█ █",
		"███",
		"  █",
		"  █",
	}

	var five = Placeholder{
		"███",
		"█  ",
		"███",
		"  █",
		"███",
	}

	var six = Placeholder{
		"███",
		"█  ",
		"███",
		"█ █",
		"███",
	}

	var seven = Placeholder{
		"███",
		"  █",
		"  █",
		"  █",
		"  █",
	}

	var eight = Placeholder{
		"███",
		"█ █",
		"███",
		"█ █",
		"███",
	}

	var nine = Placeholder{
		"███",
		"█ █",
		"███",
		"  █",
		"███",
	}

	var Digits = [...]Placeholder{
		zero, one, two, three, four, five, six, seven, eight, nine,
	}

	var colon = Placeholder{
		"   ",
		" ░ ",
		"   ",
		" ░ ",
		"   ",
	}
	screen.Clear()
	for {
		screen.MoveTopLeft()

		now := time.Now()
		hour, min, sec := now.Hour(), now.Minute(), now.Second()

		fmt.Printf("hour: %d, min: %d, sec: %d\n", hour, min, sec)

		clock := [...]Placeholder{
			Digits[hour/10], Digits[hour%10],
			colon,
			Digits[min/10], Digits[min%10],
			colon,
			Digits[sec/10], Digits[sec%10],
		}
		for line := range clock[0] {

			for index, digit := range clock {
				next := clock[index][line]
				if digit == colon && sec%2 == 0 {
					next = "   "
				}
				fmt.Print(next, " ")
			}
			fmt.Println()
		}
		time.Sleep(time.Second)
	}
}
