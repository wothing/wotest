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
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var client = &http.Client{Timeout: 30 * time.Second}

type httpLib struct {
	method string
	url    string

	reqBody string

	header map[string]string

	resp     *http.Response
	respBody []byte

	done bool
}

var httpReq = httpLib{
	header: make(map[string]string),
}

func (h *httpLib) do() {
	if !h.done {
		h.done = true

		req, err := http.NewRequest(h.method, h.url, bytes.NewBuffer([]byte(h.reqBody)))
		if err != nil {
			panic(err)
		}

		for k, v := range h.header {
			req.Header.Add(k, v)
		}

		h.resp, err = client.Do(req)
		if err != nil {
			panic(err)
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

		var bm = map[string]interface{}{}
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

// TODO may add more type assert
func dataWalker(prefix string, x interface{}) {
	data, _ := json.Marshal(x)
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

	default:
		varMap[prefix] = fmt.Sprint(x) // force write
	}
}
