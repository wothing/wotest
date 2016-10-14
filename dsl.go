/**
 * Copyright 2015-2016, Wothing Co., Ltd.
 * All rights reserved.
 *
 * Created by elvizlai on 2016/10/12 17:37.
 */

package main

import (
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var passCount = 0
var failCount = 0

const pass = "\033[0;32m ✔ \033[0m"
const failed = "\033[0;31m ✘ \033[0m"

var varMap = map[string]string{}

var funcMap = map[string]func(*string) error{
	"env": func(s *string) error {
		params := strings.Split(*s, " ")
		if len(params) > 1 {
			return errors.New("env need 1 param")
		}
		*s = os.Getenv(params[0])
		return nil
	},
	"echo": func(s *string) error {
		params := strings.Split(*s, " ")
		if len(params) > 1 {
			return errors.New("echo need 1 param")
		}
		varReplacer(&params[0])
		Debug(params[0])
		return nil
	},
	"set": func(s *string) error {
		params := strings.Split(*s, " ")
		if len(params) != 2 {
			return errors.New("set need 2 param")
		}

		varReplacer(&params[0])
		varReplacer(&params[1])

		varMap["$"+params[0]] = strings.TrimSpace(params[1])

		return nil
	},
	"get": func(s *string) error {
		params := strings.Split(*s, " ")
		if len(params) != 1 {
			return errors.New("get need 1 param")
		}

		varReplacer(&params[0])

		fmt.Printf("[get] '%s'\n", params[0])
		httptest.method = "GET"
		httptest.url = params[0]
		httptest.done = false
		return nil
	},
	"post": func(s *string) error {
		params := strings.Split(*s, " ")
		if len(params) != 1 {
			return errors.New("post need 1 param")
		}

		varReplacer(&params[0])

		fmt.Printf("[pos] '%s'\n", params[0])
		httptest.method = "POST"
		httptest.url = params[0]
		httptest.done = false
		return nil
	},
	"header": func(s *string) error {
		params := strings.Split(*s, " ")
		if len(params) != 2 {
			return errors.New("header need 2 param")
		}

		varReplacer(&params[0])
		varReplacer(&params[1])

		fmt.Printf("[hed] '%s' '%s'\n", params[0], params[1])
		httptest.header[strings.TrimSpace(params[0])] = strings.TrimSpace(params[1])
		return nil
	},
	"json": func(s *string) error {
		params := strings.Split(*s, " ")
		if len(params) > 1 {
			return errors.New("json need 1 param")
		}

		varReplacer(&params[0])
		fmt.Printf("[jsn] '%s'\n", params[0])
		httptest.reqBody = params[0]
		return nil
	},
	"ret": func(s *string) error {
		params := strings.Split(*s, " ")
		if len(params) > 1 {
			return errors.New("ret need 0 or 1 param")
		}

		httptest.do()

		fmt.Print("[ret] ")
		if params[0] == "" {
			fmt.Print(pass, "\n")
		} else {
			varReplacer(&params[0])
			fmt.Printf("'%s' '%s'", params[0], strconv.Itoa(httptest.resp.StatusCode))
			if strconv.Itoa(httptest.resp.StatusCode) == params[0] {
				passCount++
				fmt.Print(pass, "\n")
			} else {
				failCount++
				fmt.Print(failed, "\n")
			}
		}

		return nil
	},
	"match": func(s *string) error {
		params := strings.Split(*s, " ")
		if len(params) != 2 {
			return errors.New("match need 2 param")
		}

		varReplacer(&params[0])
		varReplacer(&params[1])

		fmt.Printf("[mat] '%s' '%s'", params[0], params[1])
		if params[0] == params[1] {
			passCount++
			fmt.Println(pass)
		} else {
			failCount++
			fmt.Println(failed)
		}
		return nil
	},
	"contains": func(s *string) error {
		params := strings.Split(*s, " ")
		if len(params) != 2 {
			return errors.New("contains need 2 param")
		}

		varReplacer(&params[0])
		varReplacer(&params[1])

		fmt.Printf("[cot] '%s' '%s'", params[0], params[1])
		if strings.Contains(params[0], params[1]) {
			passCount++
			fmt.Print(pass, "\n")
		} else {
			failCount++
			fmt.Print(failed, "\n")
		}

		return nil
	},
	// TODO
	"regexp": func(s *string) error {
		params := strings.Split(*s, " ")
		if len(params) != 2 {
			return errors.New("regexp need 2 param")
		}

		varReplacer(&params[0])
		//varReplacer(&params[1])

		return nil
	},
}

// parse `xyz xyz`->xyz xyz
var cmdRgx = regexp.MustCompile("\\`[\\w\\s]+\\`")

func eval(s *string) {
	*s = strings.TrimSpace(*s)

	if *s == "" {
		fmt.Println()
		return
	}

	if strings.HasPrefix(*s, "//") {
		return
	}

	replacer(s)

	//fmt.Println(*s)
	//fmt.Println(varMap)

	cmds := cmdRgx.FindAllString(*s, -1)
	for i, v := range cmds {
		temp := v
		v = strings.Trim(temp, "`")
		eval(&v)
		varMap["$.temp_"+strconv.Itoa(i)] = v
		*s = strings.Replace(*s, temp, "$.temp_"+strconv.Itoa(i), 1)
	}

	x := strings.Split(*s, " ")

	if f, ok := funcMap[x[0]]; ok {
		suffix := strings.TrimSpace(strings.TrimPrefix(*s, x[0]))
		//varReplace(&suffix)
		err := f(&suffix)
		if err != nil {
			Error(err.Error())
			os.Exit(1)
		}
		*s = suffix
	}
}

func replacer(s *string) {
	stcReplacer(s)
	quoReplacer(s)
	spcReplacer(s)
}

// set struct to varMap
var stcRgx = regexp.MustCompile(`\{[\[\]":,\w\s]*\}`)

func stcReplacer(s *string) {
	stcs := stcRgx.FindAllString(*s, -1)

	for i, v := range stcs {
		varMap["$.stc."+strconv.Itoa(i)] = v
		*s = strings.Replace(*s, v, "$.stc."+strconv.Itoa(i), 1)
	}
}

// set "xyz abc" to varMap
var quoRgx = regexp.MustCompile(`".+?"`)

func quoReplacer(s *string) {
	quos := quoRgx.FindAllString(*s, -1)

	for i, v := range quos {
		varMap["$.quo."+strconv.Itoa(i)] = strings.Trim(v, `"`)
		*s = strings.Replace(*s, v, "$.quo."+strconv.Itoa(i), 1)
	}
}

// trim space to one
var spcRgx = regexp.MustCompile(`\s{2,}`)

func spcReplacer(s *string) {
	*s = spcRgx.ReplaceAllString(*s, " ")
}

// parse $xxx
var varRgx = regexp.MustCompile("\\$[\\w\\.\\[\\]\\-]+")

// TODO not found warn
func varReplacer(s *string) {
	for _, v := range varRgx.FindAllString(*s, -1) {
		if k, ok := varMap[v]; ok {
			*s = strings.Replace(*s, v, k, 1)
		} else {
			Warn("'%s' NOT EXIST", v)
		}
	}
}
