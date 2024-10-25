package lcme

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

// ConfigRead reads a configuration file and fills the 'config' struct with the values found.
// The file must have lines in the format key=value and the keys must correspond to the fields in the struct.
func ConfigRead(filename string, config interface{}) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("error opening file: %s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		// Ignore empty lines or comments (assuming comments start with #)
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		parts := strings.Split(line, "=")
		if len(parts) != 2 {
			return fmt.Errorf("invalid line: %s", line)
		}
		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		// Check if the value is empty
		if value == "" {
			return fmt.Errorf("value for key %s is empty", key)
		}

		v := reflect.ValueOf(config).Elem()
		field := v.FieldByName(key)
		// If the field doesn't exist, ignore it
		if !field.IsValid() {
			continue
		}

		// Safe type conversion
		switch field.Kind() {
		case reflect.Int:
			intValue, err := strconv.Atoi(value)
			if err != nil {
				return fmt.Errorf("invalid integer value for %s: %s", key, err)
			}
			field.SetInt(int64(intValue))
		case reflect.Int64:
			int64Value, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				return fmt.Errorf("invalid int64 value for %s: %s", key, err)
			}
			field.SetInt(int64Value)
		case reflect.Uint:
			uintValue, err := strconv.ParseUint(value, 10, 64)
			if err != nil {
				return fmt.Errorf("invalid uint value for %s: %s", key, err)
			}
			field.SetUint(uintValue)
		case reflect.Uint64:
			uint64Value, err := strconv.ParseUint(value, 10, 64)
			if err != nil {
				return fmt.Errorf("invalid uint64 value for %s: %s", key, err)
			}
			field.SetUint(uint64Value)
		case reflect.Float32:
			float32Value, err := strconv.ParseFloat(value, 32)
			if err != nil {
				return fmt.Errorf("invalid float32 value for %s: %s", key, err)
			}
			field.SetFloat(float32Value)
		case reflect.Float64:
			float64Value, err := strconv.ParseFloat(value, 64)
			if err != nil {
				return fmt.Errorf("invalid float64 value for %s: %s", key, err)
			}
			field.SetFloat(float64Value)
		case reflect.String:
			field.SetString(value)
		case reflect.Bool:
			boolValue, err := strconv.ParseBool(value)
			if err != nil {
				return fmt.Errorf("invalid boolean value for %s: %s", key, err)
			}
			field.SetBool(boolValue)
		default:
			return fmt.Errorf("unsupported type for %s: %s", key, field.Type().Kind())
		}
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error scanning file: %s", err)
	}

	return nil
}
