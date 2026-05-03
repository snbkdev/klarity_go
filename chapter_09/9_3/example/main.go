// 9.3 Мьютексы чтения/записи: sync.RWMutex
package example

import "sync"

var mu sync.RWMutex
var balance int

func Balance() int {
	mu.RLock()
	defer mu.RUnlock()
	return balance
}

