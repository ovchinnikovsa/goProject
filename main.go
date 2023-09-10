package main

import (
	"fmt"
	"math"
	"strings"
)

type UserCreateRequest struct {
	FirstName string
	Age       int
}

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

	fmt.Println(
		DomainForLocale("google.com", "ru"),
		DomainForLocale("google.com", "en"),
	)

	fmt.Println(
		ModifySpaces("hello world", "underscore"),
		ModifySpaces("hello world", "dash"),
		ModifySpaces("hello world", "unknown"),
		ModifySpaces("hello world", ""),
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

func DomainForLocale(domain, locale string) string {
	if locale == "" {
		return "en" + "." + domain
	}
	return locale + "." + domain
}

func ModifySpaces(s, mode string) string {
	switch mode {
	case "dash":
		return strings.ReplaceAll(s, " ", "-")
	case "underscore":
		return strings.ReplaceAll(s, " ", "_")
	case "unknown":
		fallthrough
	default:
		return strings.ReplaceAll(s, " ", "*")
	}
}

func Validate(req UserCreateRequest) string {
	error_message := "invalid request"
	switch {
	case req.Age > 150:
		return error_message
	case req.Age <= 0:
		return error_message
	case req.FirstName == "":
		return error_message
	case strings.Contains(req.FirstName, " "):
		return error_message
	default:
		return ""
	}
}

func ErrorMessageToCode(msg string) int {
	switch msg {
	case "OK":
		return 0
	case "CANCELLED":
		return 1
	case "UNKNOWN":
		fallthrough
	default:
		return 2
	}
}
