package utils

import (
	"net/http"
)

func CheckAPI(url string) int {
	resp, err := http.Get(url)
	if err != nil {
		return 500 // Return 500 in case of error
	}
	defer resp.Body.Close()
	return resp.StatusCode
}
