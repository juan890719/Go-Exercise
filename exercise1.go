package main

import "fmt"

func double() {
	n *= 2
}

func pluseOneYear(p *person) {
	p.age = p.age + 1 // 等同於(*p).age = (*p).age + 1
}

func arrayPointerAdd1(array *[5]int) {
	for i := 0; i < len(array); i++ {
		(*array)[i] = (*array)[i] + 1
	}
}

func arraySort(array []int) {
	for i := 0; i < len(array); i++ {
		if i == len(array)-1 {
			break
		}
		for j := i + 1; j < len(array); j++ {
			if array[i] > array[j] {
				array[i], array[j] = array[j], array[i]
			} else {
				continue
			}
		}
	}
	for key, value := range array {
		fmt.Printf("array[%d] = %d\n", key, value)
	}
}

func beChain(something string) string {
	return something + "啊, 你快變成懲戒的鎖鏈"
}

func calculate(num1 int, num2 int) (sum int, diff int) {
	sum, diff = num1+num2, num1-num2
	return
}
