/**
 * Copyright 2015-2016, Wothing Co., Ltd.
 * All rights reserved.
 *
 * Created by elvizlai on 2016/10/12 16:37.
 */

package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/wothing/wotest/log"
)

const fileSuffix = ".wt"

// FileList parse loc
func FileList(loc string) []string {
	fs := strings.Split(loc, ";")

	data := []string{}
	for i := range fs {
		// check if is file or dir
		fi, err := os.Stat(fs[i])
		if err != nil {
			log.Fatalf(err.Error())
		}
		if fi.IsDir() {
			data = append(data, walk(fs[i], fileSuffix)...)
		} else {
			data = append(data, fs[i])
		}
	}

	return data
}

func walk(f string, suffix string) []string {
	fis, err := ioutil.ReadDir(f)
	if err != nil {
		log.Fatalf(err.Error())
	}

	list := []string{}

	for _, v := range fis {
		if !v.IsDir() && strings.HasSuffix(v.Name(), suffix) {
			list = append(list, filepath.Join(f, v.Name()))
		}
	}

	return list
}
