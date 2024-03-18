package main

import (
	"fmt"
	"net/http"
	"sync"
)

func main() {
	var endpoint string

	fmt.Println("Please enter the API endpoint to load test:")
	fmt.Scanln(&endpoint)


	var wg sync.WaitGroup
	concurrentUsers := 100 // Adjust the number of concurrent users
	requestsPerUser := 10  // Adjust the number of requests per user

	for i := 0; i < concurrentUsers; i++ {
		wg.Add(1)
		go func(userID int) {
			defer wg.Done()
			for j := 0; j < requestsPerUser; j++ {
				response, err := http.Get(endpoint) // Replace with the target API URL
				if err != nil {
					fmt.Printf("User %d: Error - %s\n", userID, err)
					continue
				}
				response.Body.Close()
				fmt.Printf("User %d: Request %d - Status Code: %d\n", userID, j+1, response.StatusCode)
			}
		}(i)
	}

	wg.Wait()
	fmt.Println("Load testing completed.")
}
