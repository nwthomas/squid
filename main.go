package main

import (
	"github.com/nwthomas/squid/cmd/input"
	"github.com/nwthomas/squid/cmd/loadtester"
)

func main() {
	inputService := input.InputService{}

	userInput := inputService.GetUserInput()

	loadtestingService := loadtester.LoadTestingService{
		UserInput: userInput,
	}

	loadtestingService.Test()
	loadtestingService.Results()
}
