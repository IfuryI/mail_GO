package main

import (
    "fmt"
    "flag"
    "unicode"
    "strconv"
    "strings"
    "unsafe"
    "errors"
)


func main() {
    expr, err := readExpressionFlags()
    if err != nil {
        fmt.Println(err)
        return
    }

    res, err := calculate(expr)
    if err != nil {
        fmt.Println(err)
        return
    }

    fmt.Println(res)
}


func calculate(expr string) (float64, error) {
    err := checkExpression(expr)
    if err != nil {
        return 0, err
    }

    myParser := parser{expr}
    parsedExpression, err := myParser.mainParse()
    if err != nil {
        return 0, err
    }

    res, err := eval(parsedExpression)
    if err != nil {
        return 0, err
    }

    return res, nil
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


func (p *parser) parseSimpleExpression() (expression, error) {
    token := (*p).parseToken()
    if token == "" {
        return expression{}, errors.New("Некорректное выражение!")
    }

    if token == "(" {
        result, err := (*p).mainParse()
        if err != nil {
            return expression{}, err
        }
        if (*p).parseToken() != ")" {
            return expression{}, errors.New("Ожидается: \")\"!")
        }
        return result, nil
    }

    if unicode.IsDigit(rune(token[0])) {
        return expression{token, []expression{}}, nil
    }

    tempResult, err := (*p).parseSimpleExpression()
    if err != nil {
        return expression{}, err
    }

    return expression{token, []expression{tempResult}}, nil
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


func (p *parser) parseBinaryExpression(minPriority int) (expression, error) {
    leftExpr, err := (*p).parseSimpleExpression()
    if err != nil {
        return expression{}, err
    }

    for {
        op := (*p).parseToken()
        priority := getPriority(op)
        if (priority <= minPriority) {
            (*(*uintptr)(unsafe.Pointer(&(*p).input)))--
            return leftExpr, nil
        }

        rightExpr, err := (*p).parseBinaryExpression(priority)
        if err != nil {
            return expression{}, err
        }
        leftExpr = expression{op, []expression{leftExpr, rightExpr}}
    }

    return expression{}, errors.New("Некорректное выражение")
}


func (p *parser) mainParse() (expression, error) {
    res, err := (*p).parseBinaryExpression(0)
    if err != nil {
        return expression{}, err
    }

    return res, nil
}


func eval(expr expression) (float64, error) {
    switch len(expr.args) {
        case 2:
            a, err := eval(expr.args[0])
            if err != nil {
                return 0, err
            }

            b, err := eval(expr.args[1])
            if err != nil {
                return 0, err
            }

            if expr.token == "+" {return a + b, nil}
            if expr.token == "-" {return a - b, nil}
            if expr.token == "*" {return a * b, nil}
            if expr.token == "/" {return a / b, nil}

            return 0, errors.New("Неизвестная операция!")

        case 1:
            a, err := eval(expr.args[0])
            if err != nil {
                return 0, err
            }

            if expr.token == "+" {return +a, nil}
            if expr.token == "-" {return -a, nil}

            return 0, errors.New("Неизвестная операция!")
        case 0:
            res, err := strconv.ParseFloat(expr.token, 64)
            if err != nil {
                return 0, err
            }

            return res, nil
    }

    return 0.0, errors.New("Ошибка")
}


func readExpressionFlags() (string, error) {
    flag.Parse()
    switch len(flag.Args()) {
        case 1:
            return flag.Args()[0], nil

        default:
            printHelp()
            return "", errors.New("Переданы некорректные аргументы!")
    }

    return "", errors.New("Переданы некорректные аргументы!")
}


func checkExpression(str string) error {
    if (str == "") {
        return errors.New("Вы ввели пустую строку!")
    }

    leftCount := 0
    rightCount := 0

    tokens := " 0123456789+-*/()"

    for _, symb := range str {
        if !strings.ContainsRune(tokens, symb) {
            return errors.New("Введен недопстимый символ!")
        }

        if symb == '(' {
            leftCount++
        }

        if symb == ')' {
            rightCount++
        }
    }

    if leftCount != rightCount {
        return errors.New("Несовпадающее количество открывающих и закрывающих скобок!")
    }

    return nil
}


func printHelp() {
    fmt.Println("Введены некорректные флаги!\n" +
                 "Пример использования: " +
                 "go run calc.go \"(1+2)*3\"")
}
