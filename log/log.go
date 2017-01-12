/**
 * Copyright 2015-2016, Wothing Co., Ltd.
 * All rights reserved.
 *
 * Created by elvizlai on 2017/1/12 14:55.
 */

package log

import (
	"log"
	"os"
)

var debug = false
var force = false

var std = log.New(os.Stdout, "", 0)

func EnableDebugMode() {
	debug = true
}

func EnableForceMode() {
	force = true
}

func Debugf(format string, para ...interface{}) {
	if debug {
		std.Printf("\033[33m"+format+"\033[0m", para...)
	}
}

func Infof(format string, para ...interface{}) {
	std.Printf("\033[32m"+format+"\033[0m", para...)
}

func Warnf(format string, para ...interface{}) {
	std.Printf("\033[35m"+format+"\033[0m", para...)
}

func Errorf(format string, para ...interface{}) {
	if force {
		std.Fatalf("\033[31m"+format+"\033[0m", para...)
	} else {
		std.Printf("\033[31m"+format+"\033[0m", para...)
	}
}

func Fatalf(format string, para ...interface{}) {
	std.Fatalf("\033[31m"+format+"\033[0m", para...)
}
