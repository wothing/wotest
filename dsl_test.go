/**
 * Copyright 2015-2016, Wothing Co., Ltd.
 * All rights reserved.
 *
 * Created by elvizlai on 2016/10/13 12:45.
 */

package main

import (
	"fmt"
	"strings"
	"testing"
)

var s1 = `{}`                                                       // empty
var s2 = `{"name":"elvizlai","age":18 , "mobile":"86-13800138000"}` // special character
var s3 = `{"arr":["a","b"]}`                                        // array
var s4 = `{"addr":{"a1":1,"a2":"abc"}}`                             // embed struct
var s5 = `"{"name":"elvizlai"}"`                                    // quot wrap

func TestStcRgx(t *testing.T) {
	s1r := stcRgx.FindAllString(s1, -1)
	if len(s1r) != 1 && s1r[0] != s1 {
		t.Fatal("s1r")
	}

	s2r := stcRgx.FindAllString(s2, -1)
	if len(s2r) != 1 && s2r[0] != s2 {
		t.Fatal("s2r")
	}

	s3r := stcRgx.FindAllString(s3, -1)
	if len(s3r) != 1 && s3r[0] != s2 {
		t.Fatal("s3r")
	}

	s4r := stcRgx.FindAllString(s4, -1)
	if len(s4r) != 1 && s4r[0] != s2 {
		t.Fatal("s4r")
	}

	s5r := stcRgx.FindAllString(s5, -1)
	if len(s5r) != 1 && s5r[0] != strings.Trim(s5, `"`) {
		t.Fatal("s5r")
	}
}

var q1 = `"xyz "`
var q2 = `"xyz 123" "abc"`
var q3 = `"Authorization"  "Bearer asdasdasd"`

func TestQuoRgx(t *testing.T) {
	q1r := quoRgx.FindAllString(q1, -1)
	if len(q1r) != 1 && q1r[0] != q1 {
		t.Fatal("q1r")
	}

	q2r := quoRgx.FindAllString(q2, -1)
	if len(q2r) != 2 && q2r[0] != "xyz 123" && q2r[1] != "abc" {
		t.Fatal("q2r")
	}

	q3r := quoRgx.FindAllString(q3, -1)
	if len(q3r) != 2 && q3r[0] != "Authorization" && q3r[1] != "Bearer asdasdasd" {
		t.Fatal("q3r")
	}

}

var c1 = `match "xyz" "abc"`
var c2 = `match {"name":"elvizlai"}  "xyz"`
var c3 = `match " " " "`

func TestCombination(t *testing.T) {
	replacer(&c1)
	if c1 != "match $.quo.0 $.quo.1" {
		t.Fatal("c1")
	}

	replacer(&c2)
	if c2 != "match $.stc.0 $.quo.0" {
		t.Fatal("c2")
	}

	replacer(&c3)
	if c3 != "match $.quo.0 $.quo.1" {
		t.Fatal("c3")
	}
}

var v1 = `$.quo.0`
var v2 = `$body_xyz`
var v3 = `$_xyz $body.data[1].goods`

func TestVarRgx(t *testing.T) {
	fmt.Println(varRgx.FindAllString(v1, -1))
	fmt.Println(varRgx.FindAllString(v2, -1))
	fmt.Println(varRgx.FindAllString(v3, -1))
}

func TestEval(t *testing.T) {
	//x1 := "set X 123"
	//eval(&x1)
	//x2 := "echo $X"
	//eval(&x2)
	//x3 := "match `env GOPATH` `env GOPATH`"
	//eval(&x3)
	//x4 := "match 1 2"
	//eval(&x4)
	//x5 := "get https://www.baidu.com"
	//eval(&x5)
	//x6 := `header auth xyz123`
	//eval(&x6)
	//x8 := "ret "
	//eval(&x8)
	//x9 := "ret 200"
	//eval(&x9)
	//x10 := "echo $resp"
	//eval(&x10)
	//x11 := "echo $resp.body"
	//eval(&x11)
}
