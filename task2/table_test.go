package main

import (
	"strconv"
	"testing"	
)


func TestTable(t *testing.T) {
	testDataArr := initData()

	for i, test := range testDataArr {

		prs := parser{test.input}

		if test.output != eval(prs.mainParse()) {
			t.Fatalf("Failed: " + strconv.Itoa(i) + " " + test.input)
		}
	}
}


type testData struct {
	input string
	output float64
}

func initData() []testData {
	return []testData {
		testData {
			"4",
			4,
		},
		testData {
			" -4",
			-4,
		},
		testData {
			"0",
			0,
		},
		testData {
			"+4",
			+4,
		},
		testData {
			"++++4",
			4,
		},
		testData {
			"    --4",
			4,
		},
		testData {
			"9+9",
			18,
		},
		testData {
			"9+9",
			18,
		},
		testData {
			"10+3*2",
			16,
		},
		testData {
			"(10+3)*2",
			26,
		},
		testData {
			"(10+3)*2/2",
			13,
		},
		testData {
			"(10+3)*2/2*3",
			39,
		},
		testData {
			"(10-5)*(3+5)+2+1",
			43,
		},
	}
}
