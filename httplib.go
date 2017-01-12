/**
 * Copyright 2015-2016, Wothing Co., Ltd.
 * All rights reserved.
 *
 * Created by elvizlai on 2016/10/12 17:39.
 */

package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

var client = &http.Client{Timeout: 30 * time.Second}

type httpReq struct {
	method string
	url    string

	header      map[string]string
	contentType string
	reqBody     string

	resp     *http.Response
	respBody []byte

	done bool
}

var req = httpReq{
	header: make(map[string]string),
}

func newReq(method string) httpReq {
	return httpReq{
		method: method,
		header: make(map[string]string),
	}
}

func (h *httpReq) do() {
	if !h.done {
		h.done = true

		req, err := http.NewRequest(h.method, h.url, bytes.NewBuffer([]byte(h.reqBody)))
		if err != nil {
			log.Fatal(err)
		}

		if h.contentType != "" {
			req.Header.Add("Content-Type", h.contentType)
		}

		for k, v := range h.header {
			req.Header.Add(k, v)
		}

		h.resp, err = client.Do(req)
		if err != nil {
			log.Fatal(err)
		}

		h.respBody, _ = ioutil.ReadAll(h.resp.Body)
		defer h.resp.Body.Close()

		// remove previous data
		for k := range varMap {
			if strings.HasPrefix(k, "$body") || strings.HasPrefix(k, "$header") || strings.HasPrefix(k, "$resp") {
				delete(varMap, k)
			}
		}

		// TODO need optimize
		rh, _ := json.Marshal(h.resp.Header)
		var rm = map[string]interface{}{}
		err = json.Unmarshal(rh, &rm)
		if err == nil {
			dataWalker("$header", rm)
			dataWalker("$resp.header", rm)
		}

		var bm interface{}
		err = json.Unmarshal(h.respBody, &bm)
		if err == nil {
			dataWalker("$body", bm)
			dataWalker("$resp.body", bm)
		} else {
			varMap["$body"] = string(h.respBody)
			varMap["$resp.body"] = string(h.respBody)
		}
	}
}
