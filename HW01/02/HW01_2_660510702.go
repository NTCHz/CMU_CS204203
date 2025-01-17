// ทิชานนท์ รตนแสนวัน
// 660510702
// HW01_2
// Problem A
// 204203 Sec 001

// THIS TEMPLATE IS OPTIONAL
// YOU ARE NOT REQUIRED TO USE IT

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	DEBUG = false
)

func checkPair(w1, w2 string) bool {
	// ความยาวของคำทั้งสอง
	len1, len2 := len(w1), len(w2)

	// กรณีที่ความยาวของคำต่างกันมากกว่า 1 ตัวอักษร
	if len1-len2 > 1 || len2-len1 > 1 {
		return false
	}

	// กรณีที่ทั้งสองคำเหมือนกัน
	if strings.Compare(w1, w2) == 0 {
		return false
	}

	// กรณีที่ความยาวของคำเท่ากัน
	if len1 == len2 {
		diff := 0
		for i := 0; i < len1; i++ {
			if w1[i] != w2[i] {
				diff++
			}
			if diff > 1 {
				return false
			}
		}
	}

	// กรณีที่ความยาวของ w1 มากกว่า w2
	if len1 > len2 {
		diff := 0
		for i, j := 0, 0; i < len1 && j < len2; {
			if w1[i] != w2[j] {
				diff++
				i++
			} else {
				i++
				j++
			}
			if diff > 1 {
				return false
			}
		}
	}

	// กรณีที่ความยาวของ w1 น้อยกว่า w2
	if len1 < len2 {
		diff := 0
		for i, j := 0, 0; i < len1 && j < len2; {
			if w1[i] != w2[j] {
				diff++
				j++
			} else {
				i++
				j++
			}
			if diff > 1 {
				return false
			}
		}
	}

	return true
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	numCases := scanToInt(scanner.Text())

	for i := 0; i < numCases; i++ {
		scanner.Scan()
		numWords := scanToInt(scanner.Text())
		if DEBUG {
			fmt.Println(numWords)
		}
		wordArray := make([]string, numWords)
		for j := 0; j < numWords; j++ {
			scanner.Scan()
			wordArray[j] = scanner.Text()
			if DEBUG {
				fmt.Println(wordArray[j])
			}
		}

		j := 1

		for ; j < numWords; j++ {
			if !checkPair(wordArray[j-1], wordArray[j]) && !checkPair(wordArray[j], wordArray[j-1]) {
				fmt.Println("F")
				break
			}
		}

		if j == numWords {
			fmt.Println("T")
		}
	}
}

func scanToInt(s string) int {
	var n int
	fmt.Sscan(s, &n)
	return n
}
