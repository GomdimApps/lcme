package utils

import (
	"fmt"
	"strconv"
	"strings"
)

// PassInts converts a list of strings into integers.
// This function was created to be used by other functions that need to convert strings into integers.
// It receives one or more string values, converts them to integers and returns a slice of integers.
func PassInts(values ...string) ([]int, error) {
	var result []int
	for _, valueStr := range values {
		trimmedValue := strings.TrimSpace(valueStr)
		intValue, err := strconv.Atoi(trimmedValue)
		if err != nil {
			return nil, fmt.Errorf("error converting '%s' to integer: %v", valueStr, err)
		}
		result = append(result, intValue)
	}
	return result, nil
}
