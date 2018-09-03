package main

import "fmt"

func main() {
	fmt.Println(reverse(1234))
}

func reverse(x int) int {
	out := 0
	num := x

	if x < 0 {
		num = num * -1
	}

	for num > 0 {
		out = out * 10
		out = out + num%10
		num = num / 10
	}

	if out > 2147483647 || out < -2147483648 {
		return 0
	}

	if x > 0 {
		return out
	}

	if x < 0 {
		return -out
	}

	return 0
}
