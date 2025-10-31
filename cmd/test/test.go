package main

import (
	"fmt"
	"net/http"
	"sync"
)

func main() {
	url := "https://vpnlp.store/admin"

	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		res, err := getResponseStatus(url)

		if err != nil {
			return
		}

		fmt.Println("Get status code: ", res)

		wg.Done()
	}()

	go func() {
		res, err := getResponseStatus(url)

		if err != nil {
			return
		}

		fmt.Println("Get status code: ", res)

		wg.Done()
	}()

	wg.Wait()
}

func getResponseStatus(url string) (int, error) {
	res, err := http.Get(url)

	if err != nil {
		return 0, err
	}

	return res.StatusCode, nil
}
