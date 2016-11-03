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
	str := ""
	for {
		data, _, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				log.Fatal(err)
			}
		}
		str = str + strings.TrimSpace(string(data)) // trim space

		// judge if line breaker
		if strings.Contains(str, ` \r`) {
			str = strings.TrimSpace(strings.TrimSuffix(str, ` \r`))
			continue
		}

		if strings.HasPrefix(str, "load ") {
			fileHandler(strings.TrimSpace(str[4:]))
			str = ""
			continue
		}

		eval(&str)
		str = ""
	}
}
