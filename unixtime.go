package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

var (
	isVerbose = false
)

func main() {
	var d = flag.String("date", "", "date string")
	var v = flag.Bool("v", false, "verbose")
	flag.Parse()

	isVerbose = *v

	switch {
	case *d != "":
		if *d == "now" {
			printNow()
		} else {
			convertDateString(*d)
		}
	case len(flag.Args()) > 0:
		convertUnixTimestamp(flag.Arg(0))
	default:
		// read from stdin
		readFromStdin()
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

	if isVerbose {
		fmt.Printf("secondStr: %v\n", secondStr)
		fmt.Printf("nanoStr: %v\n", nanoStr)
	}

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

	tm := time.Unix(second, nano)
	fmt.Printf("date: %v\tnano: %v\n", tm, nano)
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

func readFromStdin() {
	buf := make([]byte, 1024)
	for {
		n, err := os.Stdin.Read(buf)
		if err != nil {
			fmt.Printf("read from stdin failed, %s\n", err)
			return
		}

		if n == 0 {
			break
		}

		for i := 0; i < n; i++ {
			if buf[i] == '\n' || buf[i] == '\r' {
				line := string(buf[0:i])
				convertUnixTimestamp(line)
			}
		}
	}
}
