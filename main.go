package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	// url to scrape
	url := ""
	response, _ := http.Get(url)

	output, _ := io.ReadAll(response.Body)

	fmt.Println("", string(output))
}
