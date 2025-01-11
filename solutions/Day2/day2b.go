package Day2

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
)

var dampedReportStatus []bool
var safeReportsWithDamper = 0

func Day2b() {
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
	for i := 0; i < len(reportList); i++ {
		currentList := reportList[i]
		dampedReportStatus = append(dampedReportStatus, getSafetyStatus(currentList))
	}
	for _, val := range dampedReportStatus {
		if val {
			safeReportsWithDamper += 1
		}
	}
	fmt.Println("Safe Reports after damper function: ", safeReportsWithDamper)
}

func getSafetyStatus(list []int) bool {
	ascStatus := checkForAscendingLevels(list)
	descStatus := checkForDescendingLevels(list)
	levelDifference := checkLevelDifference(list)
	if (ascStatus || descStatus) && levelDifference {
		return true
	} else {
		return checkSafetyWithDamper(list)
	}
}

func checkLevelDifference(list []int) bool {
	for i := 0; i < len(list)-1; i++ {
		if math.Abs(float64(list[i])-float64(list[i+1])) >= 1 && math.Abs(float64(list[i])-float64(list[i+1])) <= 3 {
			continue
		} else {
			return false
		}
	}
	return true
}

func checkSafetyWithDamper(list []int) bool {
	for i := 0; i < len(list); i++ {
		listPart1 := list[:i]
		listPart2 := list[i+1:]
		tmpList := slices.Concat(listPart1, listPart2)
		if (checkForAscendingLevels(tmpList) || checkForDescendingLevels(tmpList)) && checkLevelDifference(tmpList) {
			return true
		} else {
			continue
		}
	}
	return false
}
