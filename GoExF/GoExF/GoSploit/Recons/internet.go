package Recons

import (
	"net/http"
)

func IsInternet() bool {
	resp, err := http.Get("http://google.com")
	if err != nil {
		return false
	}
	defer resp.Body.Close()
	return true

}
