package main

import (
	"fmt"
	"flag"
	"errors"
	"bufio"
	"os"
	"strconv"
	"strings"
)


func main() {
	fin := os.Stdin
	defer fin.Close()

	fout := os.Stdout
	defer fout.Close()

	opt, err := readFlags()
	if err != nil {
		fmt.Println(err)
		printHelp()
		return
	}

	err = readIOFileFlags(&fin, &fout)
	if err != nil {
		fmt.Println(err)
		return
	}

	strArr, err := inputStrings(fin)
	if err != nil || len(strArr) == 0 {
		fmt.Println(err)
		return
	}

	strArr, err = uniq(strArr, opt)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = outputStrings(fout, strArr)
	if err != nil {
		fmt.Println(err)
		return
	}
}


func uniq(strArr []string, opt options) ([]string, error) {
	err := checkFlags(opt)
	if err != nil {
		return []string{}, err
	}

	var result []string
	var repeatsCount int

	var currLine string
	var prevLine string = FSFlagsProccesing(strArr[0], opt)
	var prevLineSource string = strArr[0]

	for ignoreCase := 1; ignoreCase < len(strArr); ignoreCase++ {
		currLine = strArr[ignoreCase]
		currLine = FSFlagsProccesing(currLine, opt)

		if opt.ignoreCase && (strings.ToLower(prevLine) == strings.ToLower(currLine)) ||
		   (prevLine == currLine) {
			repeatsCount++
			continue
		}

		appendStrToResult(prevLineSource, opt, repeatsCount, &result)

		repeatsCount = 0
		prevLine = currLine
		prevLineSource = strArr[ignoreCase]
	}

	currLine = FSFlagsProccesing(strArr[len(strArr) - 1], opt)
	appendStrToResult(prevLineSource, opt, repeatsCount, &result)

	return result, nil
}


func FSFlagsProccesing(str string, opt options) string {
	tokens := strings.Split(str, " ")
	if len(tokens) < opt.fields {
		return "\n"
	}

	str = strings.Join(tokens[opt.fields:], " ")

	if len(str) < opt.chars {
		return "\n"
	}

	return str[opt.chars:]
}


func appendStrToResult(str string, opt options, repeatsCount int, result *[]string) {
	if (!opt.count && !opt.duplicate && !opt.uniq) ||
	   (opt.duplicate && repeatsCount != 0) || (opt.uniq && repeatsCount == 0) {
		*result = append(*result, str)
		return
	}

	if opt.count {
		*result = append(*result, strconv.Itoa(repeatsCount + 1) + " " + str)
		return
	}
}


type options struct {
	count bool
	duplicate bool
	uniq bool
	fields int
	chars int
	ignoreCase bool
}

func checkFlags(opt options) error {
	if (opt.fields < 0 || opt.chars < 0) {
		return errors.New("Аргумент флагов fields и chars не может быть меньше 0!")
	}

	if (opt.count && opt.duplicate && !opt.uniq ||
		opt.count && !opt.duplicate && opt.uniq ||
		!opt.count && opt.duplicate && opt.uniq ||
		opt.count && opt.duplicate && opt.uniq) {
		return errors.New("Не правильно переданные аргументы командной строки")
	}

	return nil
}

func readFlags() (options, error) {
	CFlag := flag.Bool("c", false, "Подсчитать количество встречаний строки во входных данных")
	DFlag := flag.Bool("d", false, "Вывести только те строки, которые повторились во входных данных")
	UFlag := flag.Bool("u", false, "Вывести только те строки, которые не повторились во входных данных")
	FFlag := flag.Int("f",  0, "Не учитывать первые N полей в строке")
	SFlag := flag.Int("s", 0, "Не учитывать первые N символов в строке")
	IFlag := flag.Bool("i", false, "Не учитывать регистр букв")

	flag.Parse()

	opt := options {
		count: *CFlag,
		duplicate: *DFlag,
		uniq: *UFlag,
		fields: *FFlag,
		chars: *SFlag,
		ignoreCase: *IFlag,
	}

	err := checkFlags(opt)
	if err != nil {
		return opt, err
	}

	return opt, nil
}


func readIOFileFlags(fin **os.File, fout **os.File) error {
	var err error

	switch len(flag.Args()) {
		case 0:
			return nil

		case 1:
			*fin, err = os.Open(flag.Args()[0])
			if err != nil {
				return err
			}

		case 2:
			*fin, err = os.Open(flag.Args()[0])
			if err != nil {
				return err
			}

			*fout, err = os.Create(flag.Args()[1])
			if err != nil {
				return err
			}

		default:
			printHelp();
			return errors.New("Введены некорректные флаги!")
	}

	return nil
}


func inputStrings(fin *os.File) (strings []string, err error) {
	strReader := bufio.NewReader(fin)

	for {
		dataString, err := strReader.ReadString('\n')
		if err != nil && len(dataString) == 0 {
			break
		}

		if err != nil {
			err = errors.New("File reading error")
			return nil, err
		}

		strings = append(strings, dataString[:len(dataString) - 1])
	}

	return strings, nil
}


func outputStrings(fout *os.File, strArr []string) error {
	for _, str := range strArr {
		_, err := fout.WriteString(str + "\n")
		if err != nil {
			return err
		}
	}

	return nil
}


func printHelp() {
	fmt.Println("Введены некорректные флаги!\n" +
				 "Использовать так: " +
				 "uniq [-c | -d | -u] [-i] [-f num] [-c chars] [input_file [output_file]]")
}
