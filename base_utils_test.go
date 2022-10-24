package main

import (
	"fmt"
	"testing"
)

func Test_SortByValue(t *testing.T) {
	p := SortByValue(map[string]int{"一": 0, "而": 9, "yi": 1, "三": 3})
	for k, v := range p {
		fmt.Println(k, v)
	}
}
