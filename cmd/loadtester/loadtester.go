package loadtester

import (
	"github.com/nwthomas/squid/cmd/input"
)

type LoadTestingService struct {
	UserInput input.UserInput
}

func (l *LoadTestingService) GetRampUserIncrementsS() int {
	// Period of time before starting up another concurrent user
	waitTime := l.UserInput.RampTimeS / l.UserInput.Users

	return waitTime
}

func (l *LoadTestingService) Test() {
	// finish
}

func (l *LoadTestingService) Results() {
	// finish
}
