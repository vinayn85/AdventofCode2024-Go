package Day1a

import (
	"encoding/csv"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

var leftList, rightList, distance []int
var totalDistance int

func Day1a() {
	csvRecords := parseCsvFile()
	generateLeftAndRightLists(csvRecords)
	slices.Sort(leftList)
	slices.Sort(rightList)
	calculateIndividualDistances()
	calucateTotalDistance()
}

func parseCsvFile() [][]string {
	inputFile, err := os.Open("./solutions/Day1a/input.txt")
	if err != nil {
		panic(err)
	}
	csvReader := csv.NewReader(inputFile)
	csvReader.Comma = ' ' // sticking with input text as copied.
	csvReader.FieldsPerRecord = -1

	csvRecords, err := csvReader.ReadAll()
	if err != nil {
		panic(err)
	}
	return csvRecords
}

func generateLeftAndRightLists(records [][]string) {
	for _, record := range records {
		tmpStrVal := reduceDelimiterToSingleSpace(record)
		tmpSlice := strings.Split(tmpStrVal, " ")
		tmpIntVal, err := strconv.Atoi(tmpSlice[0])
		if err != nil {
			panic(err)
		}
		leftList = append(leftList, tmpIntVal)
		tmpIntVal, err = strconv.Atoi(tmpSlice[1])
		if err != nil {
			panic(err)
		}
		rightList = append(rightList, tmpIntVal)
	}
}

func reduceDelimiterToSingleSpace(s []string) string {
	tmpStrVal := strings.Join(s, " ")
	spaceCount := strings.Count(tmpStrVal, " ")
	return strings.Replace(tmpStrVal, " ", "", spaceCount-1)
}

func calculateIndividualDistances() {
	for index, value := range leftList {
		if value > rightList[index] {
			distance = append(distance, value-rightList[index])
		} else {
			distance = append(distance, rightList[index]-value)
		}
	}
}

func calucateTotalDistance() {
	for _, val := range distance {
		totalDistance += val
	}
	fmt.Println("Total distance: ", totalDistance)

}
