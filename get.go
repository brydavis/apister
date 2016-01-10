package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func (r Request) Get(base string) string {
	return Get(r, base)
}

func Get(r Request, base string) string {
	req := fmt.Sprintf("%s%s", base, r.Serialize())
	fmt.Println(req, "\n")

	res, err := http.Get(req)
	if err != nil {
		fmt.Println(err)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}

	return string(body)
}
