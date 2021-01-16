# GMS

## 简介

​	GMS全称Golang Micro Service，这里取单词首字母作为项目名称。

​	项目是2020-10-24日（程序员节）正式在Github上开源的。所以默认启动端口就定为1024。

​	GMS是一个非常简单易用的微服务框架。只要您之前使用过类似Gin、beego这样的web框架。就能快速上手，使用方法和这些web框架一样简单。不用额外编写类似proto等额外的接口协议。

​	GMS的网络通信层使用[gnet](https://github.com/panjf2000/gnet) ，基于gnet的优异性能GMS也会表现不俗

## 特点：

- 非常简单、学习成本极低（不用再感叹学不动了）。GMS处于初期阶段您想参与开发也非常简单。

- 不用定义proto等协议文件。

​		写proto文件不仅麻烦还容易出错。而且使用协议文件定义服务，最终在框架内部实现都要使用反射去调用目标方法。众所周知反射调用方法性能比直接调用要差很多。所以如果其他条件不变的情况下，不使用反射直接执行目标方法的方式性能肯定比反射要好

## 缺点：

- 目前很多功能还没有完善，不建议应用在公司项目中。



## 快速开始

下载源码 进入 example 直接运行。或者按照以下步骤自己搭建Demo运行

下面我们以一个加法计算服务为类

### 1：定义请求和返回对象

```go
package model

type AdditionReq struct {
	NumberA int
	NumberB int
}

type AdditionRes struct {
	Result int
}
```

### 2：开发服务端

```go
package main

import (
	"github.com/gmsorg/gms"
	"github.com/gmsorg/gms/example/model"
	"github.com/gmsorg/gms/gmsContext"
)

func main() {
	// 添加业务处理路由（addition是业务处理方法的唯一标识，客户端调用需要使用）
	gms.AddRouter("addition", Addition)

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

	// 返回结果
	c.Result(additionRes)
	return nil
}
```

### 3：开发客户端

```go
package main

import (
	"fmt"
	"log"

	"github.com/gmsorg/gms/client"
	"github.com/gmsorg/gms/codec"
	"github.com/gmsorg/gms/discovery"
	"github.com/gmsorg/gms/example/model"
)

/*
客户端
*/
func main() {
	// 初始化一个点对点服务发现对象
	discovery := discovery.NewP2PDiscover([]string{"127.0.0.1:1024"})

	// 初始化一个客户端对象
	additionClient, err := client.NewClient(discovery)
	if err != nil {
		log.Println(err)
		return
	}

	// 设置 Msgpack 序列化器，默认也是 Msgpack
	additionClient.SetCodecType(codec.Msgpack)

	// 请求对象
	req := &model.AdditionReq{NumberA: 10, NumberB: 20}
	// 接收返回值的对象
	res := &model.AdditionRes{}

	// 调用服务
	err = additionClient.Call("addition", req, res)
	if err != nil {
		log.Println(err)
	}
	log.Println(fmt.Sprintf("%d+%d=%d", req.NumberA, req.NumberB, res.Result))
}
```

## 特别鸣谢

感谢以下开源项目作者的辛苦付出

[github.com/panjf2000/gnet](https://github.com/panjf2000/gnet)

[github.com/abronan/valkeyrie](https://github.com/abronan/valkeyrie)

[github.com/satori/go.uuid](https://github.com/satori/go.uuid)

[github.com/vmihailenco/msgpack](https://github.com/vmihailenco/msgpack)

## TODO List

- [x]  服务端支持 客户端指定序列化方式
- [x]  服务注册&服务发现
  - [x] redis 注册中心&服务发现
  - [x] etcd3 注册中心&服务发现
  - [x] consul 注册中心&服务发现
- [ ] 流控
- [ ] 熔断
- [ ] 监控统计

