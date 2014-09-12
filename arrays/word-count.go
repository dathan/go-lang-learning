package main

import (
	"code.google.com/p/go-tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	fields := strings.Fields(s)

	m := map[string]int{}

	for _, s := range fields {
		_, ok := m[s]
		if ok == true {
			m[s] += 1
		} else {
			m[s] = 1
		}
	}
	return m
}

func main() {
	wc.Test(WordCount)
}
