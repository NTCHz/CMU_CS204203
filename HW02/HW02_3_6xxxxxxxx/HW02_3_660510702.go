// ทิชานนท์ รตนแสนวัน
// 660510702
// HW02_3
// 204203 Sec 001

package main

import (
	"math"
)

func twosComplToInt(x string) int64 {

	var result int64 = 0
	var j int64 = int64(len(x)) - 2
	var front int64 = 0

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
	return front + result
}
