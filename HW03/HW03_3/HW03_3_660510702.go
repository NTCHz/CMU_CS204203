// ทิชานนท์ รตนแสนวัน
// 660510702
// HW03_3
// 204203 Sec 001

package main

import (
	// "strings"

	"math"
	"strconv"
	"strings"
	// "fmt"""
)

const fracLen = 7
const expLen = 8

const BASE uint8 = 2
const DEBUG = false

// from HW02_2
const MAX_INT = 128
const DEC_PLACE = 256

func float16bitNormed(n float32) string {

	sign := "0"
	if n < 0 {
		n = -n
		sign = "1"
	}

	// expLen = 8, fracLen = 7
	var bias = int(pow(2, expLen-1) - 1) // bias = 127
	nre := float64(n)
	if DEBUG {
		println("Bias", bias)
	}
	var minNorm float64 = 0.0000000000000000000000000000000000000117 // dummy value
	var maxNorm float64 = 338953138925153547590470800371487866880    // dummy value

	floatToBase2 := floatToBaseB(nre, 2)

	idx := strings.Index(floatToBase2, ".")

	xxx := 1
	if strings.Index(floatToBase2, "1") > idx {
		xxx = 0
	}

	E := idx - strings.Index(floatToBase2, "1") - xxx

	if DEBUG {
		println("E", E)
		println("fl", floatToBase2)
		println("EXP", int64(E+bias))
		println("idx", strings.Index(floatToBase2, "1"))
	}
	EXP := posIntToBaseB(float64(E+bias), 2)

	if DEBUG {
		println("EXP", EXP)
	}

	if (float64(n) > maxNorm) || (float64(n) < minNorm) {
		if DEBUG {
			println(n, "overflow")
		}
		return ""
	}

	floatToBase2 = strings.Replace(floatToBase2, ".", "", 1)
	result := floatToBase2[strings.Index(floatToBase2, "1")+1:]
	result = result + strings.Repeat("0", fracLen)
	return sign + " " + strings.Repeat("0", 8-len(EXP)) + EXP + " " + result[:fracLen]
}

func pow(x, y int) float64 {
	return math.Pow(float64(x), float64(y))
}

// -------------------------------------------- //from HW02_2
func floatToBaseB(x float64, b uint8) string {
	sign := ""

	if x < 0 {
		x = -x
		sign = "-"
	}
	front, back := math.Modf(x)
	frontStr := posIntToBaseB(front, b)
	backStr := fractionToBaseB(back, b)
	converted := sign + frontStr + "." + backStr

	return converted

}

func fractionToBaseB(x float64, b uint8) string {
	if x == 0 {
		return "0"
	}

	result := [DEC_PLACE]byte{}
	var currDigit float64
	var curr byte
	k := 0

	for ; k < DEC_PLACE; k++ {
		currDigit = x * float64(b)
		curr = byte(int(currDigit) + '0')
		if int(currDigit) > 9 {
			curr = 'A' + curr - '9' - 1
		}
		result[k] = curr
		x = currDigit - float64(int(currDigit))
	}
	return string(result[:])
}

func posIntToBaseB(x float64, b uint8) string {
	if x == 0 {
		return "0"
	}

	result := ""

	for i := 0; i <= MAX_INT; i++ {
		currDigit := strconv.FormatInt(int64(math.Mod(float64(x), float64(b))), 10)
		result = string(currDigit) + result
		x = x / float64(b)
	}

	result = result[strings.Index(result, "1"):]
	return result
}
