package main

import "fmt"

//闭包实现斐波那契额数列
func fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

func main() {
	f := fibonacci()

	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
