// ทิชานนท์ รตนแสนวัน
// 660510702
// HW03_1
// 204203 Sec 001

package main

import (
	"math"
	"strings"
)

func addNSubtract(n1, n2 string, bitLen uint8) (int64, int64) {
	len1, len2 := len(n1), len(n2)
	ccc := int(bitLen)

	if len1 > ccc {
		n1 = n1[int(math.Abs(float64(ccc-len(n1)))):]
	}
	if len2 > ccc {
		n2 = n2[int(math.Abs(float64(ccc-len(n2)))):]
	}

	if len1 > len2 || len1 == len2 {
		if n2[0] == '0' {
			n2 = strings.Repeat("0", int(math.Abs(float64(ccc-len(n2))))) + n2
		} else {
			n2 = strings.Repeat("1", int(math.Abs(float64(ccc-len(n2))))) + n2
		}
	}

	if len2 > len1 || len1 == len2 {
		if n1[0] == '0' {
			n1 = strings.Repeat("0", int(math.Abs(float64(ccc-len(n1))))) + n1
		} else {
			n1 = strings.Repeat("1", int(math.Abs(float64(ccc-len(n1))))) + n1
		}
	}

	n1 = n1[int(math.Abs(float64(ccc-len(n1)))):]
	n2 = n2[int(math.Abs(float64(ccc-len(n2)))):]

	x1 := make([]byte, ccc)
	x2 := make([]byte, ccc)

	i := ccc - 1

	n2re := additiveInverse(n2)

	carry_digit1 := 0
	carry_digit2 := 0

	for ; i >= 0; i -= 1 {

		currDigit1 := 0
		currDigit2 := 0

		currDigit1 += int(n1[i]) - int('0')
		currDigit2 += int(n1[i]) - int('0')

		currDigit1 += int(n2[i]) - int('0')
		currDigit2 += int(n2re[i]) - int('0')

		sum1 := currDigit1 + carry_digit1
		sum2 := currDigit2 + carry_digit2

		x1[i] = byte((sum1 % 2) + int('0'))
		x2[i] = byte((sum2 % 2) + int('0'))

		carry_digit1 = sum1 / 2
		carry_digit2 = sum2 / 2

	}

	return calculate(x1, bitLen), calculate(x2, bitLen)
}

func calculate(result []byte, bitLen uint8) int64 {
	var sum int64 = 0
	var sign int64 = 1
	if (int64(result[0]) - int64('0')) == 1 {
		sign = -1
	}

	x := len(result) - int(bitLen)
	for i := x; i < len(result); i++ {
		if i == 0 && (int64(result[0])-int64('0')) == 1 {
			sum += ((int64(result[i]) - int64('0')) * int64(math.Pow(2, float64((len(result)-i)-1))) * sign)
			continue
		}
		sum += (int64(result[i]) - int64('0')) * int64(math.Pow(2, float64((len(result)-i)-1)))
	}

	return sum
}

func additiveInverse(x string) string {
	var result string = ""
	var carry int64 = 1

	if strings.Count(x, "0") == len(x) {
		for i := 0; i < len(x); i++ {
			result += "0"
		}
		return result
	}

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

	return result

}
