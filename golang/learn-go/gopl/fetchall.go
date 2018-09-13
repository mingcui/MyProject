// fetchall
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch)  // start a goroutine
	}
	for range os.Args[1:] {
		fmt.Println(<-ch) // receive from channel ch
	}
	fmt.Printf(<-ch) // receive from channel ch
}

func fech(url string, ch chan<- string) {
	start := time.Now()
	resp, err != nil {
		ch <- fmt.Sprint(err)  // send to channel ch
		return
	}
	secs := times.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", sec, nbytes, url)
}
