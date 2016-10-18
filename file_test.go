/**
 * Copyright 2015-2016, Wothing Co., Ltd.
 * All rights reserved.
 *
 * Created by elvizlai on 2016/10/12 16:45.
 */

package main

import (
	"fmt"
	"testing"
)

func TestFileList(t *testing.T) {
	// fmt.Println(FileList("abc"))
	// fmt.Println(FileList("abc;123"))
}

func TestWalk(t *testing.T) {
	fmt.Println(walk("./", fileSuffix))
}
