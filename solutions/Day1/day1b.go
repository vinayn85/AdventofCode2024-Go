package Day1

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"AdventofCode2024-Go/solutions/utils"
)

func Day1b() {
	var leftInputList []int
	var rightInputList []int
	var similarityScore []int
	var tmpSimilarityCnt int

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

	//Calculating similarity score
	for _, value := range leftInputList {
		tmpSimilarityCnt = 0
		for _, rightListValue := range rightInputList {
			if value == rightListValue {
				tmpSimilarityCnt++
			}
		}
		similarityScore = append(similarityScore, tmpSimilarityCnt*value)
	}

	fmt.Println("Similarity Score:", utils.CalcSliceSum(similarityScore))
}
