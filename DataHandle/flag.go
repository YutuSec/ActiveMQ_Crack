package DataHandle

import (
	"flag"
	"fmt"
	"os"
)

var (
	h          bool
	TargetFile string
	Target     string
	Thread     int
)

func init() {
	fmt.Println(`  
   ___  ________________   ________  _______     ________  ___  _______ __
  / _ |/ ___/_  __/  _/ | / / __/  |/  / __ \   / ___/ _ \/ _ |/ ___/ //_/
 / __ / /__  / / _/ / | |/ / _// /|_/ / /_/ /  / /__/ , _/ __ / /__/ ,<   
/_/ |_\___/ /_/ /___/ |___/___/_/  /_/\___\_\  \___/_/|_/_/ |_\___/_/|_|  `, "\n\n\t\t\t\t\t\t\t\t\tBY: 玉兔开源漏洞工具实践项目")
	fmt.Println("\n------------------------------------------------------------------------------\n1、ActiveMQ 默认口令漏洞\n\n2、ActiveMQ任意文件写入漏洞（CVE-2016-3088）\n\n------------------------------------------------------------------------------\n")
	flag.StringVar(&TargetFile, "TF", "", "批量目标：-TF url.txt -t 500")
	flag.StringVar(&Target, "T", "", "单个目标：-T http://127.0.0.1")
	flag.IntVar(&Thread, "t", 500, "并发数量")
	flag.BoolVar(&h, "h", false, "Help")

	// 修改提示信息

	flag.Usage = usage
	flag.Parse()
	if h || ((TargetFile == "") && (Target == "")) {
		flag.Usage()
		os.Exit(0)
	}
}

func usage() {
	fmt.Fprintf(os.Stderr, "Usage:\n\n")
	flag.PrintDefaults()

}
