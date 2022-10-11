package main

import (
	"testing"
)

type TestSFStruct struct {
	input    [2]string
	expected bool
}

var (
	testsSF = []TestSFStruct{
		{[2]string{"down", "downloads"}, true},
		{[2]string{"downloads", "downloads"}, true},
		{[2]string{"don", "downloads"}, false},
		{[2]string{"hi", "downloads"}, false},
		{[2]string{"hist", "history"}, true},
		{[2]string{"d", "history"}, false},
		{[2]string{"h", "history"}, true},
		{[2]string{"d", "downloads"}, true},
	}
)

func TestIsShortForm(t *testing.T) {
	for _, v := range testsSF {
		if IsShortForm(v.input[0], v.input[1]) != v.expected {
			var r string
			if v.expected {
				r = "should be"
			} else {
				r = "shouldn't be"
			}
			t.Errorf("%s %s Short Form of %s, but the value returned suggests other wise", v.input[0], r, v.input[1])
		}
	}
}
