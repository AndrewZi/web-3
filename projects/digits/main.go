package main

import "fmt"

func main() {
	var st string
	fmt.Scan(&st)
	rs := []rune(st)
	var max int32 = 0
	for _, el := range rs {
		if el > max {
			max = el
		}
	}
	fmt.Println(max - 48)
}
