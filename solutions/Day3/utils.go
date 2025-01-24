package Day3

import (
	"log"
	"regexp"
	"strconv"
)

var getOperandsRegex = regexp.MustCompile("\\d{1,5}")

func calculateMulExpr(expr string) int {
	multiplyOperands := getOperandsRegex.FindAllString(expr, -1)
	a, err := strconv.Atoi(multiplyOperands[0])
	if err != nil {
		log.Fatal(err)
	}

	b, err := strconv.Atoi(multiplyOperands[1])
	if err != nil {
		log.Fatal(err)
	}

	return a * b
}
