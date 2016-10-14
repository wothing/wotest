/**
 * Copyright 2015-2016, Wothing Co., Ltd.
 * All rights reserved.
 *
 * Created by elvizlai on 2016/10/12 16:37.
 */

package main

import (
	"io/ioutil"
	"path/filepath"
	"strings"
)

const suffix = ".wt"

// FileList parse loc
func FileList(loc string) []string {
	fs := strings.Split(loc, ";")

	data := []string{}
	for i := range fs {
		if strings.HasSuffix(fs[i], suffix) {
			data = append(data, fs[i])
		} else {
			data = append(data, walk(fs[i], suffix)...)
		}
	}

	return data
}

func walk(f string, suffix string) []string {
	fis, err := ioutil.ReadDir(f)
	if err != nil {
		panic(err)
	}

	list := []string{}

	for _, v := range fis {
		if !v.IsDir() && strings.HasSuffix(v.Name(), suffix) {
			list = append(list, filepath.Join(f, v.Name()))
		}
	}

	return list
}
