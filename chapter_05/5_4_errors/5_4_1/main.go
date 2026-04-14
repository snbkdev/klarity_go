// 5.4.1 Стратегии обработки ошибок
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func WaitServer(url string) error {
	const timeout = 1 * time.Minute
	deadline := time.Now().Add(timeout)
	for tries := 0; time.Now().Before(deadline); tries++ {
		resp, err := http.Head(url)
		if err == nil {
			resp.Body.Close()
			return nil
		}

		log.Printf("Сервер не отвечает (%s); повтор ...", err)
		time.Sleep(time.Second << uint(tries))
	}

	return fmt.Errorf("Сервер %s не отвечет, время %s", url, timeout)
}

func main() {
	url := "https://www.google.com"
	if err := WaitServer(url); err != nil {
		fmt.Fprintf(os.Stderr, "Сервер не работает %v\n", err)
		os.Exit(1)
	}

	if err := WaitServer(url); err != nil {
		log.Fatalf("Сервер не работает %v\n", err)
	}

}