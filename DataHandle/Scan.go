package DataHandle

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func Scan() {
	ch := GETURLBase()
	for n := 0; n < Thread; n++ {
		for v := range ch {
			wg.Add(1)
			fmt.Printf("正在探测 [-%v-] ActiveMQ默认口令漏洞\n", v)
			go func(v string) {
				CheckUnauth(v, &wg)
			}(v)
		}
	}
	wg.Wait()
}
