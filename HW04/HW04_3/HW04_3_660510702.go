// ทิชานนท์ รตนแสนวัน
// 660510702
// HW04_3
// 204203 Sec 001

package main

import (
	"strings"
)

func decodeHuffman(encodedText string, codingTable map[string]string) string {
	var decoded strings.Builder

	if len(encodedText) == 0 {
		return ""
	}

	// Decode the encoded text
	var code strings.Builder
	for _, bit := range encodedText {
		code.WriteRune(bit)
		if key, found := codingTable[code.String()]; found {
			decoded.WriteString(key)
			code.Reset()
		}
	}

	return decoded.String()
}
