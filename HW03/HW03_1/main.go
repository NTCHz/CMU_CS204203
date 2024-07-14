package main

func main() {

	// var n1, n2 string
	// var bitLen uint8

	// _, err := fmt.Scan(&n1, &n2, &bitLen)

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(addNSubtract(n1, n2, bitLen))
	println(addNSubtract("11101", "1", 4)) // -4 -2

	println(addNSubtract("01101", "01000", 4)) // 5 5

	println(addNSubtract("0", "0", 4)) //0 0

	println(addNSubtract("01101", "01000", 6)) //21 5

	println(addNSubtract("01101", "01000", 5)) //-11 5
}
