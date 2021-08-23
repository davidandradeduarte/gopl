package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}
	for i, u := range os.Args[1:] {
		wd, err := os.Getwd()
		if err != nil {
			panic(err)
		}
		f, err := os.Create(fmt.Sprintf("%s/exercises/1.10.%d.%s.html", wd, i, strings.Replace(u, "/", "", -1)))
		if err != nil {
			panic(err)
		}

		_, err = f.Write([]byte(<-ch))
		if err != nil {
			panic(err)
		}

		err = f.Close()
		if err != nil {
			panic(err)
		}
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	b, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	fmt.Printf("%.2fs  %7d  %s", secs, len(b), url)
	ch <- string(b)
}
