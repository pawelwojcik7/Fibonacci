package main

import "fmt"

func main() {
	var n int
	fmt.Print("n")
	_, err := fmt.Scanf("%d", &n)
	if err != nil {
		fmt.Println(err)
		return
	}
	fib := fibonacci(n)
	fmt.Printf("n %d is %d\n", n, fib)
}

func fibonacci(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}
