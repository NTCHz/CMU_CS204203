package main

import "fmt"

// "bufio"

// "io"
// "log"
// "os"
// "strings"

func main() {

	// var line string
	// var n1, n2 string = "", ""
	// var base int = -1
	// scanner := bufio.NewScanner(os.Stdin)
	// // var result []string
	// for scanner.Scan() {
	// 	line = scanner.Text()

	// 	if strings.HasPrefix(line, "#") || strings.HasPrefix(line, "//") { // skip the line starting with #
	// 		continue
	// 	}

	// 	_, err := fmt.Sscanf(line, "%s %s %d", &n1, &n2, &base)

	// 	if err != nil {
	// 		if err != io.EOF {

	// 			log.Fatal(err)
	// 		}
	// 		break
	// 	}

	// fmt.Println(baseNAddition("11.01", "1.1", 2))   // 100.11
	// fmt.Println(baseNAddition("11.01", "1.0", 2))   // 100.01
	// fmt.Println(baseNAddition("18.50", "10.2", 10)) // 28.70
	fmt.Println(baseNAddition("5.", "10.", 10)) // 15
	fmt.Println(baseNAddition("5", "10", 10))   // 15
	// fmt.Println(baseNAddition("5", ".10", 10))      // 15
	// fmt.Println(baseNAddition(".5", "10", 10))      // 15
	// fmt.Println(baseNAddition("5", "10.", 10))      // 15
	// fmt.Println(baseNAddition(".", "10.", 10))      // 10
	// fmt.Println(baseNAddition(".", "10.", 10))      // 10
	// fmt.Println(baseNAddition(".", ".", 10))        // 0.0
	// fmt.Println(baseNAddition("0.1", "2.2", 3))     // 10.0
	// println(baseNAddition("0", "0", 10))            //0
	// println(baseNAddition(".0", "0.", 10))          //0.0
	// println(baseNAddition("0.", "0", 10))           //0.0
	// println(baseNAddition("0", "0.", 10))           //0.0
	// println(baseNAddition("0.0", "0.0", 10))        //0.0
	// println(baseNAddition("0.0", ".0", 10))         //0.0
	// println(baseNAddition("", "", 10))              //0
	// }
}
