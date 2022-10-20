package main

import (
	"bytes"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// Get the right characters.
func Index(s string, index uint) string {
	runes := bytes.Runes([]byte(s))
	for i, rune := range runes {
		if i == int(index) {
			return string(rune)
		}
	}
	return ""
}

//Perform the calculation.
func compute(num1 float64, symbol string, num2 float64, result []string) []string {
	switch symbol {
	case "+":
		result = append(result, strconv.FormatFloat(num1+num2, 'f', -1, 64))
		fmt.Print(strings.Join(result, " "), "\n")
	case "-":
		result = append(result, strconv.FormatFloat(num1-num2, 'f', -1, 64))
		fmt.Print(strings.Join(result, " "), "\n")
	case "*":
		result = append(result, strconv.FormatFloat(num1*num2, 'f', -1, 64))
		fmt.Print(strings.Join(result, " "), "\n")
	case "/":
		if num2 != 0 {
			result = append(result, strconv.FormatFloat(num1/num2, 'f', -1, 64))
			fmt.Print(strings.Join(result, " "), "\n")
		} else {
			fmt.Print("The denominator cannot be zero.\n")
		}
	default:
		fmt.Print("There was a typographical error.Please enter again:\n")
	}
	return result
}

func main() {

	var result []string
	var symbol string
	var num1, num2 float64
	var i bool = true
	//regnum1 := regexp.MustCompile("\A\d")
	//regnum2 :=regexp.MustCompile("\d\Z")
	fmt.Print("Please enter the calculation (enter the non-calculation or enter to end the program).\n")
	for i {
		fmt.Print("Input:\n")
		/*cond, _ := fmt.Scanf("%f %s %f\n", &num1, &symbol, &num2)
		fmt.Printf("cond = %v\n", cond)*/
		cond, _ := fmt.Scanf("%f%s\n", &num1, &symbol)
		if cond == 2 {
			reg1 := regexp.MustCompile(`^[-+*/]-?\d+$`)
			if reg1 == nil {
				fmt.Println("regexp err")
				return
			} else {
				//根据规则提取关键信息
				result1 := reg1.FindAllStringSubmatch(symbol, -1)
				//fmt.Printf("result1 = %T", result1)

				if result1 != nil {

					//symbol =strconv.Itoa(int(symbol))
					num2, _ = strconv.ParseFloat(symbol[1:], 64)
					symbol = Index(symbol, 0)

					//num1 = strconv.ParseFloat()
					//fmt.Printf("num1 = %v\nsymbol = %v\nnum2 = %v\n", num1, symbol, num2)

					result = compute(num1, symbol, num2, result)
				} else {
					fmt.Print("End\n")
					i = false
				}
			}
		} else {
			fmt.Print("End\n")
			i = false
		}
	}
}
