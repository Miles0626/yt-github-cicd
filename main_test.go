package main

import "testing"

func TestHello(t *testing.T) {
	expected := "Hello World CI Demo"
	result := Hello()
	if expected != result {
		t.Fatalf("输出不匹配，期待%s, 得到 %s", expected, result)
	}
}
