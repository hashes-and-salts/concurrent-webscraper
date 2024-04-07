package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {

	urls := []string{
		// sites to scrape
	}

	done := make(chan bool)

	for i, url := range urls {
		go scrape(fmt.Sprintf("%d.txt", i), url, done)
	}

	for range urls {
		<-done
	}

	fmt.Println("Scraping complete ", len(urls), " files written")

}

func scrape(filename string, url string, isComplete chan<- bool) {
	f, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error: ", err)
		isComplete <- false
		return
	}
	defer f.Close()

	response, _ := http.Get(url)

	output, _ := io.ReadAll(response.Body)

	// fmt.Println("", string(output))

	// f.WriteString(string(output))
	f.Write(output)

	fmt.Println("scraping complete, ", filename, " written")
	isComplete <- true
}
