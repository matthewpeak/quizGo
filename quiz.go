package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func readCsvFile(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("unable to read")
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV")
	}
	return records
}
func quiz(csv [][]string) int {
	count := 0
	for i := 0; i < len(csv); i++ {
		fmt.Println(csv[i][0])
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("An error occured while reading input. Please try again", err)

		}
		input = strings.TrimSuffix(input, "\n")
		num, err := strconv.Atoi(input)
		answer, err := strconv.Atoi(csv[i][1])
		if num == answer {
			count += 1
		}
	}
	return count
}
func main() {
	records := readCsvFile("./quiz.csv")
	fmt.Println(quiz(records))
}
