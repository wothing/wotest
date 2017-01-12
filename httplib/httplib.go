/**
 * Copyright 2015-2016, Wothing Co., Ltd.
 * All rights reserved.
 *
 * Created by elvizlai on 2017/1/12 12:13.
 */

package httplib

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/wothing/wotest/store"
)

var client = &http.Client{}

var done bool
var buff *bytes.Buffer
var req *http.Request
var resp *http.Response

func NewRequest(method string, url string) error {
	done = false
	buff = bytes.NewBuffer(nil)

	var err error
	req, err = http.NewRequest(method, url, nil)
	return err
}

func WithContent(value string) {
	req.Header.Set("content-type", value)
}

func WithHeader(key string, value string) {
	req.Header.Add(key, value)
}

func WithBody(data []byte) error {
	_, err := buff.Write(data)
	if err != nil {
		return err
	}

	if req.Body == nil {
		req.Body = ioutil.NopCloser(buff)
	}

	return nil
}

// TODO
func WithMultiPart() {

}

func Do() {
	if done {
		return
	}

	done = true

	var err error
	resp, err = client.Do(req)
	if err != nil {
		panic(err)
	}

	// remove previous data
	store.DeleteWithPrefix("$body")
	store.DeleteWithPrefix("$header")
	store.DeleteWithPrefix("$resp")

	// TODO need optimize
	rh, _ := json.Marshal(resp.Header)
	var rm = map[string]interface{}{}
	err = json.Unmarshal(rh, &rm)
	if err == nil {
		store.DataWalker("$header", rm)
		store.DataWalker("$resp.header", rm)
	}

	data, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	var bm interface{}
	err = json.Unmarshal(data, &bm)
	if err == nil {
		store.DataWalker("$body", bm)
		store.DataWalker("$resp.body", bm)
	} else {
		store.Set("$body", string(data))
		store.Set("$resp.body", string(data))
	}
}

func GetResp() *http.Response {
	return resp
}
