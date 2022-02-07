package main

import (
	"fmt"
	"sort"
)

type myList []int

func (list myList) Len() int {
	return len(list)
}

func (list myList) Less(i, j int) bool {
	return list[i] < list[j]
	// 希望前面的放比較小（降冪）
}

func (list myList) Swap(i, j int) {
	list[i], list[j] = list[j], list[i]
	// 傳統寫法：
	// tmp := list[i]
	// list[i] = list[j]
	// list[j] = tmp
}

func printSortMyList() {
	list := []int{1, 4, 8, 3, 5, 7, 9, 6}
	newList := myList(list) // 轉型
	fmt.Println(newList)
	sort.Sort(newList)
	fmt.Println(newList)
}
