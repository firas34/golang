package main

import (
	"fmt"
	"net/http"
)

func printer(c chan int) {

	for i := 0; i < 5; i++ {
		c <- i
	}
	close(c) // close the sender channnel (all values that are called after the end of the list, are == 0)
}

func main() {

	ch := make(chan int)

	go printer(ch)

	// time.Sleep(2 * time.Second)
	for i := 0; i < 7; i++ {
		fmt.Println(<-ch)
	}

	res, err := http.Get("https://www.google.com")
	if err != nil {
		fmt.Println("Oops")
	} else {
		fmt.Printf("%d: status code for google\n", res.StatusCode)
	}

}
