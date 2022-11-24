// @Author Bing 
// @Date 2022/11/23 15:46:00 
// @Desc
package main

import (
	"fmt"
	"log"
	"math/rand"
	"my_lru_lfu/my_lfu"
)

const (
	Cap = 3
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	//l := my_lru.InitLRU(Cap)
	l := my_lfu.InitLFUCahe(Cap)
	for i := 0; i < Cap; i++ {
		l.Put(i, rand.Intn(1000))
	}
	log.Printf("\n 初始添加了 3 个节点，%c[1;40;32m%s%c[0m\n\n", 0x1B, l.String(), 0x1B)

	data := []int{3, 4, 3, 5}
	var tmp int
	for _, v := range data {
		log.Printf("%c[1;40;32m%s%c[0m\n", 0x1B, fmt.Sprintf("\nGet %d", v), 0x1B)
		l.Get(v)
		if l.Count > tmp {
			fmt.Printf("%c[1;40;31m%s%c[0m\n", 0x1B, l.String(), 0x1B)
			tmp = l.Count
		} else {
			fmt.Printf("%c[1;40;32m%s%c[0m\n", 0x1B, l.String(), 0x1B)
		}

	}
	log.Printf("一共发生了 %d 次缺页中断", l.Count)
}
