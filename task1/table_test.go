package main

import (
	"reflect"
	"testing"
	"strconv"
)


func TestTable(t *testing.T) {
	testDataArr := initData()

	for i, test := range testDataArr {
		if !reflect.DeepEqual(test.output, uniq(test.input, test.opt)) {
			t.Fatalf("Failed: " + strconv.Itoa(i) + " " + test.testName)
		}
	}
}


type testData struct {
	testName string
	input []string
	opt options
	output []string
}


func initData() []testData {
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
				c: false,
				d: false,
				u: false,
				f: 0,
				s: 0,
				i: false,
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
			"Тест с параметром -c",
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
				c: true,
				d: false,
				u: false,
				f: 0,
				s: 0,
				i: false,
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
			"Тест с параметром -d",
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
				c: false,
				d: true,
				u: false,
				f: 0,
				s: 0,
				i: false,
			},
			[]string {
				"I love music.",
				"I love music of Kartik.",
				"I love music of Kartik.",
			},
		},
		testData {
			"Тест с параметром -u",
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
				c: false,
				d: false,
				u: true,
				f: 0,
				s: 0,
				i: false,
			},
			[]string {
				"",
				"Thanks.",
			},
		},
		testData {
			"Тест с параметром -f 1",
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
				c: false,
				d: false,
				u: false,
				f: 1,
				s: 0,
				i: false,
			},
			[]string {
				"We love music.",
				"",
				"I love music of Kartik.",
				"Thanks.",
			},
		},
		testData {
			"Тест с параметром -s 1",
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
				c: false,
				d: false,
				u: false,
				f: 0,
				s: 1,
				i: false,
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
			"Тест с параметром -i",
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
				c: false,
				d: false,
				u: false,
				f: 0,
				s: 0,
				i: true,
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
