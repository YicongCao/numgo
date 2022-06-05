package algo

import (
	"fmt"
	"testing"
)

var (
	cases = []int{-1, 0, 1, 2, 3, 4, 5}
)

func TestEcho(t *testing.T) {
	for c := range cases {
		fmt.Println(c, Echo(c))
	}
}

func TestFib(t *testing.T) {
	for c := range cases {
		fmt.Println(c, Fib(c))
	}
}

func TestSqrt(t *testing.T) {
	for c := range cases {
		fmt.Println(c, Sqrt(c))
	}
}
