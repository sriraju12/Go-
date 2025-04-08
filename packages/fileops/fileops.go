package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)


func WriteValueFloatToFile(value float64, fileName string) {
	valueText := fmt.Sprintf("%v", value)
	os.WriteFile(fileName, []byte(valueText), 0644)
}

func ReadValueFloatFromFile(fileName string) (float64, error) {
	data, err := os.ReadFile(fileName) // here _ is used to ignore the error

	if err != nil {
		return 1000, errors.New("file not found, creating a new file with default balance of 1000")
	}
	valueText := string(data) // convert byte array to string
	value, err := strconv.ParseFloat(valueText, 64)

	if err != nil {
		return 1000, errors.New("error parsing balance from file")
	}

	return value, nil // here nil is used to indicate that there is no error
}