/**
 * Copyright 2015-2016, Wothing Co., Ltd.
 * All rights reserved.
 *
 * Created by elvizlai on 2016/12/9 16:10.
 */

package main

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/wothing/wotest/log"
	"io"
	"io/ioutil"
	"strings"
)

var visitMap = make(map[string]bool)

func fileHandler(fileName string) {
	if visitMap[fileName] {
		return
	} else {
		visitMap[fileName] = true
	}

	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatalf(err.Error())
	}

	var prevEmpty = true

	reader := bufio.NewReader(bytes.NewBuffer(data))
	rawStr := ""
	for {
		data, _, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				log.Fatalf(err.Error())
			}
		}

		rawStr = rawStr + string(data)

		// line breaker
		if strings.HasSuffix(rawStr, ` \r`) {
			rawStr = strings.TrimSuffix(rawStr, ` \r`)
			continue
		}

		// trim space
		rawStr = strings.TrimSpace(rawStr)

		// empty line
		if rawStr == "" {
			if !prevEmpty {
				fmt.Println()
			}
			prevEmpty = true
			continue
		}

		// comment line
		if strings.HasPrefix(rawStr, `//`) {
			rawStr = ""
			continue
		}

		prevEmpty = false

		// load prefix
		if strings.HasPrefix(rawStr, `load `) {
			fileHandler(strings.TrimSpace(rawStr[5:]))
			rawStr = ""
			continue
		}

		eval(&rawStr)
		rawStr = ""
	}
}
