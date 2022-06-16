package main

import "fmt"

func main0201() {
	arr := [9]int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	//s := arr[1:3:5]
	//
	//fmt.Println("s = ", s)
	//fmt.Println("len(s) = ", len(s))
	//fmt.Println("cap(s) = ", cap(s))

	s1 := arr[1:5:7]
	fmt.Println("s1 = ", s1)
	fmt.Println("len(s1) = ", len(s1)) //5-1 == 4
	fmt.Println("cap(s1) = ", cap(s1)) //7-1 == 6
	fmt.Println("==========================")
	s2 := s1[1:3]
	fmt.Println("s2 = ", s2)
	fmt.Println("len(s2) = ", len(s2)) //5-1 == 4
	fmt.Println("cap(s2) = ", cap(s2)) //7-1 == 6

}
func main0202() {
	arr := [9]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s1 := arr[2:5]
	fmt.Println("s1 = ", s1)
	fmt.Println("len(s1) = ", len(s1)) //5-2 == 3
	fmt.Println("cap(s1) = ", cap(s1)) //9-2 == 7
	fmt.Println("==========================")
	s2 := arr[2:7]
	fmt.Println("s2 = ", s2)
	fmt.Println("len(s2) = ", len(s2)) //7-2 == 5
	fmt.Println("cap(s2) = ", cap(s2)) //9-2 == 7
}
func main0203() {
	s1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("s1 = ", s1, "len(s1) = ", len(s1), "cap(s1) = ", cap(s1))

	s2 := make([]int, 3, 5)
	fmt.Println("s2 = ", s2, "len(s2) = ", len(s2), "cap(s2) = ", cap(s2))

	s3 := make([]int, 3)
	fmt.Println("s3 = ", s3, "len(s3) = ", len(s3), "cap(s3) = ", cap(s3))
}
func main0204() {
	s1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	s1 = append(s1, 10)
	s1 = append(s1, 11)
	s1 = append(s1, 12)
	s1 = append(s1, 13)
	s1 = append(s1, 14)

	fmt.Println("s1 = ", s1, ", len(s1) = ", len(s1), ", cap(s1) = ", cap(s1))
}
