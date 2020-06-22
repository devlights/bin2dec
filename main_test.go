package main

import (
	"testing"
)

func TestConvert(t *testing.T) {
	type (
		testin struct {
			val string
		}
		testout struct {
			result string
		}
		testcase struct {
			name string
			in   testin
			out  testout
		}
	)

	cases := []testcase{
		{
			name: "convert 0b11111111",
			in:   testin{val: "0b11111111"},
			out:  testout{result: "255"},
		},
		{
			name: "convert 0b1000",
			in:   testin{val: "0b1000"},
			out:  testout{result: "8"},
		},
		{
			name: "convert 1000",
			in:   testin{val: "1000"},
			out:  testout{result: "8"},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			b, err := Convert(c.in.val)
			if err != nil {
				t.Error(err)
			}

			if c.out.result != b {
				t.Errorf("[want] %v\t[got] %v\n", c.out.result, b)
			}
		})
	}
}
