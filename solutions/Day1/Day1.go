package Day1

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func Day1() {
	//Variable Declarations
	var leftList, rightList, distance []int
	var totalDistance int

	//Read source data and create slices for processing
	fileData, err := os.Open("./AdventofCode2024-Go/solutions/Day1/LeftList.txt")
	if err != nil {
		fmt.Println(err)
	}
	listScanner := bufio.NewScanner(fileData)
	for listScanner.Scan() {
		tmpIntVal, err := strconv.Atoi(strings.TrimSpace(listScanner.Text()))
		if err != nil {
			fmt.Println(err)
		}
		leftList = append(leftList, tmpIntVal)
	}
	fileData, err = os.Open("./AdventofCode2024-Go/solutions/Day1/RightList.txt")
	if err != nil {
		fmt.Println(err)
	}
	listScanner = bufio.NewScanner(fileData)
	for listScanner.Scan() {
		tmpIntVal, err := strconv.Atoi(strings.TrimSpace(listScanner.Text()))
		if err != nil {
			fmt.Println(err)
		}
		rightList = append(rightList, tmpIntVal)
	}

	//sort slices
	slices.Sort(leftList)
	slices.Sort(rightList)

	//calculate distance
	for index, leftListVal := range leftList {
		if leftListVal > rightList[index] {
			distance = append(distance, leftListVal-rightList[index])
		} else {
			distance = append(distance, rightList[index]-leftListVal)
		}
	}
	for _, val := range distance {
		totalDistance += val
	}
	fmt.Println("Total distance: ", totalDistance)
}
