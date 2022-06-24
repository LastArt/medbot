package pkg

import (
	"medcalcbot/set"
	"regexp"
	"strconv"
	"strings"
)

func FloatToString(input_num float64) string {
	// to convert a float number to a string
	return strconv.FormatFloat(input_num, 'f', 6, 64)
}

func NumberValuator(msgInput string) []float64 {
	var num = []float64{}
	reg := regexp.MustCompile(set.Reg)
	isNumbers := reg.MatchString(msgInput)
	if isNumbers {
		numbers := strings.Split(msgInput, "\\")
		for _, i := range numbers {
			j, err := strconv.ParseFloat(i, 64)
			if err != nil {
				Logger(err)
				continue
			}
			num = append(num, j)
		}
	} else {
		return num
	}
	return num
}

func StringValuator(msgInput string) []string {
	var res = []string{}
	words := strings.Split(msgInput, "\\")
	for _, i := range words {
		res = append(res, i)
	}
	return res
}
