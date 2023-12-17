package main

import (
	"fmt"
	"time"
)

type Processor = func() (string, error)

func doWork(workType string, durationSecond int) string {
	fmt.Println("Work with - " + workType)
	time.Sleep(time.Duration(durationSecond) * time.Second)
	return "Result from " + workType
}

func goToRemoteServer() (result string, err error) {
	fmt.Println("RemoteServer is working...")
	result = doWork("RemoteServer", 2)
	fmt.Println("RemoteServer done")
	return result, nil
}

func goToDataBase() (result string, err error) {
	fmt.Println("DataBase is working...")
	result = doWork("DataBase", 4)
	fmt.Println("DataBase done")
	return result, nil
}

func goToDisk() (result string, err error) {
	fmt.Println("Disk is working...")
	result = doWork("Disk", 8)
	fmt.Println("Disk done")
	return result, nil
}

func processRequest(processors []Processor) {
	for _, processor := range processors {
		result, err := processor()

		if err != nil {
			fmt.Println("Error - " + err.Error())
		} else if result != "" {
			fmt.Println("-> " + result)
		}
	}
}

func main() {
	processors := []Processor{goToRemoteServer, goToDataBase, goToDisk}

	processRequest(processors)
}
