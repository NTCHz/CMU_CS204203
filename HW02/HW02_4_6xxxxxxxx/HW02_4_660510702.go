// ทิชานนท์ รตนแสนวัน
// 660510702
// HW02_4
// 204203 Sec 001

package main

import (
	"math"
	"strings"
)

func twosComplToInt(x string) int64 {

	var result int64 = 0
	var j int64 = int64(len(x)) - 2
	var front int64 = 0

	// calculate the two's complement to integer
	if x[0] == '1' {
		front = (int64(x[0]) - int64('0')) * int64(math.Pow(2, float64(j+1))) * -1
	} else {
		front = (int64(x[0]) - int64('0')) * int64(math.Pow(2, float64(j+1)))
	}
	for i := 1; i < len(x); i++ {
		if x[i] == '1' {
			result += (int64(x[i]) - int64('0')) * int64(math.Pow(2, float64(j)))
		}
		j--
	}

	// return
	return front + result
}

func additiveInverse(x string) (string, int64) {
	var result string = ""
	var carry int64 = 1

	// check if x is all 0
	if strings.Count(x, "0") == len(x) {
		for i := 0; i < len(x); i++ {
			result += "0"
		}
		return result, 0
	}

	// calculate the additive inverse
	for i := len(x) - 1; i >= 0; i-- {
		if x[i] == '1' {
			if carry == 1 {
				result = "1" + result
				carry = 0
			} else {
				result = "0" + result
			}
		} else {
			if carry == 1 {
				result = "0" + result
				carry = 1
			} else {
				result = "1" + result
			}
		}
	}

	if carry == 1 {
		result = "1" + result

	}

	// return
	return result, twosComplToInt(result)
}
