package main

import "os"
import "fmt"
import "time"
import "strconv"
import "strings"

func main() {
	const NANO = 1e9

	if len(os.Args) < 2 {
		fmt.Println("lack arguemnt")
		os.Exit(1)
	}

	var str, secondStr, nanoStr string
	str = os.Args[1]

	if strings.Contains(str, ".") {
		str = strings.ReplaceAll(str, ".", "")
	}

	secondStr = str

	if len(str) > 10 {
		secondStr = str[0:10]
		nanoStr = str[10:]
	}

	fmt.Printf("secondStr: %v\n", secondStr)
	fmt.Printf("nanoStr: %v\n", nanoStr)

	second, err := strconv.ParseInt(secondStr, 10, 64)

	if err != nil {
		panic(err)
	}

	var nano int64 = 0
	if nanoStr != "" {
		nanoFloat, err := strconv.ParseFloat("0."+nanoStr, 64)
		if err != nil {
			panic(err)
		}
		nano = int64(nanoFloat * NANO)
	}

	fmt.Printf("nano: %v\n", nano)

	tm := time.Unix(second, nano)
	fmt.Printf("date: %v\n", tm)
}
