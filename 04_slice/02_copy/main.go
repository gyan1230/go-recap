package main

import "fmt"

func main() {
	//initilization of slices
	var nilslice []int   //nil slice
	empty := []int{}     //empty slice
	empty2 := []string{} //empty slice

	fmt.Println("length and cap of nil slice is", len(nilslice), cap(nilslice))
	fmt.Println("length and cap of emp slice is", len(empty), cap(empty))
	fmt.Printf("Addres of nil slice is %p\n", nilslice)
	fmt.Printf("Addres of emp slice is %p\n", empty)
	fmt.Printf("Addres of emp2 slice is %p\n", empty2)
	empty = append(empty, 1)
	empty2 = append(empty2, "1")
	fmt.Printf("Addres of emp slice is %p\n", empty)
	fmt.Printf("Addres of emp2 slice is %p\n", empty2)

	var s1 = []int{1, 2, 3, 4, 5}

	s2 := []int{
		11,
		22,
		33,
	}

	s3 := make([]int, 3, 5)
	N := copy(s3, s2)
	fmt.Println(s1, s2, s3, N)

}
