package main

import "fmt"

func func1() {
	var (
		x int    = -3
		c string = "Runoob"
	)
	y := 4
	if y > 0 {
		fmt.Println("Hello, " + "World!")
	} else {
		fmt.Printf("%v+%q", x, c)
	}
}

func func2() {
	name := "Tony Yip"
	num := 1004
	fmt.Print(name, "'s brithday is", num)
	fmt.Printf("1\n")
	fmt.Printf("%s 's brithday is %d", name, num)
	fmt.Printf("1\n")
	fmt.Println(name, "'s brithday is", num)
	fmt.Printf("1\n")
}

func main() {
	func2()
}
