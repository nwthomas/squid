package input

import (
	"fmt"
	"testing"
)

func TestGetUserInput(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "Returns default struct values when no user input",
		},
	}

	for _, val := range tests {
		fmt.Println(val)
	}
}
