package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func getListSum(list []int) int {
	sum := 0

	for _, i := range list {
		sum += i
	}

	return sum
}

func isIdValid(id string) bool {
	mid := len(id) / 2

	firstHalf := id[:mid]
	secondHalf := id[mid:]

	return firstHalf != secondHalf

}

func main() {
	file, err := os.ReadFile("input.txt")
	checkError(err)

	var invalidIds []int
	var validIds []int

	stringRanges := strings.Split(string(file), ",")

	for _, stringRange := range stringRanges {

		rangebounds := strings.Split(stringRange, "-")
		lower, err := strconv.Atoi(rangebounds[0])
		checkError(err)
		upper, err := strconv.Atoi(rangebounds[1])
		checkError(err)

		for i := lower; i <= upper; i++ {
			if !isIdValid(strconv.Itoa(i)) {
				invalidIds = append(invalidIds, i)
			}

		}
	}

	fmt.Printf("invalidIds: %d\n", getListSum(invalidIds))
	fmt.Printf("validIds: %d\n", getListSum(validIds))
}
