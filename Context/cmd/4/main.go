package main

import (
	"Context/core"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

type Processor = func(ctx context.Context) (string, error)

func doWork(workType string, durationSecond int, result chan<- string) {
	fmt.Println("Work with - " + workType)
	time.Sleep(time.Duration(durationSecond) * time.Second)
	result <- "Result from " + workType
}

func doRequest(workType string, result chan<- string) error {
	fmt.Println("Work with - " + workType)
	var data core.UsersResponse

	response, err := http.Get("http://localhost:3000/error")
	if err != nil {
		return err
	}
	if response.StatusCode >= 500 {
		return errors.New("server error")
	}

	if bytes, err := io.ReadAll(response.Body); err == nil {
		err := json.Unmarshal(bytes, &data)
		if err != nil {
			return err
		}
	}

	result <- fmt.Sprintf("Result from %v: %v", workType, data)
	return nil
}

func goToRemoteServer(ctx context.Context) (result string, err error) {
	ctxWithTimeOut, cancel := context.WithTimeout(ctx, 1*time.Second)
	fmt.Println("RemoteServer is working...")

	resultChannel := make(chan string)

	go func() {
		err := doRequest("RemoteServer", resultChannel)
		if err != nil {
			cancel()
		}
	}()

	for {
		select {
		case <-ctxWithTimeOut.Done():
			if errors.Is(ctxWithTimeOut.Err(), context.Canceled) {
				fmt.Println("RemoteServer cancelled")
			}
			return "", ctxWithTimeOut.Err()
		case result = <-resultChannel:
			fmt.Println("RemoteServer done")
			return result, nil
		}
	}
}

func goToDataBase(ctx context.Context) (result string, err error) {
	fmt.Println("DataBase is working...")

	resultChannel := make(chan string)

	go doWork("DataBase", 4, resultChannel)

	for {
		select {
		case <-ctx.Done():
			if errors.Is(ctx.Err(), context.Canceled) {
				fmt.Println("DataBase cancelled")
				return "", nil
			}
		case result = <-resultChannel:
			fmt.Println("DataBase done")
			return result, nil
		}
	}
}

func goToDisk(ctx context.Context) (result string, err error) {
	fmt.Println("Disk is working...")

	resultChannel := make(chan string)

	go doWork("Disk", 8, resultChannel)

	for {
		select {
		case <-ctx.Done():
			if errors.Is(ctx.Err(), context.Canceled) {
				fmt.Println("Disk cancelled")
				return "", nil
			}
		case result = <-resultChannel:
			fmt.Println("Disk done")
			return result, nil
		}
	}
}

func processRequest(ctx context.Context, processors []Processor) {
	ctxWithCancel, cancel := context.WithCancel(ctx)

	var wg = &sync.WaitGroup{}
	wg.Add(len(processors)) // i = 3

	for _, processor := range processors {
		go func(p Processor) {
			result, err := p(ctxWithCancel)

			if err != nil {
				fmt.Println("Error - " + err.Error())
				cancel()
			} else if result != "" {
				fmt.Println("-> " + result)
			}

			wg.Done() // i--
		}(processor) // нужно, чтобы каждый раз в цикле была новая функция
	}

	wg.Wait() // ждем, пока i != 0

	cancel()
}

func main() {
	ctx := context.Background()

	processors := []Processor{goToRemoteServer, goToDataBase, goToDisk}

	processRequest(ctx, processors)
}
