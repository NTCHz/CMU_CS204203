// ทิชานนท์ รตนแสนวัน
// 660510702
// HW01_1
// Problem A
// 204203 Sec 001

package main

func factorial(num1 int8) int64 {
	var result int64 = 1
	for i := 1; i <= int(num1); i++ {
		result *= int64(i)
	}
	return result
}
