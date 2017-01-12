/**
 * Copyright 2015-2016, Wothing Co., Ltd.
 * All rights reserved.
 *
 * Created by elvizlai on 2016/10/12 17:37.
 */

package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/wothing/wotest/file"
	"github.com/wothing/wotest/httplib"
	"github.com/wothing/wotest/log"
	"github.com/wothing/wotest/store"
)

const pass = "\033[0;32m ✔ \033[0m"
const failed = "\033[0;31m ✘ \033[0m"

var funcMap = map[string]func(*string) error{
	"env": func(s *string) error {
		params := strings.Split(*s, " ")
		if len(params) > 1 {
			return errors.New("env need 1 param")
		}
		varReplacer(&params[0])
		*s = os.Getenv(params[0])
		return nil
	},
	"file": func(s *string) error {
		varReplacer(s)
		data, err := file.Read(*s)
		if err != nil {
			return err
		}
		*s = string(data)
		return nil
	},
	"echo": func(s *string) error {
		varReplacer(s)
		log.Debugf(*s)
		return nil
	},
	"pretty": func(s *string) error {
		varReplacer(s)
		var t interface{}
		err := json.Unmarshal([]byte(*s), &t)
		if err != nil {
			return err
		} else {
			b, err := json.MarshalIndent(t, "", "  ")
			if err != nil {
				return err
			}
			log.Debugf(string(b))
		}
		return nil
	},
	"set": func(s *string) error {
		params := strings.Split(*s, " ")
		if len(params) != 2 {
			return errors.New("set need 2 param")
		}

		varReplacer(&params[0])
		varReplacer(&params[1])

		var x interface{}
		err := json.Unmarshal([]byte(params[1]), &x)
		if err == nil {
			store.DataWalker("$"+params[0], x)
		} else {
			store.Set("$"+params[0], params[1])
		}

		return nil
	},
	"get": func(s *string) error {
		varReplacer(s)

		log.Printf("[get] '%s'", *s)

		return httplib.NewRequest("GET", *s)
	},
	"post": func(s *string) error {
		varReplacer(s)

		log.Printf("[pos] '%s'", *s)

		return httplib.NewRequest("POST", *s)
	},
	"header": func(s *string) error {
		params := strings.Split(*s, " ")
		if len(params) != 2 {
			return errors.New("header need 2 param")
		}

		varReplacer(&params[0])
		varReplacer(&params[1])

		log.Printf("[hed] '%s' '%s'", params[0], params[1])

		httplib.WithHeader(strings.TrimSpace(params[0]), strings.TrimSpace(params[1]))
		return nil
	},
	"content": func(s *string) error {
		varReplacer(s)
		log.Printf("[cnt] 'Content-Type' '%s'", *s)
		httplib.WithContent(*s)
		return nil
	},
	"body": func(s *string) error {
		varReplacer(s)

		log.Printf("[bdy] '%s'", *s)

		return httplib.WithBody([]byte(*s))
	},
	"form": func(s *string) error {
		params := strings.Split(*s, " ")
		if len(params) != 2 {
			return errors.New("form need 2 param")
		}

		varReplacer(&params[0])
		varReplacer(&params[1])

		log.Printf("[frm] '%s' '%s'", params[0], params[1])

		return httplib.WithForm(params[0], params[1])
	},
	"multipart": func(s *string) error {
		params := strings.Split(*s, " ")
		if len(params) != 2 {
			return errors.New("multipart need 2 param")
		}

		varReplacer(&params[0])
		varReplacer(&params[1])

		log.Printf("[mup] '%s' '%s'", params[0], params[1])

		return httplib.WithMultiPart(params[0], params[1])
	},
	"multifile": func(s *string) error {
		params := strings.Split(*s, " ")
		if len(params) != 2 {
			return errors.New("multifile need 2 param")
		}

		varReplacer(&params[0])
		varReplacer(&params[1])

		log.Printf("[muf] '%s' '%s'", params[0], params[1])

		return httplib.WithMultiPartFile(params[0], params[1])
	},
	"ret": func(s *string) error {
		params := strings.Split(*s, " ")
		if len(params) > 1 {
			return errors.New("ret need 0 or 1 param")
		}

		err := httplib.Do()
		if err != nil {
			log.Fatalf(err.Error())
		}

		buff := bytes.NewBuffer(nil)
		buff.WriteString("[ret] ")

		if params[0] == "" {
			buff.WriteString(strconv.Itoa(httplib.GetResp().StatusCode))
		} else {
			varReplacer(&params[0])
			buff.WriteString(fmt.Sprintf("'%s' = '%s'", params[0], strconv.Itoa(httplib.GetResp().StatusCode)))
			*s = marker(buff, strconv.Itoa(httplib.GetResp().StatusCode) == params[0])
		}

		log.Printf(buff.String())
		return nil
	},
	// has
	"has": func(s *string) error {
		params := strings.Split(*s, " ")
		if len(params) != 2 {
			return errors.New("contains need 2 param")
		}

		varReplacer(&params[0])
		varReplacer(&params[1])

		buff := bytes.NewBuffer(nil)
		buff.WriteString(fmt.Sprintf("[ast] '%s' ⊇ '%s'", params[0], params[1]))
		*s = marker(buff, strings.Contains(params[0], params[1]))
		log.Printf(buff.String())
		return nil
	},
	// equal
	"eq": func(s *string) error {
		params := strings.Split(*s, " ")
		if len(params) != 2 {
			return errors.New("match need 2 param")
		}

		varReplacer(&params[0])
		varReplacer(&params[1])

		buff := bytes.NewBuffer(nil)
		buff.WriteString(fmt.Sprintf("[ast] '%s' = '%s'", params[0], params[1]))
		*s = marker(buff, params[0] == params[1])
		log.Printf(buff.String())
		return nil
	},
	// not equal
	"neq": func(s *string) error {
		params := strings.Split(*s, " ")
		if len(params) != 2 {
			return errors.New("neq need 2 param")
		}

		varReplacer(&params[0])
		varReplacer(&params[1])

		buff := bytes.NewBuffer(nil)
		buff.WriteString(fmt.Sprintf("[ast] '%s' ≠ '%s'", params[0], params[1]))
		*s = marker(buff, params[0] != params[1])
		log.Printf(buff.String())
		return nil
	},
	// greater than
	"gt": func(s *string) error {
		params := strings.Split(*s, " ")
		if len(params) != 2 {
			return errors.New("gt need 2 param")
		}

		varReplacer(&params[0])
		varReplacer(&params[1])

		buff := bytes.NewBuffer(nil)
		buff.WriteString(fmt.Sprintf("[ast] '%s' > '%s'", params[0], params[1]))
		*s = marker(buff, params[0] > params[1])
		log.Printf(buff.String())
		return nil
	},
	// greater than or equal
	"gte": func(s *string) error {
		params := strings.Split(*s, " ")
		if len(params) != 2 {
			return errors.New("gte need 2 param")
		}

		varReplacer(&params[0])
		varReplacer(&params[1])

		buff := bytes.NewBuffer(nil)
		buff.WriteString(fmt.Sprintf("[ast] '%s' ≥ '%s'", params[0], params[1]))
		*s = marker(buff, params[0] >= params[1])
		log.Printf(buff.String())
		return nil
	},
	// less than
	"lt": func(s *string) error {
		params := strings.Split(*s, " ")
		if len(params) != 2 {
			return errors.New("lt need 2 param")
		}

		varReplacer(&params[0])
		varReplacer(&params[1])

		buff := bytes.NewBuffer(nil)
		buff.WriteString(fmt.Sprintf("[ast] '%s' < '%s'", params[0], params[1]))
		*s = marker(buff, params[0] < params[1])
		log.Printf(buff.String())
		return nil
	},
	// less than or equal
	"lte": func(s *string) error {
		params := strings.Split(*s, " ")
		if len(params) != 2 {
			return errors.New("lte need 2 param")
		}

		varReplacer(&params[0])
		varReplacer(&params[1])

		buff := bytes.NewBuffer(nil)
		buff.WriteString(fmt.Sprintf("[ast] '%s' ≤ '%s'", params[0], params[1]))
		*s = marker(buff, params[0] <= params[1])
		log.Printf(buff.String())
		return nil
	},
	// TODO
	"regexp": func(s *string) error {
		params := strings.Split(*s, " ")
		if len(params) != 2 {
			return errors.New("regexp need 2 param")
		}

		varReplacer(&params[0])
		//varReplacer(&params[1])

		return nil
	},
}

func eval(s *string) {
	replacer(s)

	// internal cmd
	cmdList := cmdRgx.FindAllString(*s, -1)
	for i, v := range cmdList {
		temp := v
		v = strings.Trim(temp, "`")
		eval(&v)
		store.Set("$.temp_"+strconv.Itoa(i), v)
		*s = strings.Replace(*s, temp, "$.temp_"+strconv.Itoa(i), 1)
	}

	x := strings.Split(*s, " ")

	if f, ok := funcMap[x[0]]; ok {
		suffix := strings.TrimSpace(strings.TrimPrefix(*s, x[0]))
		err := f(&suffix)
		if err != nil {
			log.Fatalf(err.Error())
		}
		*s = suffix
	} else {
		failCount++
		log.Warnf("no such method: '%s'", x[0])
	}
}

func marker(buff *bytes.Buffer, tf bool) string {
	if tf {
		passCount++
		buff.WriteString(pass)
		return "T"
	} else {
		failCount++
		buff.WriteString(failed)
		return "F"
	}
}
