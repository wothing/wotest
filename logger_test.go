/**
 * Copyright 2015-2016, Wothing Co., Ltd.
 * All rights reserved.
 *
 * Created by elvizlai on 2016/10/14 14:56.
 */

package main

import "testing"

func TestDebug(t *testing.T) {
	Debug("debug")
}

func TestInfo(t *testing.T) {
	Info("info")
}

func TestWarn(t *testing.T) {
	Warn("warn")
}

func TestError(t *testing.T) {
	Error("error")
}
