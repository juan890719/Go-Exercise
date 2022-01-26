package main

import "fmt"

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

	fmt.Println()
	fmt.Println(beChain("水"))

	sum, diff := calc(6, 4)
	fmt.Printf("%d\n%d", sum, diff)
	fmt.Println()

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
		fmt.Printf("jelly[%d] = %s", k, v)
		fmt.Println()
	}

	const length = 3
	hey := [length]string{"Juan", "Juan", "Huang"}
	for _, v := range hey {
		fmt.Println(v)
	}
}
