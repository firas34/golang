package main

import (
	"fmt"
	"strconv"
)

func strcustom(s string) []string {
	// convert string to Int
	n, _ := strconv.Atoi(s)
	var res []string
	if n < 10 {
		res = append(res, s)
	} else {
		for _, i := range s {
			// convert unicode to string
			// fmt.Println(string(i))
			res = append(res, string(i))
		}
	}
	return res
}
func compress(chars []byte) int {
	_chars := &chars
	var current byte = (*_chars)[0]
	var count int = 1
	res := []string{}
	for i := 1; i < len(*_chars); i++ {
		if (*_chars)[i] == current {
			count++
			fmt.Println(count)
		} else {
			if count == 1 {
				res = append(res, string(current))
			} else {
				res = append(res, string(current))
				// add 2 slices
				res = append(res, strcustom(strconv.Itoa(count))...)
			}
			current = (*_chars)[i+1]
			count = 1
		}
	}
	if count == 1 {
		res = append(res, string(current))
	} else {
		res = append(res, string(current))
		res = append(res, strcustom(strconv.Itoa(count))...)
	}
	//chars = res
	//resB:=[]byte{}
	// st_ := ""
	// for _, i := range res {
	// 	st_ = st_ + i
	// }
	var s string = ""
	for _, i := range res {
		s = s + i
	}
	// var b = chars
	a := []byte(s)

	for idx, _ := range a {
		(*_chars)[idx] = a[idx]
	}

	return len(res)
	//return chars
}

func main() {

	chars := []byte{'a', 'b', 'c'}
	fmt.Println(chars)
	fmt.Println(compress(chars))
	fmt.Println(chars)
}
