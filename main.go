/**
 * Copyright 2015-2016, Wothing Co., Ltd.
 * All rights reserved.
 *
 * Created by elvizlai on 2016/10/12 16:23.
 */

package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"
)

var f = "."

func main() {
	flag.BoolVar(&debugMode, "d", debugMode, "debug mode")
	flag.StringVar(&f, "f", f, "test files, if multi, using ';' to separate")
	flag.Parse()

	if os.Getenv("debug") != "" || os.Getenv("DEBUG") != "" {
		debugMode = true
	}

	var start = time.Now()
	// read then run code
	for _, v := range FileList(f) {
		fileHandler(v)
	}

	if failCount == 0 {
		Info("\nall %d assert passed, using: %s", passCount, time.Now().Sub(start))
	} else {
		Error("\n%d assert failed", failCount)
		os.Exit(2)
	}
}

var visitMap = make(map[string]bool)

func fileHandler(fileName string) {
	if visitMap[fileName] {
		return
	} else {
		visitMap[fileName] = true
	}

	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}

	reader := bufio.NewReader(bytes.NewBuffer(data))
	rawStr := ""
	for {
		data, _, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				log.Fatal(err)
			}
		}

		rawStr = rawStr + string(data)

		// line breaker
		if strings.HasSuffix(rawStr, ` \r`) {
			rawStr = strings.TrimSuffix(rawStr, ` \r`)
			continue
		}

		// trim space
		rawStr = strings.TrimSpace(rawStr)

		// empty line
		if rawStr == "" {
			fmt.Println()
			continue
		}

		// comment line
		if strings.HasPrefix(rawStr, `//`) {
			rawStr = ""
			continue
		}

		// load prefix
		if strings.HasPrefix(rawStr, `load `) {
			fileHandler(strings.TrimSpace(rawStr[5:]))
			rawStr = ""
			continue
		}

		eval(&rawStr)
		rawStr = ""
	}
}
