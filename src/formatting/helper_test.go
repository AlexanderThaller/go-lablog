package formatting

import (
	"testing"

	testhelper "github.com/AlexanderThaller/lablog/src/testing"
)

func Test_HeaderIndent(t *testing.T) {
	tests := map[int]string{
		-1: "",
		0:  "",
		1:  "=",
		2:  "==",
		3:  "===",
		4:  "====",
		5:  "=====",
		6:  "======",
	}

	for input, expected := range tests {
		got := HeaderIndent("=", input)
		testhelper.CompareGotExpected(t, nil, got, expected)
	}
}
