package main

import "net/http"

//ExecuteURL Check that the ingress urls responds
func ExecuteURL(url string) int {

	resp, err := http.Get(url)

	if err != nil {
		return -1
	}

	return resp.StatusCode

}
