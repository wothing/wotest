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
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"

	"github.com/wothing/wotest/store"
)

var client = &http.Client{}

var done bool
var buff *bytes.Buffer
var writer *multipart.Writer

var req *http.Request
var resp *http.Response

func NewRequest(method string, url string) error {
	done = false
	buff = bytes.NewBuffer(nil)

	var err error
	req, err = http.NewRequest(method, url, nil)
	return err
}

func WithContent(val string) {
	req.Header.Set("Content-Type", val)
}

func WithHeader(key string, val string) {
	req.Header.Add(key, val)
}

func WithBody(data []byte) error {
	_, err := buff.Write(data)
	return err
}

// TODO
func WithMultiPart(key string, val string) error {
	if writer == nil {
		writer = multipart.NewWriter(buff)
	}

	return writer.WriteField(key, val)
}

func WithMultiPartFile(name string, filePath string) error {
	if writer == nil {
		writer = multipart.NewWriter(buff)
	}

	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	part, err := writer.CreateFormFile(name, filePath)
	if err != nil {
		return err
	}

	_, err = io.Copy(part, file)

	return err
}

func Do() error {
	if done {
		return nil
	}

	done = true

	if writer != nil {
		req.Header.Set("Content-Type", writer.FormDataContentType())
		if err := writer.Close(); err != nil {
			return err
		}
		writer = nil
	}

	req.Body = ioutil.NopCloser(buff)

	var err error
	resp, err = client.Do(req)
	if err != nil {
		return err
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

	return nil
}

func GetResp() *http.Response {
	return resp
}
