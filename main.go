package main

import "fmt"

func main() {
	fmt.Println(Hello())
	fmt.Println(Ping())
}

func Hello() string {
	return "Hello World CI Demo"
}

func Ping() string {
	return "Ping..."
}
