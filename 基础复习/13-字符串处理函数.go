package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "I love my work and I love my family too"
	s := strings.Split(str, " ")
	fmt.Println(s)
}
