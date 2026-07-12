package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println(Hello())
	fmt.Println(Ping())

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello World Go K8s Demo")
	})
	fmt.Println("服务启动，监听8080端口")
	http.ListenAndServe(":8080", nil)
}

func Hello() string {
	return "Hello World CI Demo"
}

func Ping() string {
	return "Ping..."
}
