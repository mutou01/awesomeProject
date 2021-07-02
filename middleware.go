package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func Middleware() gin.HandlerFunc {
	return func(c *gin.Context){
		t := time.Now()
		fmt.Println("中间件开始执行")
		//设置变量到Context的key中，即可通过Get（）获取
		c.Set("request","中间件")
		//执行函数
		c.Next()//添加next（）后，差距80us
		//中间件执行完成后续的操作
		status := c.Writer.Status()
		fmt.Println("中间件执行完毕",status)
		t2 := time.Since(t)
		fmt.Println("time:",t2)
	}
}

func myTiem() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		since := time.Since(start)
		fmt.Println("程序用时：", since)
	}
}

func main() {
	//创建路由，并使用两个默认的中间件：Logger（）、Recovery（）
	r := gin.Default()
	//局部中间件
	//r.GET("/ce", Middleware(), GetMiddleHandler)
	r.Use(Middleware())
	{
		r.GET("/ce", GetMiddleHandler)
	}
	//r.Use(myTime())此功能与Middleware（）重复
	ShoppingGroup := r.Group("/shopping")
	{
		ShoppingGroup.GET("/time1",GetTime1)
		ShoppingGroup.GET("/time2",GetTime2)
	}

	r.Run(":1237")
}

//http://:1237/ce
func GetMiddleHandler(ctx *gin.Context) {
	//取值
	req , _ :=ctx.Get("request")
	fmt.Println("request:",req)
	//页面接收
	ctx.JSON(http.StatusOK,gin.H{
		"request":req,
	})
}
//{
//    "request": "中间件"
//}

func GetTime1(ctx *gin.Context)  {
	time.Sleep(5 * time.Second)
}

func GetTime2(ctx *gin.Context)  {
	time.Sleep(3 * time.Second)
}

















