/**
 * Copyright 2015-2016, Wothing Co., Ltd.
 * All rights reserved.
 *
 * Created by elvizlai on 2016/10/12 16:37.
 */

package file

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/wothing/wotest/log"
)

const fileSuffix = ".wt"

// List return all files if loc is dir, or it returns self array
func List(loc string, suffix string) []string {
	data := []string{}
	// check if is file or dir
	fi, err := os.Stat(loc)
	if err != nil {
		log.Fatalf(err.Error())
	}
	if fi.IsDir() {
		data = append(data, walk(loc, suffix)...)
	} else {
		data = append(data, loc)
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

// Read return contents of file
func Read(fileName string) ([]byte, error) {
	return ioutil.ReadFile(fileName)
}
