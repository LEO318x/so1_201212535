package controller

import "strconv"

func RealizarDivision(num1 int64, num2 int64) string {
	if num2 == 0 {
		return "\"ERR\""
	} else {
		return strconv.FormatInt(num1/num2, 10)
	}
}
