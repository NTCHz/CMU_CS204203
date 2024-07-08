// ทิชานนท์ รตนแสนวัน
// 660510702
// HW02_2
// 204203 Sec 001
package main

import "strings"

const MAX_INT = 64
const DEC_PLACE = 6

func floatToBaseB(x float64, b uint8) string {
	sign := ""

	if x < 0 { // turn negative numbers to positive
		x = -x
		sign = "-"
	}
	// split at the decimal point
	front := int64(x)
	back := x - float64(front)

	frontStr := posIntToBaseB(front, b)
	backStr := fractionToBaseB(back, b)
	// putting every part together
	converted := sign + frontStr + "." + backStr

	return converted

}

func fractionToBaseB(x float64, b uint8) string {
	// this function is working correctly
	if x == 0 {
		return "000000"
	}

	result := [DEC_PLACE]byte{}
	var currDigit float64
	var curr byte
	k := 0

	for ; k < DEC_PLACE; k++ {
		// calculate and convert back to char
		currDigit = x * float64(b)
		curr = byte(int(currDigit) + '0')
		if int(currDigit) > 9 {
			curr = 'A' + curr - '9' - 1
		}
		result[k] = curr
		x = currDigit - float64(int(currDigit))
	}
	// convert from byte array to string
	return string(result[:])
}

func posIntToBaseB(x int64, b uint8) string {
	// this function is working correctly
	if x == 0 {
		return "0"
	}

	result := []byte(strings.Repeat("x", MAX_INT))
	k := MAX_INT - 1
	var currDigit byte

	for x > 0 {
		// calculate and convert back to char
		currDigit = byte((x % int64(b)) + int64('0'))
		if currDigit > '9' {
			currDigit = 'A' + currDigit - '9' - 1
		}
		result[k] = currDigit
		x = x / int64(b)
		k--
	}
	// convert from byte array to string
	return string(result[k+1:])
}
