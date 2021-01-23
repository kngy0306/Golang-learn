package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// 本来1000回（ｎ）の数だけサイト訪問が更新されるはずだが、mu.Lock()がないと保証されない


type Visited struct {
	mu      sync.Mutex // visitedマップをガードする
	visited map[string]int
}

func (v *Visited) VisitLink(url string) int {
	//v.mu.Lock()
	//defer v.mu.Unlock()

	// 訪れたサイトの回数を更新する
	count := v.visited[url]
	count++
	v.visited[url] = count
	return count
}

func ClickHinata(v *Visited, c chan int) {
	r := rand.Intn(2) + 1
	time.Sleep(time.Duration(r) * time.Second)

	count := v.VisitLink("hinata.com")
	c <- count
}

func main() {
	hinata := &Visited{visited: map[string]int{"hinata.com": 0}}
	fmt.Println("hinataサイトの訪問回数: ", hinata.visited["hinata.com"])
	c := make(chan int)
	count := 0

	n := 1000
	for i := 0; i < n; i++ {
		go ClickHinata(hinata, c)
	}

	for i := 0; i < n; i++ {
		count = <-c
	}

	fmt.Println("hinataサイトの訪問回数: ", hinata.visited["hinata.com"], count)
}