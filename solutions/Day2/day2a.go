package Day2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reportList [][]int
var reportSafetyStatus map[int]bool

func Day2a() {
	reportSafetyStatus = make(map[int]bool)
	inputFile, err := os.Open("./solutions/Day2/input.txt")
	if err != nil {
		fmt.Println("Error opening input file: ", err.Error())
	}
	defer func(fileToClose *os.File) {
		err := inputFile.Close()
		if err != nil {
			fmt.Println("Error closing input file: ", err.Error())
		}
	}(inputFile)

	inputFileScanner := bufio.NewScanner(inputFile)
	for inputFileScanner.Scan() {
		currentLevelAsInt := convertRawStringToInt(inputFileScanner.Text())
		reportList = append(reportList, currentLevelAsInt)
	}
	checkSafetyStatus()
	fmt.Println("Safe Report Count: ", getSafeReportCount())
}

func getSafeReportCount() int {
	var tmpCount int
	for _, v := range reportSafetyStatus {
		if v {
			tmpCount++
		}
	}
	return tmpCount
}

func checkSafetyStatus() {
	for index, value := range reportList {
		if (checkForAscendingLevels(value) && checkLevelDifferenceForAscending(value)) || (checkForDescendingLevels(value) && checkLevelDifferenceForDescending(value)) {
			reportSafetyStatus[index] = true
		} else {
			reportSafetyStatus[index] = false
		}
	}
}

func checkForAscendingLevels(valueList []int) bool {
	for index, value := range valueList {
		if index == len(valueList)-1 {
			return true
		} else {
			if value < valueList[index+1] {
				continue
			} else {
				return false
			}
		}
	}
	if checkLevelDifferenceForAscending(valueList) {
		return true
	} else {
		return false
	}
}

func checkForDescendingLevels(valueList []int) bool {
	for index, value := range valueList {
		if index == len(valueList)-1 {
			return true
		} else {
			if value > valueList[index+1] {
				continue
			} else {
				return false
			}
		}
	}
	return false
}

func checkLevelDifferenceForAscending(valueList []int) bool {
	for i := 0; i < len(valueList)-1; i++ {
		if valueList[i+1]-valueList[i] >= 1 && valueList[i+1]-valueList[i] <= 3 && valueList[i+1]-valueList[i] > 0 {
			continue
		} else {
			return false
		}
	}
	return true
}

func checkLevelDifferenceForDescending(valueList []int) bool {
	for i := 0; i < len(valueList)-1; i++ {
		if valueList[i]-valueList[i+1] >= 1 && valueList[i]-valueList[i+1] <= 3 && valueList[i]-valueList[i+1] > 0 {
			continue
		} else {
			return false
		}
	}
	return true
}

func convertRawStringToInt(rawString string) []int {
	tmpStrSlice := strings.Split(rawString, " ")
	var tmpIntSlice []int
	for _, value := range tmpStrSlice {
		tmpIntValue, err := strconv.Atoi(value)
		if err != nil {
			panic(err)
		}
		tmpIntSlice = append(tmpIntSlice, tmpIntValue)
	}
	return tmpIntSlice
}
