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

func main() {
	file, err := os.ReadFile("input.txt")
	checkError(err)

	lines := strings.Split(string(file), "\n")

	dial := 50
	zeroCount := 0

	for _, line := range lines {
		rotation := getRotation(line)
		dial = rotateDial(dial, rotation)
		fmt.Printf("Dial: %d\n", dial)
		if dial == 0 {
			zeroCount += 1
		}

	}

	fmt.Printf("Zero count: %d\n", zeroCount)
}

type Direction string

const (
	DirLeft  Direction = "L"
	DirRight Direction = "R"
)

func (s Direction) String() string {
	return string(s)
}

type Rotation struct {
	direction Direction
	clicks    int
}

func getRotation(input string) Rotation {
	clicks, err := strconv.Atoi(input[1:])
	if err != nil {
		panic(err)
	}

	return Rotation{
		direction: Direction(input[0]),
		clicks:    clicks,
	}
}

func rotateDialModulo(dial int, rotation Rotation) int {
	switch rotation.direction {
	case DirLeft:
		dial = ((dial-rotation.clicks)%10 + 10) % 10
	case DirRight:
		dial = (dial + rotation.clicks) % 10
	}

	return dial

}

func rotateDial(dial int, rotation Rotation) int {
	min := 0
	max := 99

	var dialval int

	if rotation > 100

	switch rotation.direction {
	case DirLeft:
		dialval = dial - rotation.clicks
	case DirRight:
		dialval = dial + rotation.clicks
	}

	switch {
	// exceeds max bounds
	case dialval > max:
		return dialval - max - 1
	// exceeds min bounds
	case dialval < min:
		return dialval + max + 1
	default:
		return dialval
	}
}
