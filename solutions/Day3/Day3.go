package Day3

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const regex1 = "mul\\(\\d{1,5},\\d{1,5}\\)"

var matches []string
var result int

func Day3() {
	inputFile, err := os.Open("./solutions/Day3/input.txt")
	if err != nil {
		panic(err)
	}
	defer func(inputFile *os.File) {
		err = inputFile.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(inputFile)

	regexExp, err := regexp.Compile(regex1)
	if err != nil {
		panic(err)
	}

	inputScanner := bufio.NewScanner(inputFile)
	for inputScanner.Scan() {
		matches = append(matches, regexExp.FindAllString(inputScanner.Text(), -1)...)
	}
	fmt.Println(matches)
	solve()
	fmt.Println("Sum of all multiplications: ", result)
}

func solve() {
	for _, match := range matches {
		tmp := strings.SplitAfter(match, "mul(")
		tmp1 := strings.ReplaceAll(tmp[1], ")", "")
		tmp2 := strings.Split(tmp1, ",")
		intInput1, err := strconv.Atoi(tmp2[0])
		if err != nil {
			panic(err)
		}
		intInput2, err := strconv.Atoi(tmp2[1])
		if err != nil {
			panic(err)
		}
		result += intInput1 * intInput2
	}
}
