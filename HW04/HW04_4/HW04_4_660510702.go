// ทิชานนท์ รตนแสนวัน
// 660510702
// HW04_4
// 204203 Sec 001
package main

type Triplet struct {
	Offset int
	Length int
	Next   byte
}

func decodeLZ77(matches []Triplet) string {
	decodedString := []byte{}
	for _, match := range matches {
		x := len(decodedString) - match.Offset
		decodedString = append(decodedString, decodedString[x:x+match.Length]...)
		if match.Next != 0 {
			decodedString = append(decodedString, match.Next)
		}
	}
	return string(decodedString)
}
