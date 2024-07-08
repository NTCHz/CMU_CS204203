// ทิชานนท์ รตนแสนวัน
// 660510702
// HW02_1
// 204203 Sec 001

package main

import (
	"fmt"
	"log"
	"strings"
)

// set the maximum digit length. why 36 and not 35?
const MAX = 36

func main() {
	// why are we using string?
	var n1, n2 string

	// read in two string (can be multiple lines)
	_, err := fmt.Scan(&n1, &n2)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(addition(n1, n2))
}

func addition(n1, n2 string) string {
	// this is just a skeleton and will give out wrong result
	result := []byte(strings.Repeat("x", MAX))
	len1 := len(n1)
	len2 := len(n2)

	// loop from the left most digit
	i, j, k := len1-1, len2-1, MAX-1

	// carry from the previous digit
	carry_digit := 0

	// loop from right most position
	for ; i >= 0 || j >= 0; i, j, k = i-1, j-1, k-1 {
		// current digit
		currDigit := 0

		// add the value from the current digit of n1 when it is not out of bound
		if i >= 0 {
			currDigit += int(n1[i]) - int('0')
		}

		// add the value from the current digit of n2 when it is not out of bound
		if j >= 0 {
			currDigit += int(n2[j]) - int('0')
		}

		sum := currDigit + carry_digit
		// convert back to byte (one char is called byte)
		result[k] = byte((sum % 10) + int('0'))

		// carry to the next digit
		carry_digit = sum / 10

	}

	// if don't have any carry digit, return the result
	if carry_digit > 0 {
		// if have carry digit, add it to the result
		result[k] = byte(carry_digit + int('0'))

		// convert array of bytes to string
		return string(result[k:])

	} else {
		// convert array of bytes to string
		return string(result[k+1:])

	}
}
