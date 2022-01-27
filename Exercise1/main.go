package main

import "fmt"

type person struct {
	name string
	age  uint
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
				v1 := array[i]
				array[i] = array[j]
				array[j] = v1
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
	return something + "啊，你快變成懲戒的鎖鏈"
}

func calc(num1 int, num2 int) (sum int, diff int) {
	sum, diff = num1+num2, num1-num2
	return
}

func main() {
	num1 := 1.25 // num1 is float
	num2 := 2    // num2 is int
	fmt.Println(num1 * float64(num2))

	var com1 complex64 = 5 + 3i
	var com2 complex64 = 1.3 + 2i
	fmt.Println(com1 + com2)

	// hello()

	var a rune = '哈'
	var b rune = '囉'
	fmt.Printf("%c%c", a, b)

	var s string = "絹絹"
	fmt.Println(s)

	variant := fmt.Sprintf("%v %v %v", "月球基地", 3.14159, true)
	fmt.Println(variant)

	num := 3
	if num < 5 {
		fmt.Println("The number is smail.")
	}

	num3 := -3
	if num3 > 0 {
		fmt.Println("+")
	} else {
		fmt.Println("-")
	}

	for i := 1; i < 11; i++ {
		fmt.Printf("%d ", i)
	}

	fmt.Println("\n" + beChain("水"))

	sum, diff := calc(6, 4)
	fmt.Printf("%d\n%d\n", sum, diff)

	// Array
	// var food [6]string
	// food[0] = "Juan"
	// food[1] = "Wei"
	// food[2] = "Dad eat 2"
	// food[3] = "Mom"
	// food[4] = "Jack"
	// food[5] = "Amy"
	var food = [6]string{
		"Amy",
		"Jack",
		"Fen",
		"Juan",
		"Wei",
		"Kate",
	}
	fmt.Println(food[0])

	jelly := [...]string{"黃絹", "吳承蓉", "吳敏綸"}
	for k, v := range jelly {
		fmt.Printf("jelly[%d] = %s\n", k, v)
	}

	const length = 3
	hey := [length]string{"Juan", "Juan", "Huang"}
	for _, v := range hey {
		fmt.Println(v)
	}

	// Slice
	slice1 := food[0:3]
	fmt.Printf("slice1's capacity: %d\n", cap(slice1))
	fmt.Printf("slice's length: %d\n", len(slice1))

	slice2 := make([]string, 3, 5)
	slice2[0] = "hey"
	slice2[1] = "hello"
	slice2[2] = "mello"

	slice3 := []string{}
	slice3 = append(slice3, "oh")
	slice3 = append(slice3, "my")
	slice3 = append(slice3, "god")

	// Map
	tall1 := make(map[string]int)
	tall1["Juan"] = 155
	tall1["Jong"] = 170
	tall1["Minlun"] = 154

	tall2 := map[string]int{
		"Juan":   155,
		"Jong":   170,
		"Minlun": 154,
	}
	// 第一個值val是應的value，第二個值ok是表示此key->value是否存在的布林值
	val, ok := tall2["Jong"]
	fmt.Println(val, ok)
	val, ok = tall2["Jack"]
	fmt.Println(val, ok)

	for k, v := range tall2 {
		fmt.Printf("%s的身高是%d\n", k, v)
	}

	delete(tall2, "Juan")
	for k := range tall2 {
		fmt.Println(k)
	}

	// exercise1 利用迴圈印出九九乘法表
	for i := 2; i < 10; i++ {
		for j := 2; j < 10; j++ {
			fmt.Printf("%d X %d = %d\t", j, i, j*i)
		}
		fmt.Println()
	}

	// exercise2 由小到到大排序一個整數陣列
	exerciseArray := [...]int{3, 1, 2, 5, 1}
	arraySort(exerciseArray[:])

	// pointer
	var p *int
	p = &num
	fmt.Println(p, " : ", *p)

	// array不是指標
	array1 := [5]int{1, 3, 4, 2, 5}
	arrayPointerAdd1(&array1)
	for _, v := range array1 {
		fmt.Print(v, " ")
	}
	fmt.Println()

	// type struct 定義型態別名
	var sakura person
	sakura.name = "小櫻"
	sakura.age = 10
	fmt.Println(sakura)
	fmt.Printf("%T\n", sakura)

	juan := person{
		name: "Huang Juan",
		age:  21,
	}
	fmt.Println(juan)

	// 不使用type, 直接用struct宣告
	jong := struct {
		name string
		age  uint
	}{
		name: "Jong",
		age:  21,
	}
	fmt.Println(jong)
	pluseOneYear(&sakura)
	fmt.Println(sakura)

	// 用new()宣告一個struct的指標變數
	minlun := new(person)
	minlun.name = "Minlun"
	minlun.age = 22
	fmt.Println(minlun)  // &{Minlun 22}
	fmt.Println(&minlun) // 指標：0x1400000e030
	fmt.Println(*minlun) // {Minlun 22}
}
