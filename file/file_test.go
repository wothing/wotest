/**
 * Copyright 2015-2016, Wothing Co., Ltd.
 * All rights reserved.
 *
 * Created by elvizlai on 2016/10/12 16:45.
 */

package file

import (
	"fmt"
	"testing"
)

func TestList(t *testing.T) {
	fmt.Println(List(".", "go"))
}

func TestWalk(t *testing.T) {
	fmt.Println(walk(".", "go"))
}

func TestRead(t *testing.T) {
	fmt.Println(Read("./file.go"))
}
