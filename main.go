package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	var name string = "Ovchie"

	fmt.Println(Greetings(name))
	fmt.Println(MinInt(3623763243, 23425235))

	fmt.Println(
		IsValid(0, "hello world"),
		IsValid(-22, "hello world"),
		IsValid(22, ""),
		IsValid(225, "hello world"),
	)
}

func multiply(x int, y int) int {
	return x * y
}

func MinInt(x, y int) int {
	return int(math.Min(float64(x), float64(y)))
}

func IsValid(id int, text string) bool {
	return id > 0 && text != ""
}

func Greetings(name string) string {
	name = strings.Trim(name, " ")
	name = strings.ToLower(name)
	name = strings.Title(name)
	return "Привет, " + name + "!"
}
