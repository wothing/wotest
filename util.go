/**
 * Copyright 2015-2016, Wothing Co., Ltd.
 * All rights reserved.
 *
 * Created by elvizlai on 2016/11/7 09:33.
 */

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
)

// TODO may add more type assert
func dataWalker(prefix string, x interface{}) {
	data, err := json.Marshal(x)
	if err != nil {
		log.Fatal(err)
	}
	varMap[prefix] = string(data)

	switch s := x.(type) {
	case map[string]interface{}:
		for l, v := range s {
			dataWalker(prefix+"."+l, v)
		}

	case []interface{}:
		for i, v := range s {
			dataWalker(prefix+"["+strconv.Itoa(i)+"]", v)
		}
	case float64:
		varMap[prefix] = strconv.FormatFloat(s, 'f', -1, 64)

	default:
		varMap[prefix] = fmt.Sprint(x) // force write
	}
}
