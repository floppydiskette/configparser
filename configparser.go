package configparser

// this file will parse the config file and return the config as a map

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func LoadConfig(fileName string) (map[string]interface{}, error) {
	// open the config file
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// create a map to store the config
	config := make(map[string]interface{})

	// for each line in the config file
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// split the line into key and value
		line := scanner.Text()
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			return nil, fmt.Errorf("invalid config line: %s", line)
		}
		key := strings.TrimSpace(parts[0])
		value := parseValue(parts[1])
		config[key] = value
	}

	// check for errors
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return config, nil
}

func parseValue(value string) interface{} {
	// value without whitespace
	value_c := strings.TrimSpace(value)

	// check if value_c has quotes
	if strings.HasPrefix(value_c, "\"") && strings.HasSuffix(value_c, "\"") {
		// return the content of the quotes
		return strings.Trim(value, "\"")
	}

	// check if the lowercase value is a boolean
	if strings.ToLower(value_c) == "true" || strings.ToLower(value_c) == "false" {
		return strings.ToLower(value_c) == "true"
	}

	// check if the value is an integer
	if i, err := strconv.ParseInt(value_c, 10, 64); err == nil {
		return i
	}

	// check if the value is a float
	if f, err := strconv.ParseFloat(value_c, 64); err == nil {
		return f
	}

	// throw an error otherwise
	return fmt.Errorf("invalid value: %s", value)
}
