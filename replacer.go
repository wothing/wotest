/**
 * Copyright 2015-2016, Wothing Co., Ltd.
 * All rights reserved.
 *
 * Created by elvizlai on 2016/11/7 12:21.
 */

package main

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/wothing/wotest/log"
	"github.com/wothing/wotest/store"
)

//var varMap = map[string]string{}

func replacer(s *string) {
	structReplacer(s)
	quoteReplacer(s)
	spaceReplacer(s)
}

// set struct to varMap
// STRUCT{}
// ARRAY[]
var structRgx = regexp.MustCompile(`\s{1,}?[\{|\[].+[\}|\]]$`)

func structReplacer(s *string) {
	ss := structRgx.FindAllString(*s, -1)
	for i, v := range ss {
		store.Set("$.stc."+strconv.Itoa(i), strings.TrimSpace(v))
		*s = strings.Replace(*s, v, " $.stc."+strconv.Itoa(i), 1)
	}
}

// set quote to varMap
// " "
var quoteRgx = regexp.MustCompile(`"(.*)"`)

func quoteReplacer(s *string) {
	qs := quoteRgx.FindAllString(*s, -1)
	for i, v := range qs {
		store.Set("$.quo."+strconv.Itoa(i), strings.Trim(v, `"`))
		*s = strings.Replace(*s, v, "$.quo."+strconv.Itoa(i), 1)
	}
}

// trim space to one
var spaceRgx = regexp.MustCompile(`\s{2,}`)

func spaceReplacer(s *string) {
	*s = spaceRgx.ReplaceAllString(*s, " ")
}

// parse $xxx
var varRgx = regexp.MustCompile(`\$[\w\.\[\]\-]+`)

// TODO not found warn
func varReplacer(s *string) {
	for _, v := range varRgx.FindAllString(*s, -1) {
		if k, ok := store.Get(v); ok {
			*s = strings.Replace(*s, v, k, 1)
			varReplacer(s) // TODO is this ok?
		} else {
			log.Warnf("'%s' NOT EXIST", v)
		}
	}
}

// parse `xyz xyz`->xyz xyz;
// `NOT_EMPTY` should work
var cmdRgx = regexp.MustCompile("`.+`")
