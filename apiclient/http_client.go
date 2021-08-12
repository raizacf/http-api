package apiclient

import (
	"crypto/md5"
	"encoding/hex"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// HTTPClient interface wraps the basic client methods for http requests
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// HashResponse return the MD5 hash of an url get response
func HashResponse(addr string, client HTTPClient) (string, string, error) {
	url, err := validateURL(addr)
	if err != nil {
		return url, "", err
	}

	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return url, "", err
	}

	resp, err := client.Do(request)
	if err != nil {
		return url, "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return url, "", err
	}

	hash := md5.Sum(body)

	return url, hex.EncodeToString(hash[:]), nil
}

func validateURL(addr string) (string, error) {
	if !strings.HasPrefix(addr, "http://") && !strings.HasPrefix(addr, "https://") {
		addr = "http://" + addr
	}
	_, err := url.ParseRequestURI(addr)
	if err != nil {
		return addr, err
	}
	return addr, nil
}
