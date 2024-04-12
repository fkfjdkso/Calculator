package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	outOfRange             = "Введенное число не принадлежит отрезку от 1 до 10"                                                //ПРОВЕРЕНО
	nonIntInput            = "Введенное число(числа) не является(являются) целым(и)"                                            //ПРОВЕРЕНО
	romanianNegativeOrZero = "Римская система счисления не имеет числа 0 и отрицательных чисел"                                 //ПРОВЕРЕНО
	numberSystemMixed      = "Используются разные системы счисления"                                                            //ПРОВЕРЕНО
	tooManyOrWrongActions  = "Формат математической операции не удовлетворяет заданию или не является математической операцией" //ПРОВЕРЕНО
)

var operators = map[string]func(x, y int) int{
	"+": func(x, y int) int { return x + y },
	"-": func(x, y int) int { return x - y },
	"*": func(x, y int) int { return x * y },
	"/": func(x, y int) int { return x / y },
}

var romanian = map[string]int{
	"I":    1,
	"II":   2,
	"III":  3,
	"IV":   4,
	"V":    5,
	"VI":   6,
	"VII":  7,
	"VIII": 8,
	"IX":   9,
	"X":    10,
	"XL":   40,
	"L":    50,
	"XC":   90,
	"C":    100,
}

func checkInput(x string, y string, z string) {
	fmt.Println(z)
	digit1, errDigit1 := strconv.Atoi(x)
	digit2, errDigit2 := strconv.Atoi(y)
	if errDigit1 != nil && errDigit2 != nil {
		if romanian[x] == 0 || romanian[y] == 0 {
			panic(outOfRange)
		}
	} else if errDigit1 == nil && errDigit2 == nil {
		if (digit1 >= 10 || digit1 < 1) || (digit2 >= 10 || digit2 < 1) {
			panic(outOfRange)
		}
	} else if (errDigit1 == nil && romanian[y] != 0) || (errDigit2 == nil && romanian[x] != 0) {
		panic(numberSystemMixed)
	} else if errDigit1 != nil || errDigit2 != nil {
		_, errFloat1 := strconv.ParseFloat(x, 64)
		_, errFloat2 := strconv.ParseFloat(y, 64)
		if errFloat1 == nil || errFloat2 == nil {
			panic(nonIntInput)
		}
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)
		elems := strings.Fields(text)
		switch len(elems) {
		case 3:
			if elems[1] != "+" && elems[1] != "-" && elems[1] != "/" && elems[1] != "*" {
				panic(tooManyOrWrongActions)
			} else if elems[1] == "-" && romanian[elems[0]] <= romanian[elems[2]] {
				panic(romanianNegativeOrZero)
			} else {
				checkInput(elems[0], elems[2], elems[1])
			}
		default:
			panic(tooManyOrWrongActions)
		}

	}
}
