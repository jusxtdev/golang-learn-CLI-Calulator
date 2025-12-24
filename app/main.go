package main

import (
	"flag"
	"fmt"
)

// take operation as flag input
// select function acc to input, selection happens in main func
// flags
//    -add, -subtract, -multiply, -divide, -help
//    -motivate(name) -- sends a motivational quote to the user to help them study maths

func add_func(x, y int) int {
	return x + y
}

func subtract_func(x, y int) int {
	return x - y
}

func multiply_func(x, y int) int {
	return x * y
}

func floor_func(x, y int) int {
	return x / y
}

func main() {

	//define flags
	add := flag.Bool("add", false, "Add two numbers")
	subtract := flag.Bool("subtract", false, "Subtract two numbers")
	multiply := flag.Bool("multiply", false, "Multiply two numbers")
	floor := flag.Bool("floor", false, "floor value of division of two numbers")
	//help := flag.Bool("help", false, "Usage help")
	//motivate := flag.String("name", "Young Man", "Get Motivated !!")

	options := map[int]func(int, int) int{
		1: add_func,
		2: subtract_func,
		3: multiply_func,
		4: floor_func,
	}

	result := 0
	if *add {
		result = options[1](1, 2)
	} else if *subtract {
		result = options[2](1, 2)
	} else if *multiply {
		result = options[3](1, 22)
	} else if *floor {
		result = options[4](23, 9)
	}
	fmt.Println("Answer", result)
}
