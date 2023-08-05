package main

import (
	"fmt"
	"strconv"
)

func strcustom(s string) []string{
	// convert string to Int
	n,_ := strconv.Atoi(s)
	fmt.Println(n)
	var res []string
	if n < 10{
		res = append(res, s) 
	} else {
		for _,i:= range(s){
			// convert unicode to string
			fmt.Println(string(i))
			res = append(res, string(i))
		}
	}
	return res
}
func compress(chars []byte) []string {
	var current byte = chars[0]
	var count int = 1
	res := []string{}
	for i:=1; i< len(chars); i++ {
		fmt.Println(string(chars[i]))
		if chars[i] == current {
			count++
			fmt.Println("YES - ",string(current), count)
		} else{ 
			if count == 1{
				fmt.Println("CORRECT")
				res = append (res, string(current))
			}else{
				res = append (res, string(current))
				fmt.Println("current: ", string(current), "count: ", strconv.Itoa(count))
				res = append (res, strconv.Itoa(count))
			}
			current = chars[i+1]
			count = 1
		}
	}
	if count == 1{
		res = append(res, string(current))
	}else{
		res = append(res, string(current))
		res = append(res, strconv.Itoa(count))
	}
	return res
}

func main() {
	
	// chars := []byte{'a','b','b','b','b','b','b','b','b','b','b','b','b'}
	//fmt.Println(compress(chars))
	fmt.Println(strcustom("10"))
	// s,err:=strconv.Atoi("10")
	// if err!=nil{
	// 	fmt.Println("wrong input")
	// }else{
	// 	fmt.Println(s+1)
	// }
	

}
