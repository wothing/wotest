/**
 * Copyright 2015-2016, Wothing Co., Ltd.
 * All rights reserved.
 *
 * Created by elvizlai on 2017/1/12 12:13.
 */

package httplib

import (
	"fmt"
	"testing"
)

func TestGetWithHeader(t *testing.T) {
	var err error
	err = NewRequest("GET", "https://echo.paw.cloud")
	if err != nil {
		t.Fatal(err)
	}

	WithHeader("key", "value")

	fmt.Println(req)
}

func TestPostWithEmptyBody(t *testing.T) {
	err := NewRequest("POST", "https://echo.paw.cloud")
	if err != nil {
		t.Fatal(err)
	}

	WithHeader("key", "value")

	fmt.Println(req)
}

func TestPostWithBody(t *testing.T) {
	err := NewRequest("POST", "https://echo.paw.cloud")
	if err != nil {
		t.Fatal(err)
	}

	WithContent("application/json")
	WithHeader("key", "value")

	WithBody([]byte("{"))
	WithBody([]byte(`"key":"value"`))
	WithBody([]byte(`}`))

	fmt.Println(req)
}
