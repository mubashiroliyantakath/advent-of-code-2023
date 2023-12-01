package main

import (
	_ "embed"
	"fmt"
	log "github.com/sirupsen/logrus"
	"regexp"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

const firstNumberInStringRegex = `^([a-zA-Z]+)?(\d).*`
const lastNumberInStringRegex = `.*(\d)([a-zA-Z]+)?$`

func init() {
	input = strings.TrimRight(input, "/n")
	if len(input) == 0 {
		log.Fatal("Empty input.txt file.")
	}
}

func main() {
	calibrationValuesList := getCalibrationValues(input)
	calibrationValuesSum := sumListOfInt(calibrationValuesList)
	log.Info("Total sum of the list of integers ", calibrationValuesSum)
}

func getCalibrationValues(input string) []int {
	var calibrationValues []int
	firstNumberPattern := regexp.MustCompile(firstNumberInStringRegex)
	lastNumberPattern := regexp.MustCompile(lastNumberInStringRegex)
	for _, line := range strings.Split(input, "\n") {
		firstNumberMatches := firstNumberPattern.FindStringSubmatch(line)
		fistNumber := firstNumberMatches[2]
		lastNumberMatches := lastNumberPattern.FindStringSubmatch(line)
		lastNumber := lastNumberMatches[1]
		log.Info("The calibration value for this line ", line, " is ", fmt.Sprintf("%s%s", fistNumber, lastNumber))
		intValue, _ := strconv.Atoi(fmt.Sprintf("%s%s", fistNumber, lastNumber))
		calibrationValues = append(calibrationValues, intValue)
	}
	return calibrationValues
}

func sumListOfInt(listOfInt []int) int {
	var sum int
	for _, item := range listOfInt {
		sum = sum + item
	}
	return sum
}
