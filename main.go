// 入口包
package main

import (
	"log"
	"time"

	"gitee.com/xiaoyutab/gatway/router"
	"github.com/fvbock/endless"
)

// 使用gin框架搭建自己的网关服务
func main() {
	time.Local = time.FixedZone("CST", 3600*8) // 设置时区为东八区
	r := router.RouterExplode()
	if err := endless.ListenAndServe(":2002", r); err != nil {
		log.Fatalf("listen: %s\n", err)
	}
}
