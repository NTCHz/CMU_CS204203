// ทิชานนท์ รตนแสนวัน
// 660510702
// HW03_2
// 204203 Sec 001

package main

import (
	"strconv"
	"strings"
)

const MAX = 36 // really?

func baseNAddition(r1, r2 string, base int) string {
	decPointPos1 := strings.Index(r1, ".")
	decPointPos2 := strings.Index(r2, ".")

	if (r1 == "" && r2 == "") || (r1 == "0" && r2 == "0") {
		return "0"
	}

	if r1 == "." && r2 == "." {
		return "0.0"
	}

	r1re := [2]string{}
	r2re := [2]string{}

	if decPointPos1 != -1 {
		r1re[0] = r1[:decPointPos1]
		r1re[1] = r1[decPointPos1+1:]
		if r1re[1] == "" {
			r1re[1] = "0"
		}
	} else {
		if r1 == "." {
			r1re[0] = "0"
			r1re[1] = "0"
		} else {
			r1re[0] = r1
			r1re[1] = ""
		}
	}

	if decPointPos2 != -1 {
		r2re[0] = r2[:decPointPos2]
		r2re[1] = r2[decPointPos2+1:]
		if r2re[1] == "" {
			r2re[1] = "0"
		}
	} else {
		if r2 == "." {
			r2re[0] = "0"
			r2re[1] = "0"
		} else {
			r2re[0] = r2
			r2re[1] = ""
		}
	}

	if len(r1re[0]) > len(r2re[0]) {
		r2re[0] = strings.Repeat("0", len(r1re[0])-len(r2re[0])) + r2re[0]
	} else {
		r1re[0] = strings.Repeat("0", len(r2re[0])-len(r1re[0])) + r1re[0]
	}

	if len(r1re[1]) > len(r2re[1]) {
		r2re[1] = r2re[1] + strings.Repeat("0", len(r1re[1])-len(r2re[1]))
	} else {
		r1re[1] = r1re[1] + strings.Repeat("0", len(r2re[1])-len(r1re[1]))
	}

	a := string(r1re[0]) + string(r1re[1])
	b := string(r2re[0]) + string(r2re[1])

	i := len(a) - 1

	x := make([]byte, len(a))
	carry_digit := 0

	for ; i >= 0; i -= 1 {

		currDigit1 := 0

		currDigit1 += int(a[i]) - int('0')

		currDigit1 += int(b[i]) - int('0')

		sum := currDigit1 + carry_digit

		x[i] = byte((sum % base) + int('0'))

		carry_digit = sum / base

	}

	result := ""

	for i := 0; i < len(x); i++ {
		if i == len(r1re[0]) {
			result += "."
		}
		result += string(x[i])
	}

	if carry_digit > 0 {
		return strconv.Itoa(carry_digit) + result
	}

	return result
}
