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
	"runtime"
	"strings"
	"time"

	"github.com/wothing/wotest/file"
	"github.com/wothing/wotest/log"
)

const wtSuffix = ".wt"

var passCount = 0
var failCount = 0
var files = "."

var debugMode = false
var forceMode = false

var ver = "1.0"

func init() {
	var version = false
	flag.BoolVar(&debugMode, "d", debugMode, "debug mode, echo enable")
	flag.BoolVar(&forceMode, "f", forceMode, "force mode, stop test if err")
	flag.StringVar(&files, "file", files, "wotest file or dir, if multiple, using ';' to separate")
	flag.BoolVar(&version, "version", version, "show version info")
	flag.Parse()

	if version {
		log.Infof("ver%s, using %s", ver, runtime.Version())
		os.Exit(0)
	}

	if debugMode {
		log.EnableDebugMode()
	}

	if forceMode {
		log.EnableForceMode()
	}
}

func main() {
	var start = time.Now()

	if args := flag.Args(); len(args) != 0 {
		files = strings.Join(args, ";")
	}

	// read then run code
	for _, v := range strings.Split(files, ";") {
		for _, v := range file.List(v, wtSuffix) {
			fileHandler(v)
		}
	}

	if failCount == 0 {
		log.Infof("\nall %d assert passed, using: %s", passCount, time.Now().Sub(start))
	} else {
		log.Fatalf("\n%d assert passed, %d assert failed, using: %s", passCount, failCount, time.Now().Sub(start))
	}
}
