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
	"os"
	"time"
)

var f = "."

func main() {
	flag.BoolVar(&debugMode, "d", debugMode, "debug mode")
	flag.StringVar(&f, "f", f, "test files, if multi, using ';' to separate")
	flag.Parse()

	var start = time.Now()
	// read then run code
	for _, v := range FileList(f) {
		data, err := ioutil.ReadFile(v)
		if err != nil {
			panic(err)
		}
		reader := bufio.NewReader(bytes.NewBuffer(data))
		for {
			data, _, err := reader.ReadLine()
			if err != nil {
				if err == io.EOF {
					break
				} else {
					panic(err)
				}
			}
			temp := string(data)
			eval(&temp)
		}
	}

	if failCount == 0 {
		Info("\nall %d assert passed, using: %s", passCount, time.Now().Sub(start))
	} else {
		Error("\n%d assert failed", failCount)
		os.Exit(2)
	}

}
