package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "I love my work and I love my family too"
	s := strings.Split(str, " ")
	for _, v := range s {
		fmt.Println(v)
	}
}
