package input

import "fmt"

type UserInput struct {
	Url     string // Endpoint to load test
	Workers int    // Number of concurrent workers to use in test
}

func GetUserInput() UserInput {
	var endpoint string
	var workers int

	fmt.Println("Please enter the API endpoint to load test:")
	fmt.Scanln(&endpoint)

	fmt.Println("Please enter the max number of concurrent requests:")
	fmt.Scanln(&workers)

	return UserInput{
		Url:     endpoint,
		Workers: workers,
	}
}
