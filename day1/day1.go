package day1

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func readAndReturn(filePath string) ([]int, []int) {

	file, err := os.Open("day1/" + filePath)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	firstItems := []int{}
	secondItems := []int{}

	for scanner.Scan() {
		line := scanner.Text()
		splitLine := strings.Split(line, "   ")
		if len(splitLine) == 0 {
			continue
		}
		first, err := strconv.Atoi(splitLine[0])
		second, err := strconv.Atoi(splitLine[1])

		if err != nil {
			fmt.Println(err)
			return nil, nil
		}

		firstItems = append(firstItems, first)
		secondItems = append(secondItems, second)
	}
	return firstItems, secondItems
}

func Day1Part1(filePath string) {

	firstItems, secondItems := readAndReturn(filePath)

	slices.Sort(firstItems)
	slices.Sort(secondItems)

	length := len(firstItems)
	sum := 0
	for i := range length {
		sum += int(math.Abs(float64(firstItems[i]) - float64(secondItems[i])))
	}

	fmt.Printf("The sum is: %v\n", sum)

}

func Day1Part2(filePath string) {

	firstItems, secondItems := readAndReturn(filePath)
	similarityScoreMap := map[int]int{}
	for _, v := range firstItems {
		similarityScoreMap[v] = 0
	}
	for _, v := range secondItems {
		if _, ok := similarityScoreMap[v]; ok {
			similarityScoreMap[v]++
		}
	}
	similarityScore := 0

	for _, item := range firstItems {
		similarityScore += item * similarityScoreMap[item]
	}
	fmt.Printf("The similarity score value is: %v\n", similarityScore)
}
