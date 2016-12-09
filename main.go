/**
 * Copyright 2015-2016, Wothing Co., Ltd.
 * All rights reserved.
 *
 * Created by elvizlai on 2016/10/12 16:23.
 */

package main

import (
	"flag"
	"os"
	"strings"
	"time"
)

var passCount = 0
var failCount = 0
var f = "."

func main() {
	flag.BoolVar(&debugMode, "d", debugMode, "debug mode")
	flag.StringVar(&f, "f", f, "test files, if multi, using ';' to separate")
	flag.Parse()

	if os.Getenv("debug") != "" || os.Getenv("DEBUG") != "" {
		debugMode = true
	}

	var start = time.Now()

	args := flag.Args()
	if len(args) != 0 {
		f = strings.Join(args, ";")
	}

	// read then run code
	for _, v := range FileList(f) {
		fileHandler(v)
	}

	if failCount == 0 {
		Infof("\nall %d assert passed, using: %s", passCount, time.Now().Sub(start))
	} else {
		Errorf("\n%d assert passed, %d assert failed, using: %s", passCount, failCount, time.Now().Sub(start))
		os.Exit(2)
	}
}
