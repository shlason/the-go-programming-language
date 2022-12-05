package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"
)

type Memo struct {
	requests chan request
}

type entry struct {
	res   result
	ready chan struct{}
}

func (e *entry) call(f Func, key string) {
	e.res.value, e.res.err = f(key)
	close(e.ready)
}

func (e *entry) deliver(response chan<- result) {
	<-e.ready
	response <- e.res
}

func (memo *Memo) server(f Func) {
	cache := make(map[string]*entry)

	for request := range memo.requests {
		e := cache[request.key]

		if e == nil {
			e = &entry{ready: make(chan struct{})}
			cache[request.key] = e
			go e.call(f, request.key)
		}
		go e.deliver(request.response)
	}
}

func (memo *Memo) Get(key string) (interface{}, error) {
	response := make(chan result)
	memo.requests <- request{key, response}
	res := <-response
	return res.value, res.err
}

func (memo *Memo) Close() {
	close(memo.requests)
}

type request struct {
	key      string
	response chan<- result
}

type result struct {
	value interface{}
	err   error
}

type Func func(key string) (interface{}, error)

func NewMemo(f Func) *Memo {
	memo := &Memo{requests: make(chan request)}
	go memo.server(f)
	return memo
}

var urls = []string{
	"https://go.dev/play/",
	"https://go.dev/",
	"https://go.dev/learn/",
	"https://go.dev/doc/",
	"https://go.dev/blog/",
	"https://go.dev/play/",
	"https://go.dev/",
	"https://go.dev/learn/",
	"https://go.dev/doc/",
	"https://go.dev/blog/",
	"https://pkg.go.dev/",
	"https://go.dev/solutions/",
	"https://go.dev/play/",
	"https://go.dev/",
	"https://go.dev/learn/",
	"https://go.dev/doc/",
	"https://go.dev/blog/",
	"https://go.dev/play/",
	"https://go.dev/",
	"https://go.dev/learn/",
	"https://go.dev/doc/",
	"https://go.dev/blog/",
	"https://pkg.go.dev/",
	"https://go.dev/solutions/",
	"https://go.dev/play/",
	"https://go.dev/",
	"https://go.dev/learn/",
	"https://go.dev/doc/",
	"https://go.dev/blog/",
	"https://go.dev/play/",
	"https://go.dev/",
	"https://go.dev/learn/",
	"https://go.dev/doc/",
	"https://go.dev/blog/",
	"https://pkg.go.dev/",
	"https://go.dev/solutions/",
}

func main() {
	memo := NewMemo(httpGetBody)
	start := time.Now()
	var n sync.WaitGroup
	for _, url := range urls {
		n.Add(1)
		go func(url string) {
			start := time.Now()
			value, err := memo.Get(url)
			if err != nil {
				log.Println(err)
			}
			fmt.Printf("%s, %s, %d bytes\n", url, time.Since(start), len(value.([]byte)))
			n.Done()
		}(url)
	}
	n.Wait()
	fmt.Printf("total time: %s\n", time.Since(start))
}

func httpGetBody(url string) (interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}
