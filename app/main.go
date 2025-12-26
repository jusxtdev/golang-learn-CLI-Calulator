package main

import (
	"flag"
	"fmt"
	"math"
	"math/rand/v2"
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

func divide_func(x, y int) float64 {
	return float64(x) / float64(y)
}

func motivate_func() {
	quotes := map[int]string{
		1: `"Mathematics is not about numbers, equations, computations, or algorithms: it is about understanding." - William Paul Thurston`,
		2: `"Pure mathematics is, in its way, the poetry of logical ideas." - Albert Einstein`,
		3: `"Mathematics knows no races or geographic boundaries; for mathematics, the cultural world is one country." - David Hilbert`,
		4: `"The only way to learn mathematics is to do mathematics." - Paul Halmos`,
		5: `"In mathematics, you don't understand things. You just get used to them." - John von Neumann`,
	}
	max := 5
	min := 0
	quote_num := rand.IntN(max-min+1) + min
	fmt.Println(quotes[quote_num])
}

func int_input() int {
	fmt.Println("Enter number :")
	var x int
	fmt.Scanln(&x)
	return x
}

func main() {

	//define flags
	add := flag.Bool("add", false, "Add two numbers")
	subtract := flag.Bool("subtract", false, "Subtract two numbers")
	multiply := flag.Bool("multiply", false, "Multiply two numbers")
	floor := flag.Bool("floor", false, "floor value of division of two numbers")
	divide := flag.Bool("divide", false, "Division of two numbers")
	help := flag.Bool("help", false, "Usage help")
	motivate := flag.Bool("motivate", false, "Get Motivated !!")

	//Parse flags
	flag.Parse()

	math_opr := *add || *subtract || *multiply || *floor || *divide

	if math_opr {
		// math operations

		// Take user input
		x := int_input()
		y := int_input()
		result := 0.0

		if *add {
			result = float64(add_func(x, y))
		} else if *subtract {
			result = float64(subtract_func(x, y))
		} else if *multiply {
			result = float64(multiply_func(x, y))
		} else if *floor {
			result = float64(floor_func(x, y))
		} else if *divide {
			result = float64(divide_func(x, y))
		}
		fmt.Println("Answer", math.Round(result*100)/100)
	} else if *help {
		//help statemtn
		help_str := `
Usage :
	./main [options]
Options :
	-add      : 	Perform addition of two nums
	-subtract :   Perform subtraction of two nums
	-multiply :   Perform multiplication of two nums
	-floor    :   Perform division of two nums and print it's floor value
	-division :   Perform division
	-help     :   This option
	-motivate :   Get Motivated !!
`
		fmt.Println(help_str)
	} else if *motivate {
		// motivate
		motivate_func()
	} else {
		fmt.Println("No flag Provided\nTry -help")
	}
}
