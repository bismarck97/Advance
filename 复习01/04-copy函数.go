package main

import "fmt"

func main04() {
	data := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s1 := data[8:] //[8,9]
	fmt.Println("s1=", s1)
	s2 := data[0:5] //[0,1,2,3,4]
	fmt.Println("s2=", s2)
	copy(s2, s1)
	fmt.Println("拷贝后：s2=", s2)
}
