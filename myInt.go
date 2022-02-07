package main

import "fmt"

type myInt int

func (num *myInt) absoulte() myInt {
	if *num < 0 {
		*num = -(*num)
	}
	return *num
}

func printAbsoulteNum() {
	var number myInt = -3
	pointer := &number
	fmt.Println(pointer.absoulte())
	fmt.Println(number)
}
