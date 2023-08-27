package router

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"gitee.com/xiaoyutab/gatway/model"
	"gitee.com/xiaoyutab/xgotool"
	"gitee.com/xiaoyutab/xgotool/https"
	"github.com/gin-gonic/gin"
	wraphh "github.com/turtlemonvh/gin-wraphh"
)

var r *gin.Engine

func RouterExplode() *gin.Engine {
	if r == nil {
		// 路由加载
		r = gin.Default()
		// 路由格式：
		// <app>/router
		// _ 作为预留APP名称，禁止使用，即APP名称不允许为单独一个_
		// 测试路由
		r.GET("/_/test", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"code": 0,
				"msg":  "测试接口返回结果",
				"data": gin.H{},
				"flag": true,
				"time": time.Now().Format(time.DateTime),
			})
		})
		// 平滑重启路由信息
		r.GET("/_/reload", func(ctx *gin.Context) {
			// 进程PID
			pid := os.Getpid()
			// 平滑重启pid
			xgotool.Exec("kill", "-1", fmt.Sprintf("%d", pid))
			ctx.JSON(200, gin.H{
				"code": 0,
				"msg":  "重启完成",
				"data": gin.H{},
				"flag": true,
				"time": time.Now().Format(time.DateTime),
			})
		})
		// 测试网关服务
		r.Any("www", wraphh.WrapHH(func(h http.Handler) http.Handler {
			u := URL{AppRouter: model.AppRouter{Targets: "[\"https://csapi1.xinfushenghuo.cn/.env\"]"}}
			return https.Gatway(h, &u)
		}))
	}
	return r
}
