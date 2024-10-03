package utils

import (
	"fmt"
	"strconv"
	"strings"
)

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
