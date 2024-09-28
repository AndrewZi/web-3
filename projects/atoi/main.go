package main

import "fmt"

func main() {
	var st string
	fmt.Scan(&st)
	rs := []rune(st)
	for _, el := range rs {
		fmt.Print((el - 48) * (el - 48))
	}
}
