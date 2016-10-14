/**
 * Copyright 2015-2016, Wothing Co., Ltd.
 * All rights reserved.
 *
 * Created by elvizlai on 2016/10/13 11:49.
 */

package main

import "fmt"

func Debug(s string, para ...interface{}) {
	s = "\033[0;33m" + s + "\033[0m\n"
	fmt.Printf(s, para...)
}

func Info(s string, para ...interface{}) {
	s = "\033[0;32m" + s + "\033[0m\n"
	fmt.Printf(s, para...)
}

func Warn(s string, para ...interface{}) {
	s = "\033[0;35m" + s + "\033[0m\n"
	fmt.Printf(s, para...)
}

func Error(s string, para ...interface{}) {
	s = "\033[0;31m" + s + "\033[0m\n"
	fmt.Printf(s, para...)
}
