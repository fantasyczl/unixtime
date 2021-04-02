package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func main() {
	var d = flag.String("date", "", "date string")
	flag.Parse()

	if *d != "" {
		convertDateString(*d)
	} else {
		// 处理时间戳
		if len(flag.Args()) == 0 {
			fmt.Printf("lack timestamp")
			printNow()
			return
		}

		convertUnixTimestamp(flag.Arg(0))
	}
}

func convertUnixTimestamp(str string) {
	const NANO = 1e9

	var secondStr, nanoStr string

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

func convertDateString(dateStr string) {
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Printf("LoadLocation failed, %s\n", err)
		return
	}

	t, err := time.ParseInLocation("2006-01-02 15:04:05", dateStr, loc)
	if err != nil {
		fmt.Printf("parse time failed, %s\n", err)
		return
	}

	fmt.Printf("time:\n\t%s\n", t)
	fmt.Printf("unix:\n\t%d\n", t.Unix())
}

func printNow() {
	ts := time.Now()
	fmt.Printf("\nnow:\n%s\nts: %d\n", ts, ts.Unix())
}
