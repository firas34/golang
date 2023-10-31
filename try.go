package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup

func urlCall(web string) {
	defer wg.Done() //  defer is used to delay the execution of a function or a statement until the nearby function returns
	res, err := http.Get(web)
	if err != nil {
		fmt.Println("Oops")
	} else {
		fmt.Printf("%d: status code for %s\n", res.StatusCode, web)
	}
	// wg.Done() // or we use it without defer since it's at the end of the function
}

func main() {

	urls := []string{
		"https://google.com",
		"https://fb.com",
		"https://go.dev",
		"https://google.com",
		"https://fb.com",
		"https://go.dev",
	}

	for _, url := range urls {
		go urlCall(url)
		wg.Add(1)
	}

	wg.Wait()
}
