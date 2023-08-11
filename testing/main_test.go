package main

import "testing"

/*
Things you can do:
1) Check your coverage with this command:
    go test -cover

2) Get your coverage in the browser with this command:
    go test -coverprofile=coverage.out && go tool cover -html=coverage.out
*/

// 1) Manually Way
func TestDivide(t *testing.T) {
	_, err := divider(100.0, 10.0)
	if err != nil {
		t.Error("Got an error when is should not have")
	}
}

func TestBadDivide(t *testing.T) {
	_, err := divider(100.0, 0)
	if err == nil {
		t.Error("Got an error when is should have")
	}
}

// 2) with a function and struct
var divideTest []struct {
	name     string
	dividend float32
	divisor  float32
	expected float32
	isError  bool
}{
	{"valid-data", 100.0, 10.0, 10.0, false},
	{"invalid-data", 100.0, 0.0, 0.0, true},
	{"expect-5", 50.0, 10.0, 5.0, false},
	{"expect-fraction", -1.0, -777.0, 0.0012870013, false},
}

func TestDivision(t *testing.T) {
	for _, tt := range tests {
		got, err := divider(t.divident, t.divisor)
		if tt.isErr{

			if err == nil {
				t.Error("Expected an error but not have one")
			}
		} else {
			if err != nil {
				t.Error("Did not expect and error but got one")
			}
		} else {
			t.Error("Expected %f but got %f", tt.expected, got)
		}
	}
}