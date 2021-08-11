package httpclient

// import (
// 	"net/http"
// )

// // HTTPClient interface
// type HTTPClient interface {
// 	Do(req *http.Request) (*http.Response, error)
// }

// validate and parse url

// create and send request

// get md5 hash

// print result
//fmt.Println(request.URL, hex.EncodeToString(hash[:]))
//fmt.Println()

// request, _ := http.NewRequest(http.MethodGet, "", nil)

// if request.URL.Scheme == "" {
// 	request.URL.Scheme = "http"
// }
// if request.URL.Host == "" {
// 	request.URL.Host = strings.ReplaceAll(addr, "http://", "")
// }

// client := &http.Client{}
// resp, err := client.Do(request)
// if err != nil {
// 	fmt.Println(err)
// }

// defer resp.Body.Close()
// body, _ := ioutil.ReadAll(resp.Body)
// hash := md5.Sum(body)
// fmt.Println(request.URL, hex.EncodeToString(hash[:]))
// fmt.Println()
