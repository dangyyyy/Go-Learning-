package datafile

import (
	"bufio"
	"os"
	"strconv"
)

func GetFloats(filename string) ([]float64, error) {

	var numbers []float64

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	scaner := bufio.NewScanner(file)
	for scaner.Scan() {
		number, err := strconv.ParseFloat(scaner.Text(), 64)
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, number)
	}
	err = file.Close()
	if err != nil {
		return nil, err
	}
	if scaner.Err() != nil {
		return nil, scaner.Err()
	}
	return numbers, nil
}
