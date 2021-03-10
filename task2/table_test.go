package main

import (
	"strconv"
	"testing"
	"fmt"
)


func TestTablePositive(t *testing.T) {
	testDataArr := initDataPositive()

	for i, test := range testDataArr {
		parser := parser{test.input}

		parsedExpression, err := parser.mainParse()
	    if err != nil {
	        fmt.Println(err)
	        return
	    }

		res, err := eval(parsedExpression)
		if err != nil || test.output != res {
			fmt.Println(err)
			t.Fatalf("Failed: " + strconv.Itoa(i) + " " + test.input + " " + strconv.FormatFloat(test.output, 'f', 1, 64) + " " + strconv.FormatFloat(res, 'f', 1, 64))
		}
	}
}


func TestTableNegative(t *testing.T) {
	testDataArr := initDataNegative	()

	for i, test := range testDataArr {
		parser := parser{test.input}

		parsedExpression, err1 := parser.mainParse()

		res, err2 := eval(parsedExpression)
		if err2 == nil || err1 == nil || res == test.output {
			t.Fatalf("Failed2: " + strconv.Itoa(i) + " " + test.input)
		}
	}
}


type testData struct {
	input string
	output float64
}


func initDataPositive() []testData {
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
			"(3*2) + 10",
			16,
		},
		testData {
			"(10+3) * 2",
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


func initDataNegative() []testData {
	return []testData {
		testData {
			"qew",
			4,
		},
		testData {
			"",
			-4,
		},
		testData {
			"@2",
			2,
		},
		testData {
			"q2",
			2,
		},
		testData {
			"++w++4",
			4,
		},
		testData {
			" (2 + 2",
			4,
		},
		testData {
			"(9 + ( 9 + ( 2 + 3 ",
			18,
		},
		testData {
			"()",
			2.0,
		},
		testData {
			"qw       1231231",
			16,
		},
		testData {
			"( +",
			26,
		},
		testData {
			"()(2)",
			13,
		},
		testData {
			"(2**2)",
			39,
		},
		testData {
			" (2 + 2 + (22 * 3) ",
			43,
		},
	}
}
