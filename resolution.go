package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func calculate(rawsymb string) int {
	find_rim := "(.*)([+|\\-|*|/])(.*)"
	rawsymb = strings.ToUpper(rawsymb)
	rawsymb = strings.ReplaceAll(rawsymb, " ", "")
	reg := regexp.MustCompile(find_rim)
	result := reg.FindStringSubmatch(rawsymb)
	if len(result) == 0 {
		fmt.Println("shit")
		return 0
	}

	action := result[2]
	num0, num1 := parse_int(result[1], result[3])
	switch action {
	case "+":
		return num0 + num1
	case "-":
		return num0 - num1
	case "*":
		return num0 * num1
	case "/":
		return num0 / num1
	default:
		return 0
	}

}

func parse_roman(rawnum string) int {
	var befnum int
	var resultnum int
	var addnum int

	if resultnum := checkint(rawnum); resultnum > 0 {
		return resultnum
	}

	roman_numerals := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	for _, symb := range rawnum {
		addnum = roman_numerals[string(symb)]
		if befnum < addnum {
			resultnum += addnum - befnum*2
		} else {
			resultnum += addnum
		}
		befnum = addnum
	}
	return resultnum
}

func checkint(rawnum string) int {
	num, err := strconv.Atoi(rawnum)
	if err == nil {
		return num
	}
	return 0

}

func parse_int(rawnum0 string, rawnum1 string) (int, int) {
	var num0 int
	var num1 int
	num0 = parse_roman(rawnum0)
	num1 = parse_roman(rawnum1)
	return num0, num1
}

func main() {
	fmt.Println(calculate("100-III"))   // 97
	fmt.Println(calculate("100+III"))   // 103
	fmt.Println(calculate("100*III"))   // 300
	fmt.Println(calculate("100/III"))   // 33
	fmt.Println(calculate("MMXVIII+8")) // 2026
}
