package main

import (
	"reflect"
	"testing"
	"strconv"
)


func TestTablePositive(t *testing.T) {
	testDataArr := initDataPositive()

	for ignoreCase, test := range testDataArr {
		res, err := uniq(test.input, test.opt)
		if err != nil || !reflect.DeepEqual(test.output, res) {
			t.Fatalf("Failed: " + strconv.Itoa(ignoreCase) + " " + test.testName)
		}
	}
}


func TestTableNegative(t *testing.T) {
	testDataArr := initDataNegative()
	for ignoreCase, test := range testDataArr {
		res, err := uniq(test.input, test.opt)
		if err == nil || !reflect.DeepEqual([]string{}, res) {
			t.Fatalf("Failed: " + strconv.Itoa(ignoreCase) + " " + test.testName)
		}
	}
}


type testData struct {
	testName string
	input []string
	opt options
	output []string
}


func initDataPositive() []testData {
	return []testData {
		testData {
			"Тест без параметров",
			[]string {
				"I love music.",
				"I love music.",
				"I love music.",
				"",
				"I love music of Kartik.",
				"I love music of Kartik.",
				"Thanks.",
				"I love music of Kartik.",
				"I love music of Kartik.",
			},
			options {
				count: false,
				duplicate: false,
				uniq: false,
				fields: 0,
				chars: 0,
				ignoreCase: false,
			},
			[]string {
				"I love music.",
				"",
				"I love music of Kartik.",
				"Thanks.",
				"I love music of Kartik.",
			},
		},
		testData {
			"Тест с параметром -count",
			[]string {
				"I love music.",
				"I love music.",
				"I love music.",
				"",
				"I love music of Kartik.",
				"I love music of Kartik.",
				"Thanks.",
				"I love music of Kartik.",
				"I love music of Kartik.",
			},
			options {
				count: true,
				duplicate: false,
				uniq: false,
				fields: 0,
				chars: 0,
				ignoreCase: false,
			},
			[]string {
				"3 I love music.",
				"1 ",
				"2 I love music of Kartik.",
				"1 Thanks.",
				"2 I love music of Kartik.",
			},
		},
		testData {
			"Тест с параметром -duplicate",
			[]string {
				"I love music.",
				"I love music.",
				"I love music.",
				"",
				"I love music of Kartik.",
				"I love music of Kartik.",
				"Thanks.",
				"I love music of Kartik.",
				"I love music of Kartik.",
			},
			options {
				count: false,
				duplicate: true,
				uniq: false,
				fields: 0,
				chars: 0,
				ignoreCase: false,
			},
			[]string {
				"I love music.",
				"I love music of Kartik.",
				"I love music of Kartik.",
			},
		},
		testData {
			"Тест с параметром -uniq",
			[]string {
				"I love music.",
				"I love music.",
				"I love music.",
				"",
				"I love music of Kartik.",
				"I love music of Kartik.",
				"Thanks.",
				"I love music of Kartik.",
				"I love music of Kartik.",
			},
			options {
				count: false,
				duplicate: false,
				uniq: true,
				fields: 0,
				chars: 0,
				ignoreCase: false,
			},
			[]string {
				"",
				"Thanks.",
			},
		},
		testData {
			"Тест с параметром -fields 1",
			[]string {
				"We love music.",
				"I love music.",
				"They love music.",
				"",
				"I love music of Kartik.",
				"We love music of Kartik.",
				"Thanks.",
			},
			options {
				count: false,
				duplicate: false,
				uniq: false,
				fields: 1,
				chars: 0,
				ignoreCase: false,
			},
			[]string {
				"We love music.",
				"",
				"I love music of Kartik.",
				"Thanks.",
			},
		},
		testData {
			"Тест с параметром -chars 1",
			[]string {
				"I love music.",
				"A love music.",
				"C love music.",
				"",
				"I love music of Kartik.",
				"We love music of Kartik.",
				"Thanks.",
			},
			options {
				count: false,
				duplicate: false,
				uniq: false,
				fields: 0,
				chars: 1,
				ignoreCase: false,
			},
			[]string {
				"I love music.",
				"",
				"I love music of Kartik.",
				"We love music of Kartik.",
				"Thanks.",
			},
		},
		testData {
			"Тест с параметром -ignoreCase",
			[]string {
				"I LOVE MUSIC.",
				"I love music.",
				"I LoVe MuSiC.",
				"",
				"I love MuSIC of Kartik.",
				"I love music of kartik.",
				"Thanks.",
				"I love music of kartik.",
				"I love MuSIC of Kartik.",
			},
			options {
				count: false,
				duplicate: false,
				uniq: false,
				fields: 0,
				chars: 0,
				ignoreCase: true,
			},
			[]string {
				"I LOVE MUSIC.",
				"",
				"I love MuSIC of Kartik.",
				"Thanks.",
				"I love music of kartik.",
			},
		},
	}
}


