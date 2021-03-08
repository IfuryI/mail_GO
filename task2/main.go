package main

import (
	"fmt"
	"flag"
	"unicode"
	"strconv"
	"unsafe"
)


func printHelp() {
	fmt.Println("Введены некорректные флаги!\n" +
				 "Пример использования: " +
				 "go run calc.go \"(1+2)*3\"")
}


func readExpressionFlags() string {
	flag.Parse()
	switch len(flag.Args()) {
		case 1:
			return flag.Args()[0]

		default:
			printHelp();
			return ""
	}

	return ""
}


type expression struct {
	token string
	args []expression
}


type parser struct {
	input string
}


func (p *parser) parseToken() string{
	// Пропуск пробелов
	for _, symb := range (*p).input {
		if !unicode.IsSpace(symb) {
			break
		}
		*(*uintptr)(unsafe.Pointer(&(*p).input))++
	}

	// Получение числа
	if unicode.IsDigit(rune((*p).input[0])) {
		var number string
		for _, symb := range (*p).input {
			if !unicode.IsDigit(symb) && symb != '.' {
				break
			}
			number = number + string(symb)
			*(*uintptr)(unsafe.Pointer(&(*p).input))++
		}

		return number
	}

	// Получение знака или скобок
	tokens := []string{ "+", "-", "*", "/", "(", ")" }
	for _, tok := range tokens {
		if string((*p).input[0]) == tok {
			for i := 0; i < len(tok); i++ {
				(*(*uintptr)(unsafe.Pointer(&(*p).input)))++
			}

			return tok
		}
	}

	return ""
}


func (p *parser) parseSimpleExpression() expression {
	token := (*p).parseToken()
	if token == "" {
		fmt.Println("Некорректное выражение!")
	}

	if token == "(" {
		result := (*p).mainParse()
		if (*p).parseToken() != ")" {
			fmt.Println("Ожидается: \")\"")
		}
		return result
	}

	if unicode.IsDigit(rune(token[0])) {
		return expression{token, []expression{}}
	}

	return expression{token, []expression{(*p).parseSimpleExpression()}}
}


func getPriority(binary_op string) int {
	switch binary_op {
		case "+":
			return 1
		case "-":
			return 1
		case "*":
			return 2
		case "/":
			return 2
		default:
			return 0
	}
}


func (p *parser) parseBinaryExpression(minPriority int) expression {
	leftExpr := (*p).parseSimpleExpression()

	for {
		op := (*p).parseToken()
		priority := getPriority(op)
		if (priority <= minPriority) {
			(*(*uintptr)(unsafe.Pointer(&(*p).input)))--
			return leftExpr
		}

		rightExpr := (*p).parseBinaryExpression(priority)
		leftExpr = expression{op, []expression{leftExpr, rightExpr}}
	}
}


func (p *parser) mainParse() expression {
	return (*p).parseBinaryExpression(0)
}


func eval(expr expression) float64 {
	switch len(expr.args) {
		case 2:
			a := eval(expr.args[0])
			b := eval(expr.args[1])
			if expr.token == "+" {return a + b}
			if expr.token == "-" {return a - b}
			if expr.token == "*" {return a * b}
			if expr.token == "/" {return a / b}

			fmt.Println("Неизвестная операция!")
		case 1:
			a := eval(expr.args[0])
			if expr.token == "+" {return +a}
			if expr.token == "-" {return -a}

			fmt.Println("Неизвестная операция!")

		case 0:
			res, err := strconv.ParseFloat(expr.token, 64)

			if err != nil {
				fmt.Println("ОШИБКА")
			}

			return res
	}

	return 0.0
}


func main() {
	expression := readExpressionFlags()

	fmt.Println(expression)

	myParser := parser{expression}

	fmt.Println(eval(myParser.mainParse()))
}
