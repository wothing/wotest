/**
 * Copyright 2015-2016, Wothing Co., Ltd.
 * All rights reserved.
 *
 * Created by elvizlai on 2016/10/13 11:49.
 */

package main

import (
	"fmt"
	"os"
)

var debugMode = false

func Debugf(s string, para ...interface{}) {
	if debugMode {
		s = "\033[0;33m" + s + "\033[0m\n"
		fmt.Printf(s, para...)
	}
}

func Infof(s string, para ...interface{}) {
	s = "\033[0;32m" + s + "\033[0m\n"
	fmt.Printf(s, para...)
}

func Warnf(s string, para ...interface{}) {
	s = "\033[0;35m" + s + "\033[0m\n"
	fmt.Printf(s, para...)
}

func Errorf(s string, para ...interface{}) {
	s = "\033[0;31m" + s + "\033[0m\n"
	fmt.Printf(s, para...)
}

func Fatal(e error) {
	err := "\033[0;31m" + e.Error() + "\033[0m\n"
	fmt.Print(err)
	os.Exit(2)
}
