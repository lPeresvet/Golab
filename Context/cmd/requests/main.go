package main

import (
	"Context/core"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	var data core.UsersResponse

	response, err := http.Get("http://localhost:3000/users")
	if err != nil {
		log.Fatal(err)
	}

	if bytes, err := io.ReadAll(response.Body); err == nil {
		err := json.Unmarshal(bytes, &data)
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println(data.Users)
}
