/**
 * Copyright 2015-2016, Wothing Co., Ltd.
 * All rights reserved.
 *
 * Created by elvizlai on 2016/11/7 12:22.
 */

package main

import (
	"fmt"
	"testing"
)

var s1 = `set x {}`                                                       // empty
var s2 = `set x {"name":"elvizlai","age":18 , "mobile":"86-13800138000"}` // special character
var s3 = `set x {"arr":["a","b"]}`                                        // array
var s4 = `set x {"addr":{"a1":1,"a2":"abc"}}`                             // embed struct
var s5 = `set x ["b","c",123]`

func TestStruct(t *testing.T) {
	structReplacer(&s1)
	structReplacer(&s2)
	structReplacer(&s3)
	structReplacer(&s4)
	structReplacer(&s5)

	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println(s4)
	fmt.Println(s5)
}

var q1 = `""`
var q2 = `"xyz"`
var q3 = `"xyz abc"`

func TestQuote(t *testing.T) {
	quoteReplacer(&q1)
	quoteReplacer(&q2)
	quoteReplacer(&q3)

	fmt.Println(q1)
	fmt.Println(q2)
	fmt.Println(q3)

	fmt.Println(varMap)
}
