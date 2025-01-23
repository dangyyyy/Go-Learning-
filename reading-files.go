package datafile

import (
	"bufio"
	"os"
	"strconv"
)

func getfloats(filename string) ([3]float64, error) {

	var numbers [3]float64

	file, err := os.Open(filename)
	if err != nil {
		return numbers, err
	}
	i := 0
	scaner := bufio.NewScanner(file)
	for scaner.Scan() {
		numbers[i], err = strconv.ParseFloat(scaner.Text(), 64)
		if err != nil {
			return numbers, err
		}
		i++
	}
	err = file.Close()
	if err != nil {
		return numbers, err
	}
	if scaner.Err() != nil {
		return numbers, scaner.Err()
	}
	return numbers, nil
}
