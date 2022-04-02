package DataHandle

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"regexp"
	"time"
)

var tr = &http.Transport{
	MaxIdleConns:    100,
	TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
}
var Client = &http.Client{
	Transport: tr,
	Timeout:   30 * time.Second,
}

const (
	ip4Pattern = `((25[0-5]|2[0-4]\d|[01]?\d\d?)\.){3}(25[0-5]|2[0-4]\d|[01]?\d\d?)`
	ip6Pattern = `(([0-9A-Fa-f]{1,4}:){7}([0-9A-Fa-f]{1,4}|:))|` +
		`(([0-9A-Fa-f]{1,4}:){6}(:[0-9A-Fa-f]{1,4}|((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3})|:))|` +
		`(([0-9A-Fa-f]{1,4}:){5}(((:[0-9A-Fa-f]{1,4}){1,2})|:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3})|:))|` +
		`(([0-9A-Fa-f]{1,4}:){4}(((:[0-9A-Fa-f]{1,4}){1,3})|((:[0-9A-Fa-f]{1,4})?:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:))|` +
		`(([0-9A-Fa-f]{1,4}:){3}(((:[0-9A-Fa-f]{1,4}){1,4})|((:[0-9A-Fa-f]{1,4}){0,2}:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:))|` +
		`(([0-9A-Fa-f]{1,4}:){2}(((:[0-9A-Fa-f]{1,4}){1,5})|((:[0-9A-Fa-f]{1,4}){0,3}:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:))|` +
		`(([0-9A-Fa-f]{1,4}:){1}(((:[0-9A-Fa-f]{1,4}){1,6})|((:[0-9A-Fa-f]{1,4}){0,4}:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:))|` +
		`(:(((:[0-9A-Fa-f]{1,4}){1,7})|((:[0-9A-Fa-f]{1,4}){0,5}:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:))`

	// 同时匹配 IP4 和 IP6
	ipPattern = "(" + ip4Pattern + ")|(" + ip6Pattern + ")"
	// 匹配域名
	domainPattern = `[a-zA-Z0-9][a-zA-Z0-9_-]{0,62}(\.[a-zA-Z0-9][a-zA-Z0-9_-]{0,62})*(\.[a-zA-Z][a-zA-Z0-9]{0,10}){1}`
	urlPattern    = `((https|http|ftp|rtsp|mms)?://)?` + // 协议
		`(([0-9a-zA-Z]+:)?[0-9a-zA-Z_-]+@)?` + // pwd:user@
		"(" + ipPattern + "|(" + domainPattern + "))" + // GetInfo 或域名
		`(:\d{1,5})?` + // 端口
		`(/+[a-zA-Z0-9][a-zA-Z0-9_.-]*)*/*` + // path
		`(\?([a-zA-Z0-9_-]+(=.*&?)*)*)*` // query
)

func GETURLBase() chan string {
	var vals []string
	var err error
	if TargetFile != "" && Target == "" {
		vals, err = ReadConf(TargetFile)
		if err != nil {
			fmt.Println(err)
			return nil
		}
	} else if Target != "" {
		vals = append(vals, Target)
	}
	var ch = make(chan string, len(vals))
	for _, v := range vals {

		if match, _ := regexp.MatchString(urlPattern, v); match { //获取url根目录
			reg := regexp.MustCompile(`((https|http|ftp|rtsp|mms)?://)?` + // 协议
				`(([0-9a-zA-Z]+:)?[0-9a-zA-Z_-]+@)?` + // pwd:user@
				"(" + ipPattern + "|(" + domainPattern + "))" + // GetInfo 或域名
				`(:\d{1,5})?`)
			url := reg.FindString(v)
			ch <- url
		}
	}
	close(ch)
	return ch
}
