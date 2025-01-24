package Day3

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

func Day3a() {
	multiplyExprRegex := regexp.MustCompile("mul\\(\\d{1,5},\\d{1,5}\\)")
	var multiplyExprSlice []string
	var sumOfMulExp int

	inputFile, err := os.Open("./solutions/Day3/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(inputFile)

	inputFileReader := bufio.NewScanner(inputFile)

	for inputFileReader.Scan() {
		line := inputFileReader.Text()
		multiplyExprSlice = append(multiplyExprSlice, multiplyExprRegex.FindAllString(line, -1)...)
	}

	for _, expr := range multiplyExprSlice {
		sumOfMulExp += calculateMulExpr(expr)
	}

	fmt.Println("Sum of mul expression: ", sumOfMulExp)
}
