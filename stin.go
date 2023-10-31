package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res, _ := reader.ReadString('\n')
	num, err := strconv.ParseFloat(strings.TrimSpace(res), 64) // FOCUS ON TRIM // res = "5\n"
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(num)
}
