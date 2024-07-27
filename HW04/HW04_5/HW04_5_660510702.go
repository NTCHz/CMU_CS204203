// ทิชานนท์ รตนแสนวัน
// 660510702
// HW04_5
// 204203 Sec 001

package main

type Triplet struct {
	Offset int
	Length int
	Next   byte
}

const DEBUG = true

func encodeLZ77(windowSize int, bufferSize int, inputString string) []Triplet {
	var result []Triplet
	var inputStringarr = []byte(inputString)
	for i := 0; i < len(inputStringarr); i++ {
		var offset int
		var length int
		var next byte
		var buffer int
		var searchStart int

		if i-windowSize < 0 {
			searchStart = 0
		} else {
			searchStart = i - windowSize
		}

		if i+bufferSize > len(inputStringarr) {
			buffer = len(inputStringarr)
		} else {
			buffer = i + bufferSize
		}

		for j := searchStart; j < i; j++ {
			X := 0
			for k := 0; i+k < buffer && j+k < i && inputStringarr[j+k] == inputStringarr[i+k]; k++ {
				X++
			}
			if X >= length {
				length = X
				offset = i - j
				if i+length < len(inputStringarr) {
					next = inputStringarr[i+length]
				} else {
					next = 0
				}
			}
		}
		if length > 0 {
			result = append(result, Triplet{offset, length, next})
			i += length
		} else {
			result = append(result, Triplet{0, 0, inputStringarr[i]})
		}
	}

	return result
}
