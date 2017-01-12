/**
 * Copyright 2015-2016, Wothing Co., Ltd.
 * All rights reserved.
 *
 * Created by elvizlai on 2016/10/12 16:23.
 */

package main

import (
	"flag"
	"strings"
	"time"

	"github.com/wothing/wotest/log"
)

var passCount = 0
var failCount = 0
var f = "."

var debugMode = false
var forceMode = false

func main() {
	flag.BoolVar(&debugMode, "d", debugMode, "debug mode, echo enable")
	flag.BoolVar(&forceMode, "f", forceMode, "force mode, stop test if err")
	flag.StringVar(&f, "file", f, "test files, if multi, using ';' to separate")
	flag.Parse()

	if debugMode {
		log.EnableDebugMode()
	}

	if forceMode {
		log.EnableForceMode()
	}

	var start = time.Now()

	if args := flag.Args(); len(args) != 0 {
		f = strings.Join(args, ";")
	}

	// read then run code
	for _, v := range FileList(f) {
		fileHandler(v)
	}

	if failCount == 0 {
		log.Infof("\nall %d assert passed, using: %s", passCount, time.Now().Sub(start))
	} else {
		log.Fatalf("\n%d assert passed, %d assert failed, using: %s", passCount, failCount, time.Now().Sub(start))
	}
}
