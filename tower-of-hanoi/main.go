package main

import "fmt"

func move(n int, A string, B string, C string) {
	if n == 1 {
		fmt.Println(A, "-->", C)
	} else {
		move(n-1, A, C, B)
		fmt.Println(A, "-->", C)
		move(n-1, B, A, C)
	}
}

func main() {
	move(3, "A", "B", "C")
}
