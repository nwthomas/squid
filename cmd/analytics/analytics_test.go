package analytics

import (
	"fmt"
	"testing"
)

func TestTest(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "",
		},
	}

	for _, val := range tests {
		fmt.Println(val)
	}
}
