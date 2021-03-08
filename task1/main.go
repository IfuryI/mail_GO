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

	strArr = uniq(strArr, opt)

	err = outputStrings(fout, strArr)
	if err != nil {
		fmt.Println(err)
		return
	}
}


func uniq(strArr []string, opt options) []string {
	var result []string
	var repeatsCount int

	var currLine string
	var prevLine string = FSFlagsProccesing(strArr[0], opt)
	var prevLineSource string = strArr[0]

	for i := 1; i < len(strArr); i++ {
		currLine = strArr[i]
		currLine = FSFlagsProccesing(currLine, opt)

		if opt.i && strings.ToLower(prevLine) == strings.ToLower(currLine) {
			repeatsCount++
			continue
		}
		if prevLine == currLine {
			repeatsCount++
			continue
		}

		appendStrToResult(prevLineSource, opt, repeatsCount, &result)

		repeatsCount = 0
		prevLine = currLine
		prevLineSource = strArr[i]
	}

	currLine = FSFlagsProccesing(strArr[len(strArr) - 1], opt)
	appendStrToResult(prevLineSource, opt, repeatsCount, &result)

	return result
}


func FSFlagsProccesing(str string, opt options) string {
	tokens := strings.Split(str, " ")
	if len(tokens) < opt.f {
		return "\n"
	}

	str = strings.Join(tokens[opt.f:], " ")
	
	if len(str) < opt.s {
		return "\n"
	}

	return str[opt.s:]
}


func appendStrToResult(str string, opt options, repeatsCount int, result *[]string) {
	if !opt.c && !opt.d && !opt.u {
		*result = append(*result, str)
		return
	}

	if opt.c {
		*result = append(*result, strconv.Itoa(repeatsCount + 1) + " " + str)
		return
	}

	if (opt.d && repeatsCount != 0) || (opt.u && repeatsCount == 0) {
		*result = append(*result, str)
		return
	}
}


type options struct {
	c bool
	d bool
	u bool
	f int
	s int
	i bool
}


func readFlags() (options, error) {
	CFlag := flag.Bool("c", false, "")
	DFlag := flag.Bool("d", false, "")
	UFlag := flag.Bool("u", false, "")
	FFlag := flag.Int("f",  0, "")
	SFlag := flag.Int("s", 0, "")
	IFlag := flag.Bool("i", false, "")

	flag.Parse()

	opt := options {
		c: *CFlag,
		d: *DFlag,
		u: *UFlag,
		f: *FFlag,
		s: *SFlag,
		i: *IFlag,
	}

	if (opt.c && opt.d && !opt.u || opt.c && !opt.d && opt.u || !opt.c && opt.d && opt.u) {
		return opt, errors.New("Не правильно переданные аргументы командной строки")
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
				 "uniq [-c | -d | -u] [-i] [-f num] [-s chars] [input_file [output_file]]")
}
