package httpclient

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func sanitizeURL(url string) (string, error) {
	if !strings.HasPrefix(url, "http://") {
		return "http://" + url, nil
	}
	return url, nil
}

// HandleURL ..
func HandleURL(addr string) {
	// validate and parse url
	url, _ := sanitizeURL(addr)

	// create and send request
	request, _ := http.NewRequest(http.MethodGet, url, nil)

	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
	}

	// get md5 hash
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	hash := md5.Sum(body)

	fmt.Println(request.URL, hex.EncodeToString(hash[:]))
}
