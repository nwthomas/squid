package analytics

type AnalyticsService struct{}

func (a *AnalyticsService) Test() {
	// finish
}

/*
This package should track the following data:
1. Total number of requests
2. Total errors
3. Total successes

This package should also surface the following data when load testing is done:
1. Mean/Mode/Median averages of response times
2. Slowest/fastest response times
3. Total errors
4. 50th, 90th, 95th, 96th, 97th, 98th, and 99th percentiles
*/
