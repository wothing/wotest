/**
 * Copyright 2015-2016, Wothing Co., Ltd.
 * All rights reserved.
 *
 * Created by elvizlai on 2016/10/13 12:45.
 */

package main

import (
	"fmt"
	"testing"
)

var s1 = `{} {"name":"elvizlai"}`
var s2 = `{"name":"elvizlai", "age":18}`
var s3 = `{"name":"elvizlai","age":18,"add":["a","b"]}`
var s4 = `"{"name":"elvizlai", "age":18}"`

func TestStcRgx(t *testing.T) {
	fmt.Println(stcRgx.FindAllString(s1, -1))
	fmt.Println(stcRgx.FindAllString(s2, -1))
	fmt.Println(stcRgx.FindAllString(s3, -1))
	fmt.Println(stcRgx.FindAllString(s4, -1))
}

func TestStcReplace(t *testing.T) {
	stcReplacer(&s1)
	fmt.Println(s1)

	stcReplacer(&s3)
	fmt.Println(s3)
}

var q1 = `"xyz "`
var q2 = `"xyz 123"`
var q3 = `"xyz 123" "abc"`
var q4 = `"Authorization"  "Bearer asdasdasd"`

func TestQuoRgx(t *testing.T) {
	fmt.Println(quoRgx.FindAllString(q1, -1))
	fmt.Println(quoRgx.FindAllString(q2, -1))
	fmt.Println(quoRgx.FindAllString(q3, -1))
	fmt.Println(quoRgx.FindAllString(q4, -1))
}

func TestQuoReplacer(t *testing.T) {
	quoReplacer(&q1)
	fmt.Println(q1)

	quoReplacer(&q3)
	fmt.Println(q3)

	quoReplacer(&q4)
	fmt.Println(q4)
}

var c1 = `match "xyz" "abc"`
var c2 = `match {"name":"elvizlai"}  "xyz"`
var c3 = `match " " " "`

func TestCombination(t *testing.T) {
	replacer(&c1)
	fmt.Println(c1)

	replacer(&c2)
	fmt.Println(c2)

	replacer(&c3)
	fmt.Println(c3)

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
