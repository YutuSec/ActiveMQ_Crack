package DataHandle

import (
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
)

func RequestHead(Main string, url string, bodys io.Reader, head map[string]string) (*http.Response, string, string) {
	resq, err := http.NewRequest(Main, url, bodys)
	if err != nil {
		return nil, "", ""
	}
	for key, val := range head {
		resq.Header.Add(key, val)
	}
	resqbody, err := httputil.DumpRequest(resq, true)
	if err != nil {
		return nil, "", ""

	}
	resp, err := Client.Do(resq) //排除全局变量引起的问题，排除resq为空引发的问题,排除因err定义的问题
	if err != nil {
		return nil, "", ""
	}
	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return nil, "", ""
	}
	return resp, string(body), string(resqbody)
}
func RequestHeadUnClose(Main string, url string, bodys io.Reader, head map[string]string) *http.Response {
	resq, err := http.NewRequest(Main, url, bodys)
	if err != nil {
		return nil
	}
	for key, val := range head {
		resq.Header.Add(key, val)
	}
	if err != nil {
		return nil

	}
	resp, err := Client.Do(resq) //排除全局变量引起的问题，排除resq为空引发的问题,排除因err定义的问题
	if err != nil {
		return nil
	}
	return resp
}
