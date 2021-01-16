package main

import (
	"net/http"
	_ "net/http/pprof"

	"github.com/gmsorg/gms"
	"github.com/gmsorg/gms/example/model"
	"github.com/gmsorg/gms/gmsContext"
)

func main() {

	// 添加业务处理路由（addition是业务处理方法的唯一标识，客户端调用需要使用）
	gms.AddRouter("addition", Addition)

	go func() {
		http.ListenAndServe("0.0.0.0:8899", nil)
	}()

	// 启动，以1024 为启动端口
	gms.Run("127.0.0.1", 1024)
	// gms.DefaultRun()
}

/*
加法计算
*/
func Addition(c *gmsContext.Context) error {
	additionReq := &model.AdditionReq{}
	// 绑定请求参数
	c.Param(additionReq)

	// 结果对象
	additionRes := &model.AdditionRes{}
	additionRes.Result = additionReq.NumberA + additionReq.NumberB

	// fmt.Println(additionRes.Result)
	// 返回结果
	c.Result(additionRes)
	return nil
}
