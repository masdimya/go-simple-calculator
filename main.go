package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

func instruction() {
	fmt.Println("1. Type Integers")
	fmt.Println("2. Type operator \"+\" for sum")
	fmt.Println("3. Type operator \"-\" for sub")
	fmt.Println("4. Type operator \"x\" for multiply")
	fmt.Println("5. Type operator \"/\" for devide")
	fmt.Println("6. Type operator \"=\" for get total")
	fmt.Println("7. Press enter for show the result")
}

func display(value string){
	fmt.Print("Display :", value)
}

func calculate(arrInput *[]string,input string)  {
	if input != "=" {
		*arrInput = append(*arrInput, input)
	} else {
		var total int = 0
		var operator string = ""

		for i,v := range *arrInput {
			if i == 0 {
				parseInt, _ := strconv.Atoi(v)
				total = parseInt
			} else if v != "=" {
				switch v {
				case "+","-","x","/":
					operator = v
				default:
					if operator != "" {
						parseInt, _ := strconv.Atoi(v)
						switch operator {
						case "+":
							total += parseInt
						case "-":
							total -= parseInt
						case "*":
							total = total * parseInt
						case "/":
							total = total / parseInt
						}

						operator = ""

					}
				}

			}


		}

		*arrInput = []string{ strconv.Itoa(total) }

	}
	
}

func main()  {
	var inputStore = []string{}
	var input string

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	
	instruction()
	display("")

	for scanner.Scan() {
		input  = scanner.Text()
		calculate(&inputStore, input)
		display(strings.Join(inputStore,""))
	}

}