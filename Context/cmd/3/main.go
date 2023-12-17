package main

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"
)

type Processor = func(ctx context.Context) (string, error)

func doWork(workType string, durationSecond int) string {
	fmt.Println("Work with - " + workType)
	time.Sleep(time.Duration(durationSecond) * time.Second)
	return "Result from " + workType
}

func goToRemoteServer(ctx context.Context) (result string, err error) {
	//fmt.Println("RemoteServer is working...")
	//result = doWork("RemoteServer", 2)
	//fmt.Println("RemoteServer done")
	//return result, nil
	return "", errors.New("REMOTE SERVER NOT WORKING")
}

func goToDataBase(ctx context.Context) (result string, err error) {

	if errors.Is(ctx.Err(), context.Canceled) {
		fmt.Println("DataBase cancelled")
		return "", nil
	}

	fmt.Println("DataBase is working...")
	result = doWork("DataBase", 4)
	fmt.Println("DataBase done")
	return result, nil
}

func goToDisk(ctx context.Context) (result string, err error) {

	if errors.Is(ctx.Err(), context.Canceled) {
		fmt.Println("Disk cancelled")
		return "", nil
	}

	fmt.Println("Disk is working...")
	result = doWork("Disk", 8)
	fmt.Println("Disk done")
	return result, nil
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
