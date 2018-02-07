package gormDemo

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type MyInterface map[string]interface{}

func (m MyInterface) Value() (driver.Value, error) {
	b, e := json.Marshal(m)
	if e != nil {
		return "", e
	}
	return string(b), nil
}
func (m *MyInterface) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("assert pg json failed")
	}
	e := json.Unmarshal(b, m)
	if e != nil {
		return e
	}
	return nil
}

type IntSlice []int

func NewIntSlice(args ...int) IntSlice {
	s := make(IntSlice, 0)
	for _, arg := range args {
		s = append(s, arg)
	}
	return s
}

func (s IntSlice) Value() (driver.Value, error) {
	if len(s) == 0 {
		return nil, nil
	}
	str := "{"
	for _, v := range s {
		str += fmt.Sprintf("%d,", v)
	}
	return strings.TrimRight(str, ",") + "}", nil
}

func (s *IntSlice) Scan(input interface{}) error {
	b, ok := input.([]byte)
	if !ok {
		return errors.New("assert pg array failed")
	}
	str := string(b)
	str = strings.TrimLeft(str, "{")
	str = strings.TrimRight(str, "}")
	for _, v := range strings.Split(str, ",") {
		if i, e := strconv.Atoi(v); e == nil {
			*s = append(*s, i)
		} else {
			return errors.New("pg array strToint failed")
		}
	}
	return nil
}

type StrSlice []string

func NewStringSlice(args ...string) StrSlice {
	s := make(StrSlice, 0)
	for _, arg := range args {
		s = append(s, arg)
	}
	return s
}

func (s StrSlice) Value() (driver.Value, error) {
	return "{" + strings.Join(s, ",") + "}", nil
}

func (s *StrSlice) Scan(input interface{}) error {
	b, ok := input.([]byte)
	if !ok {
		return errors.New("assert pg array failed")
	}
	str := string(b)
	str = strings.TrimLeft(str, "{")
	str = strings.TrimRight(str, "}")
	for _, v := range strings.Split(str, ",") {
		*s = append(*s, v)
	}
	return nil
}
