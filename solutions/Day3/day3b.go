package Day3

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

func Day3b() {
	multiplyExprRegex := regexp.MustCompile("mul\\(\\d{1,5},\\d{1,5}\\)")
	var multiplyExprSlice []string
	//var inputData []string
	var cleanedInputSlice []string
	//var splitByNewLineInputLines []string
	var sumOfMulExp int

	inputData, err := os.ReadFile("./solutions/Day3/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	inputData = []byte(strings.ReplaceAll(string(inputData), "\n", " "))
	inputData = []byte(strings.ReplaceAll(string(inputData), "don't", "\ndon't"))
	inputData = []byte(strings.ReplaceAll(string(inputData), "do", "\ndo"))
	cleanedInputSlice = strings.SplitN(string(inputData), "\n", -1)

	for _, exp := range cleanedInputSlice {
		if !strings.HasPrefix(exp, "don't") {
			multiplyExprSlice = append(multiplyExprSlice, multiplyExprRegex.FindAllString(exp, -1)...)
		}
	}

	for _, exp := range multiplyExprSlice {
		sumOfMulExp += calculateMulExpr(exp)
	}

	fmt.Println("Sum of enabled multiply expressions:", sumOfMulExp)

}
