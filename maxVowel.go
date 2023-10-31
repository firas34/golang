package main

import (
	"fmt"
)

func isVowel(char byte) bool {
	if char == 'a' || char == 'e' || char == 'o' || char == 'i' || char == 'u' {
		return true
	} else {
		return false
	}
}
func vowelCount(s string) int {
	res := 0
	for _, i := range s {
		if isVowel(byte(i)) {
			res = res + 1
		}
	}
	return res
}
func maxVowels(s string, k int) int {

	initV := vowelCount(s[0:k])
	max := initV

	for i := k; i < len(s); i++ {
		if isVowel(s[i-k]) {
			initV--
		}
		if isVowel(s[i]) {
			initV++
		}
		if max < initV {
			max = initV
		}
	}

	return max
}

func main() {

	s := "leetcode"
	fmt.Println(maxVowels(s, 2))
}
