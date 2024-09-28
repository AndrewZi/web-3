package main

import (
	"fmt"
	"strings"
)

func main() {
	var st string
	fmt.Scan(&st)
	fmt.Println(strings.Join(strings.Split(st, ""), "*"))
}
