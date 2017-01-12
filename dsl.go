/**
 * Copyright 2015-2016, Wothing Co., Ltd.
 * All rights reserved.
 *
 * Created by elvizlai on 2016/10/12 17:37.
 */

package main

import (
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
	"file": func(s *string) error {
		varReplacer(s)
		data, err := file.Read(*s)
		if err != nil {
			return err
		}
		*s = string(data)
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

		fmt.Printf("[get] '%s'\n", *s)

		httplib.NewRequest("GET", *s)
		return nil
	},
	"post": func(s *string) error {
		varReplacer(s)

		fmt.Printf("[pos] '%s'\n", *s)

		httplib.NewRequest("POST", *s)
		return nil
	},
	"header": func(s *string) error {
		params := strings.Split(*s, " ")
		if len(params) != 2 {
			return errors.New("header need 2 param")
		}

		varReplacer(&params[0])
		varReplacer(&params[1])

		fmt.Printf("[hed] '%s' '%s'\n", params[0], params[1])

		httplib.WithHeader(strings.TrimSpace(params[0]), strings.TrimSpace(params[1]))
		return nil
	},
	"content": func(s *string) error {
		varReplacer(s)
		fmt.Printf("[cnt] '%s'\n", *s)
		httplib.WithContent(*s)
		return nil
	},
	"body": func(s *string) error {
		varReplacer(s)

		fmt.Printf("[bdy] '%s'\n", *s)

		httplib.WithBody([]byte(*s))
		return nil
	},
	"ret": func(s *string) error {
		params := strings.Split(*s, " ")
		if len(params) > 1 {
			return errors.New("ret need 0 or 1 param")
		}

		httplib.Do()

		fmt.Print("[ret] ")

		if params[0] == "" {
			fmt.Print(httplib.GetResp().StatusCode, "\n")
		} else {
			varReplacer(&params[0])
			fmt.Printf("'%s' = '%s'", params[0], strconv.Itoa(httplib.GetResp().StatusCode))
			if strconv.Itoa(httplib.GetResp().StatusCode) == params[0] {
				passCount++
				fmt.Print(pass, "\n")
			} else {
				failCount++
				fmt.Print(failed, "\n")
			}
		}

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

		fmt.Printf("[ast] '%s' ⊇ '%s'", params[0], params[1])

		if strings.Contains(params[0], params[1]) {
			passCount++
			fmt.Print(pass, "\n")
		} else {
			failCount++
			fmt.Print(failed, "\n")
		}

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

		fmt.Printf("[ast] '%s' = '%s'", params[0], params[1])

		if params[0] == params[1] {
			passCount++
			fmt.Println(pass)
		} else {
			failCount++
			fmt.Println(failed)
		}
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

		fmt.Printf("[ast] '%s' ≠ '%s'", params[0], params[1])

		if params[0] != params[1] {
			passCount++
			fmt.Print(pass, "\n")
		} else {
			failCount++
			fmt.Print(failed, "\n")
		}
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

		fmt.Printf("[ast] '%s' > '%s'", params[0], params[1])

		if params[0] > params[1] {
			passCount++
			fmt.Print(pass, "\n")
		} else {
			failCount++
			fmt.Print(failed, "\n")
		}
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

		fmt.Printf("[ast] '%s' ≥ '%s'", params[0], params[1])
		if params[0] >= params[1] {
			passCount++
			fmt.Print(pass, "\n")
		} else {
			failCount++
			fmt.Print(failed, "\n")
		}
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
	cmds := cmdRgx.FindAllString(*s, -1)
	for i, v := range cmds {
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
