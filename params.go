package main

import (
	"fmt"
	"strings"
)

func (r Request) Serialize() string {
	return Serialize(r)
}

func Serialize(r Request) (params string) {
	for k, v := range r {
		switch t := v.(type) {
		case string:
			var s string
			for _, word := range strings.Split(t, " ") {
				s += fmt.Sprintf("%s+", word)
			}
			params += fmt.Sprintf("&%s=%s", k, s[:len(s)-1])

		case int:
			params += fmt.Sprintf("&%s=%d", k, t)

		case float64:
			params += fmt.Sprintf("&%s=%f", t)

		case []float64:
			var s string
			for _, f := range t {
				s += fmt.Sprintf("%v,", f)
			}
			params += fmt.Sprintf("&%s=%s", k, s[:len(s)-1])

		case []int:
			var s string
			for _, f := range t {
				s += fmt.Sprintf("%v,", f)
			}
			params += fmt.Sprintf("&%s=%s", k, s[:len(s)-1])

		case []uint:
			var s string
			for _, f := range t {
				s += fmt.Sprintf("%v,", f)
			}
			params += fmt.Sprintf("&%s=%s", k, s[:len(s)-1])

		case []interface{}:
			var s string
			for _, f := range t {
				s += fmt.Sprintf("%v,", f)
			}
			params += fmt.Sprintf("&%s=%s", k, s[:len(s)-1])

		case []string:
			var s string

			for _, f := range t {
				s += fmt.Sprintf("%v,", f)
			}
			params += fmt.Sprintf("&%s=%s", k, s[:len(s)-1])

		default:
			params += fmt.Sprintf("&%s=%v", t)

		}
	}
	return
}
