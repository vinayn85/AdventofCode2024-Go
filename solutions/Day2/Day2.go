package Day2

//TODO Need to implement damper functionality fully

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"slices"
	"strconv"
	"strings"
)

var sourceLists [][]int
var reportStatus = make(map[int]string)
var safeReportCount int
var dampedSafeReportCount int
var dampedLists [][]int

func Day2() {
	createInputListsFromFile()
	verifyReportSafetyStatus(sourceLists)
	for _, val := range reportStatus {
		if val == "safe" {
			safeReportCount++
		}
	}
	fmt.Println("Original Safe report count: ", safeReportCount)
	//fmt.Println(reportStatus)
	applyDamper()
	verifyReportSafetyStatus(dampedLists)
	for _, val := range reportStatus {
		if val == "safe" {
			dampedSafeReportCount++
		}
	}
	fmt.Println("Safe report count after damper: ", dampedSafeReportCount)

	//fmt.Println(reportStatus)
}

func applyDamper() {
	for _, list := range sourceLists {
		fmt.Println(list)
		for i := 0; i < len(list)-1; i++ {
			if !(list[i+1]-list[i] >= 1 || list[i+1]-list[i] <= 3 && list[i+1]-list[i] > 0) {
				list = append(list[:i], list[i+1:]...)
				fmt.Println(list)
				dampedLists = append(dampedLists, list)
				break
			}
		}
	}
}

func verifyReportSafetyStatus(inputLists [][]int) {
	//verify if ascending or descending then check if the differences between levels is acceptable
	for index, val := range inputLists {
		var t = make([]int, len(val))
		copy(t, val) // creating a copy to compare cleanly
		slices.Sort(t)
		if reflect.DeepEqual(val, t) { // using reflection here. can write own function to do the same for production
			reportStatus[index] = "safe"
			a := checkAdjacentLevelDifferenceforAscending(val)
			if !a {
				reportStatus[index] = "unsafe"
			}
		} else {
			slices.Reverse(t)
			if reflect.DeepEqual(val, t) {
				reportStatus[index] = "safe"
				a := checkAdjacentLevelDifferenceforDescending(val)
				if !a {
					reportStatus[index] = "unsafe"
				}
			}
		}
	}
}

func checkAdjacentLevelDifferenceforAscending(inputSlice []int) bool {
	for index, val := range inputSlice {
		if index == len(inputSlice)-1 {
			return true
		} else {
			if inputSlice[index+1]-val >= 1 && inputSlice[index+1]-val <= 3 {
				continue
			} else {
				return false

			}
		}
	}
	return false
}

func checkAdjacentLevelDifferenceforDescending(inputSlice []int) bool {
	for index, val := range inputSlice {
		if index == len(inputSlice)-1 {
			return true
		} else {
			if val-inputSlice[index+1] >= 1 && val-inputSlice[index+1] <= 3 {
				continue
			} else {
				return false
			}
		}
	}
	return false
}

func createInputListsFromFile() {
	fileData, err := os.Open("./solutions/Day2/input.txt")
	if err != nil {
		panic(err)
	}
	inputScanner := bufio.NewScanner(fileData)
	for inputScanner.Scan() {
		currentLine := inputScanner.Text()
		currentLineAsSlice := strings.Split(currentLine, " ")
		sourceLists = append(sourceLists, convertStringSliceToIntSlice(currentLineAsSlice))
	}
}

func convertStringSliceToIntSlice(inputSlice []string) []int {
	var outputSlice []int
	for _, val := range inputSlice {
		tmpIntVal, err := strconv.Atoi(val)
		if err != nil {
			panic(err)
		}
		outputSlice = append(outputSlice, tmpIntVal)
	}
	return outputSlice
}
