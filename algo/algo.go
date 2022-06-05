package algo

func Echo(num int) int {
	return num
}

func Fib(num int) int {
	if num < 2 {
		return 1
	}

	return Fib(num-1) + Fib(num-2)
}

func Sqrt(num int) int {
	if num < 0 {
		return 0
	}
	i, j := 0, 0
	for ; i*i <= num; i++ {
		j = i
	}
	return j
}
