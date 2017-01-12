/**
 * Copyright 2015-2016, Wothing Co., Ltd.
 * All rights reserved.
 *
 * Created by elvizlai on 2017/1/12 14:23.
 */

package store

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"
)

// TODO a tree may be better

var varMap = map[string]string{}

func Set(key string, value string) {
	varMap[key] = value
}

func Get(key string) (string, bool) {
	value, ok := varMap[key]
	return value, ok
}

// for $body $header $resp
func DeleteWithPrefix(prefix string) {
	for key := range varMap {
		if strings.HasPrefix(key, prefix) {
			delete(varMap, key)
		}
	}
}

func DataWalker(prefix string, x interface{}) {
	data, err := json.Marshal(x)
	if err != nil {
		log.Fatal(err)
	}
	Set(prefix, string(data))

	switch s := x.(type) {
	case map[string]interface{}:
		for l, v := range s {
			DataWalker(prefix+"."+l, v)
		}

	case []interface{}:
		for i, v := range s {
			DataWalker(prefix+"["+strconv.Itoa(i)+"]", v)
		}
	case float64:
		Set(prefix, strconv.FormatFloat(s, 'f', -1, 64))

	default:
		Set(prefix, fmt.Sprint(x)) // force write
	}
}
