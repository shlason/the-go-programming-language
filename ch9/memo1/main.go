package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"
)

type entry struct {
	res   result
	ready chan struct{}
}

type Memo struct {
	f     Func
	mux   sync.Mutex
	cache map[string]*entry
}

type Func func(key string) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

func NewMemo(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]*entry)}
}

func (memo *Memo) Get(key string) (interface{}, error) {
	memo.mux.Lock()
	e := memo.cache[key]

	if e == nil {
		e = &entry{ready: make(chan struct{})}
		memo.cache[key] = e
		memo.mux.Unlock()

		e.res.value, e.res.err = memo.f(key)

		close(e.ready)
	} else {
		memo.mux.Unlock()
		<-e.ready
	}
	return e.res.value, e.res.err
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
