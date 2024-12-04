package day4

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
)

func readAndReturn(filePath string) []string {

	file, err := os.Open("day4/" + filePath)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	lines := []string{}

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	return lines
}

func Day4Part1(filePath string) {
	lines := readAndReturn(filePath)
	matrix := make([][]rune, len(lines))
	visitedMatrix := make([][]bool, len(lines))

	for i, line := range lines {
		matrix[i] = []rune(line)
		visitedMatrix[i] = make([]bool, len([]rune(line)))
	}

	foundWords := 0

	for i := range matrix {
		for j := range matrix[i] {
			letter := matrix[i][j]
			position := Position[int, int]{i, j}
			if letter == 'X' {
				foundWords += searchForString(matrix, position, "XMAS", UP, NONE)
				foundWords += searchForString(matrix, position, "XMAS", DOWN, NONE)
				foundWords += searchForString(matrix, position, "XMAS", RIGHT, NONE)
				foundWords += searchForString(matrix, position, "XMAS", LEFT, NONE)
				foundWords += searchForString(matrix, position, "XMAS", UP, LEFT)
				foundWords += searchForString(matrix, position, "XMAS", UP, RIGHT)
				foundWords += searchForString(matrix, position, "XMAS", UP, LEFT)
				foundWords += searchForString(matrix, position, "XMAS", UP, RIGHT)
			}
		}
	}

	fmt.Printf("The amount of XMAS is: %v\n", foundWords)

}

type Direction = int

const (
	UP Direction = iota
	DOWN
	LEFT
	RIGHT
	NONE
)

type Position[T, U any] struct {
	X T
	Y U
}

func searchForString(matrix [][]rune, position Position[int, int], str string, directionX Direction, directionY Direction) int {
	currentPosition := Position[int, int]{position.X, position.Y}
	for _, letter := range []rune(str) {
		newLetter := matrix[currentPosition.X][currentPosition.Y]
		if newLetter != letter {
			return 0
		}
		newPosition, err := navigate(matrix, currentPosition, directionX, directionY)
		if err != nil {
			return 0
		}
		currentPosition = newPosition
	}
	return 1
}

func navigate(matrix [][]rune, position Position[int, int], directionX Direction, directionY Direction) (Position[int, int], error) {

	navigatedPosition := navigateOnDirection(navigateOnDirection(position, directionX), directionY)

	isOutOfBoundsX := navigatedPosition.X < 0 || navigatedPosition.X >= len(matrix)
	isOutOfBoundsY := navigatedPosition.Y < 0 || navigatedPosition.Y >= len(matrix[0])

	if isOutOfBoundsX || isOutOfBoundsY {
		return navigatedPosition, errors.New("Out of bounds")
	}
	return navigatedPosition, nil
}

func navigateOnDirection(position Position[int, int], direction Direction) Position[int, int] {
	newPosition := Position[int, int]{position.X, position.Y}
	switch direction {
	case UP:
		newPosition.X--
	case DOWN:
		newPosition.X++
	case LEFT:
		newPosition.Y--
	case RIGHT:
		newPosition.Y++
	case NONE:
	default:
	}
	return newPosition
}
