package main

import (
	"bufio"
	"flag"
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {
	// Define command-line arguments
	targetURL := flag.String("u", "", "The target URL")
	wordlistFile := flag.String("w", "", "The wordlist file")
	delay := flag.Int("delay", 0, "Delay between requests in milliseconds")
	userAgent := flag.String("useragent", "Mozilla/5.0", "User-Agent header for the requests")
	timeout := flag.Int("timeout", 10, "Timeout for the HTTP requests in seconds")
	outputFile := flag.String("o", "", "The output file to write results to")

	// Parse command-line arguments
	flag.Parse()

	// Check if the necessary arguments are provided
	if *targetURL == "" || *wordlistFile == "" {
		fmt.Println("Usage: go run main.go -url <targetURL> -wordlist <wordlistFile> -delay <delayInMilliseconds> -useragent <User-Agent> -timeout <timeoutInSeconds> -output <outputFile>")
		return
	}

	// Open the wordlist file
	file, err := os.Open(*wordlistFile)
	if err != nil {
		fmt.Printf("Error opening wordlist file: %v\n", err)
		return
	}
	defer file.Close()

	// Create a scanner to read the wordlist file line by line
	scanner := bufio.NewScanner(file)

	// Create an HTTP client with the specified timeout
	client := &http.Client{
		Timeout: time.Duration(*timeout) * time.Second,
	}

	// Prepare the output file if specified
	var output *os.File
	if *outputFile != "" {
		output, err = os.Create(*outputFile)
		if err != nil {
			fmt.Printf("Error creating output file: %v\n", err)
			return
		}
		defer output.Close()
	}

	// Loop through each line in the wordlist file
	for scanner.Scan() {
		// Construct the full URL by appending the directory path to the target URL
		url := fmt.Sprintf("%s/%s", *targetURL, scanner.Text())

		// Create a new HTTP request
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			fmt.Printf("Error creating request for %s: %v\n", url, err)
			continue
		}

		// Set the User-Agent header
		req.Header.Set("User-Agent", *userAgent)

		// Send the HTTP request
		resp, err := client.Do(req)
		if err != nil {
			fmt.Printf("Error accessing %s: %v\n", url, err)
			continue
		}

		// Check the response status code
		if resp.StatusCode == http.StatusOK {
			result := fmt.Sprintf("Directory %s exists (Status Code: %d) at %s\n", scanner.Text(), resp.StatusCode, url)
			fmt.Print(result)
			if output != nil {
				_, err := output.WriteString(result)
				if err != nil {
					fmt.Printf("Error writing to output file: %v\n", err)
				}
			}
		}

		// Close the response body to prevent resource leaks
		resp.Body.Close()

		// Sleep for the specified delay duration
		time.Sleep(time.Duration(*delay) * time.Millisecond)
	}

	// Check for scanner errors
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading wordlist file: %v\n", err)
	}
}
