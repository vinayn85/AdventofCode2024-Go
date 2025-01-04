package Day1

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"

	"AdventofCode2024-Go/solutions/utils"
)

func Day1a() {
	var leftInputList []int
	var rightInputList []int
	var distance []int

	spaceRegex := regexp.MustCompile("\\s+")

	inputFile, err := os.Open("./solutions/Day1/input.txt")
	if err != nil {
		panic(err)
	}
	defer func(fileToClose *os.File) {
		err := fileToClose.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(inputFile)

	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		currentLine := scanner.Text()
		currentLineWithSingleSpace := spaceRegex.ReplaceAllString(currentLine, " ")
		currentLineSplit := strings.Split(currentLineWithSingleSpace, " ")

		tmpVar1, err := strconv.Atoi(strings.TrimSpace(currentLineSplit[0]))
		if err != nil {
			panic(err)
		}
		leftInputList = append(leftInputList, tmpVar1)

		tmpVar2, err := strconv.Atoi(strings.TrimSpace(currentLineSplit[1]))
		if err != nil {
			panic(err)
		}
		rightInputList = append(rightInputList, tmpVar2)
	}

	// sort the slices
	slices.Sort(leftInputList)
	slices.Sort(rightInputList)

	//Calculate Distance
	for index, value := range leftInputList {
		if value > rightInputList[index] {
			distance = append(distance, value-rightInputList[index])
		} else {
			distance = append(distance, rightInputList[index]-value)
		}
	}

	fmt.Println("Total distance:", utils.CalcSliceSum(distance))
}
