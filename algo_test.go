package numgo_test

import (
	"fmt"
	"numgo"
	"testing"
)

var (
	cases = []int{-1, 0, 1, 2, 3, 4, 5}
)

func TestEcho(t *testing.T) {
	for c := range cases {
		fmt.Println(c, numgo.Echo(c))
	}
}

func TestFib(t *testing.T) {
	for c := range cases {
		fmt.Println(c, numgo.Fib(c))
	}
}

func TestSqrt(t *testing.T) {
	for c := range cases {
		fmt.Println(c, numgo.Sqrt(c))
	}
}
