package myculc

import (
	"fmt"
	"testing"
)

func TestPlus(t *testing.T) {
	expect := 30
	result := plus(10, 20)
	if result != expect {
		t.Errorf("Plus(10, 20) = \"%v\", want \"%v\"", result, expect)
	}
}

func TestSubstract(t *testing.T) {
	s := substract(20, 10)
	if s != 10 {
		t.Error("Eroor")
	}
}

type testData struct {
	n1   int
	n2   int
	want int
}

func TestAdd(t *testing.T) {
	tests := []testData{
		{n1: 10, n2: 20, want: 30},
		{n1: 28, n2: 2, want: 30},
		{n1: 15, n2: 15, want: 30},
		// miss
		{n1: 21, n2: 10, want: 30},
	}

	for _, test := range tests {
		res := plus(test.n1, test.n2)
		if res != test.want {
			t.Errorf(errorString(res, test.want))
		}
	}
}

func errorString(res, want int) string {
	return fmt.Sprintf("%v を期待しましたが、 %v が出力されました。", want, res)
}
