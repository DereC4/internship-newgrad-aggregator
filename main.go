package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type JobListing struct {
	Company     string
	Role        string
	Location    string
	Application string
	Age         string
}

func dogWorker(url string, ch chan string) {
	// function signature in Go is variableName dataType
	// channels are type safe in Go so you have to define what type a channel takes
	response, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprintf("Error fetching %s: %v", url, err)
		return
	}

	bodyBytes, err := io.ReadAll(response.Body)

	if err != nil {
		ch <- fmt.Sprintf("Error reading body for %s: %v", url, err)
		return
	}

	ch <- string(bodyBytes)

}

func main() {
	urls := []string{
		// this is a slice
		"https://raw.githubusercontent.com/vanshb03/Summer2027-Internships/dev/README.md",
		"https://raw.githubusercontent.com/SimplifyJobs/Summer2026-Internships/refs/heads/dev/README.md",
	}

	resultsChannel := make(chan string)
	// make a channel type so we can talk to main
	// loop through urls and start a thread for each one
	fmt.Println("Starting fetches")
	for _, url := range urls {
		// underscore is so we ignore the index, discard it
		go dogWorker(url, resultsChannel)
	}

	// you have to open the file before reading from channel
	file, err := os.OpenFile("testing.md", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}

	for i := 0; i < len(urls); i++ {
		results := <-resultsChannel
		// channels will get consumed when you read them all one by one, so our two for loop approach was writing nothing
		fmt.Printf("--- Document Received #%d ---\n", i+1)
		fmt.Println(results)
		fmt.Println("-------------------------------")

		separator := fmt.Sprintf("\n\n# --- Document Received #%d ---\n\n", i+1)

		if _, err := file.WriteString(separator); err != nil {
			fmt.Printf("Error writing separator to file: %v\n", err)
		}

		if _, err := file.WriteString(results); err != nil {
			fmt.Printf("Error writing content to file: %v\n", err)
		}

		fmt.Printf("Saved document #%d to combined_output.md\n", i+1)
	}

	for i := 0; i < len(urls); i++ {
	}

}
