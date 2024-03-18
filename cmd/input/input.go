package input

import "fmt"

type UserInput struct {
	Endpoint       string // Endpoint to load test
	ExecutionTimeS int    // Length of time in seconds to run load test
	RampTimeS      int    // Length of time in seconds to ramp up to max users
	Users          int    // Number of concurrent users to simulate
}

type InputService struct{}

func (i *InputService) GetUserInput() UserInput {
	var (
		endpoint      string
		executionTime int
		rampTime      int
		users         int
	)

	fmt.Println("Enter an endpoint to load test:")
	fmt.Scanln(&endpoint)

	fmt.Println("Enter max number of concurrent users:")
	fmt.Scanln(&users)

	fmt.Println("Enter number of seconds to ramp to max current users:")
	fmt.Scanln(&rampTime)

	fmt.Println("Enter number of seconds to run load test:")
	fmt.Scanln(&executionTime)

	return UserInput{
		ExecutionTimeS: executionTime,
		Endpoint:       endpoint,
		RampTimeS:      rampTime,
		Users:          users,
	}
}