func initDataNegative() []testData {
	return []testData {
		testData {
			"Тест без параметров",
			[]string {
				"I love music.",
				"I love music.",
				"I love music.",
				"",
				"I love music of Kartik.",
				"I love music of Kartik.",
				"Thanks.",
				"I love music of Kartik.",
				"I love music of Kartik.",
			},
			options {
				count: true,
				duplicate: true,
				uniq: true,
				fields: 0,
				chars: 0,
				ignoreCase: false,
			},
			[]string {
				"I love music.",
				"",
				"I love music of Kartik.",
				"Thanks.",
				"I love music of Kartik.",
			},
		},
		testData {
			"Тест без параметров",
			[]string {
				"I love music.",
				"I love music.",
				"I love music.",
				"",
				"I love music of Kartik.",
				"I love music of Kartik.",
				"Thanks.",
				"I love music of Kartik.",
				"I love music of Kartik.",
			},
			options {
				count: false,
				duplicate: true,
				uniq: true,
				fields: 0,
				chars: 0,
				ignoreCase: false,
			},
			[]string {
				"I love music.",
				"",
				"I love music of Kartik.",
				"Thanks.",
				"I love music of Kartik.",
			},
		},
		testData {
			"Тест с параметром -duplicate",
			[]string {
				"I love music.",
				"I love music.",
				"I love music.",
				"",
				"I love music of Kartik.",
				"I love music of Kartik.",
				"Thanks.",
				"I love music of Kartik.",
				"I love music of Kartik.",
			},
			options {
				count: true,
				duplicate: true,
				uniq: false,
				fields: 0,
				chars: 0,
				ignoreCase: false,
			},
			[]string {
				"I love music.",
				"I love music of Kartik.",
				"I love music of Kartik.",
			},
		},
		testData {
			"Тест без параметров",
			[]string {
				"I love music.",
				"I love music.",
				"I love music.",
				"",
				"I love music of Kartik.",
				"I love music of Kartik.",
				"Thanks.",
				"I love music of Kartik.",
				"I love music of Kartik.",
			},
			options {
				count: true,
				duplicate: false,
				uniq: true,
				fields: 0,
				chars: 0,
				ignoreCase: false,
			},
			[]string {
				"I love music.",
				"",
				"I love music of Kartik.",
				"Thanks.",
				"I love music of Kartik.",
			},
		},
		testData {
			"Тест без параметров",
			[]string {
				"I love music.",
				"I love music.",
				"I love music.",
				"",
				"I love music of Kartik.",
				"I love music of Kartik.",
				"Thanks.",
				"I love music of Kartik.",
				"I love music of Kartik.",
			},
			options {
				count: false,
				duplicate: true,
				uniq: true,
				fields: -3,
				chars: 0,
				ignoreCase: false,
			},
			[]string {
				"I love music.",
				"",
				"I love music of Kartik.",
				"Thanks.",
				"I love music of Kartik.",
			},
		},
		testData {
			"Тест без параметров",
			[]string {
				"I love music.",
				"I love music.",
				"I love music.",
				"",
				"I love music of Kartik.",
				"I love music of Kartik.",
				"Thanks.",
				"I love music of Kartik.",
				"I love music of Kartik.",
			},
			options {
				count: false,
				duplicate: true,
				uniq: true,
				fields: 0,
				chars: -2,
				ignoreCase: false,
			},
			[]string {
				"I love music.",
				"",
				"I love music of Kartik.",
				"Thanks.",
				"I love music of Kartik.",
			},
		},
		testData {
			"Тест без параметров",
			[]string {
				"I love music.",
				"I love music.",
				"I love music.",
				"",
				"I love music of Kartik.",
				"I love music of Kartik.",
				"Thanks.",
				"I love music of Kartik.",
				"I love music of Kartik.",
			},
			options {
				count: false,
				duplicate: true,
				uniq: true,
				fields: -12,
				chars: -12,
				ignoreCase: false,
			},
			[]string {
				"I love music.",
				"",
				"I love music of Kartik.",
				"Thanks.",
				"I love music of Kartik.",
			},
		},
	}
}
