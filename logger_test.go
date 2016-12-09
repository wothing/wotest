/**
 * Copyright 2015-2016, Wothing Co., Ltd.
 * All rights reserved.
 *
 * Created by elvizlai on 2016/10/14 14:56.
 */

package main

import "testing"

func TestDebug(t *testing.T) {
	Debugf("hidden")
	debugMode = true
	Debugf("debug")
}

func TestInfo(t *testing.T) {
	Infof("info")
}

func TestWarn(t *testing.T) {
	Warnf("warn")
}

func TestError(t *testing.T) {
	Errorf("error")
}
