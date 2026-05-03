package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"
)

func httpGetBody(url string) (interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

type Memo struct {
	f Func
	mu sync.Mutex
	cache map[string]*entry
}

type Func func(key string) (interface{}, error)
type result struct {
	value interface{}
	err error
}

type entry struct {
	res result
	ready chan struct{}
}

func New(f Func) *Memo{
	return &Memo{f: f, cache: make(map[string]*entry)}
}

func (memo *Memo) Get(key string) (value interface{}, err error) {
	memo.mu.Lock()
	e := memo.cache[key]
	if e == nil {
		e = &entry{ready: make(chan struct{})}
		memo.cache[key] = e
		memo.mu.Unlock()

		e.res.value, e.res.err = memo.f(key)
		close(e.ready)
	} else {
		memo.mu.Unlock()
		<-e.ready
	}

	return e.res.value, e.res.err
}

func incomingURLs() []string {
    return []string{
        "https://example.com",
        "https://google.com",
        "https://golang.org",
    }
}

func main() {
	m := New(httpGetBody)
	var n sync.WaitGroup
	urls := incomingURLs()
	
	for _, url := range urls {
		n.Add(1)
		go func(url string) {
			start := time.Now()
			value, err := m.Get(url)
			if err != nil {
				log.Print(err)
			}
			fmt.Printf("%s, %s, %d байтов \n", url, time.Since(start), len(value.([]byte)))
			n.Done()
		}(url)
	}
	n.Wait()
}