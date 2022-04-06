package file

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func getfile(path string) *os.File {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	return file
}

func ParseFile(path string) ([]int, [][]string) {
	file := getfile(path)
	defer file.Close()
	var registers []int
	var rawCostumers [][]string

	fScanner := bufio.NewScanner(file)
	for fScanner.Scan() {
		line := fScanner.Text()
		if len(line) == 1 {
			r, err := strconv.Atoi(line)
			if err != nil {
				panic(err)
			}
			registers = append(registers, r)
			continue
		}
		rawCostumers = append(rawCostumers, strings.Fields(line))

	}
	if err := fScanner.Err(); err != nil {
		panic(err)
	}

	return registers, rawCostumers
}
