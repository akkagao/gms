package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"runtime/pprof"
	"sync"
	"time"

	"github.com/google/gops/agent"

	"github.com/gmsorg/gms/client"
	"github.com/gmsorg/gms/discovery"
	"github.com/gmsorg/gms/example/model"
	"github.com/gmsorg/gms/serialize"
)

/*
客户端
*/
func main() {
	cpuf, err := os.Create("cpu_profile")
	if err != nil {
		log.Fatal(err)
	}
	pprof.StartCPUProfile(cpuf)
	defer pprof.StopCPUProfile()

	fm, err := os.OpenFile("mem.out", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	pprof.WriteHeapProfile(fm)
	fm.Close()

	// profile.Start(profile.CPUProfile,profile.ProfilePath(".")).Stop()
	// profile.Start(profile.MemProfile, profile.ProfilePath("."), profile.NoShutdownHook)

	// p := profile.Start(profile.MemProfileAllocs, profile.ProfilePath("."), profile.NoShutdownHook)

	if err := agent.Listen(agent.Options{}); err != nil {
		log.Fatal(err)
	}
	// time.Sleep(time.Hour)

	// 初始化一个点对点服务发现对象
	discovery := discovery.NewP2PDiscover([]string{"127.0.0.1:1024"})

	// 初始化一个客户端对象
	additionClient, err := client.NewClient(discovery)
	if err != nil {
		log.Println(err)
		return
	}

	// 设置 Msgpack 序列化器，默认也是 Msgpack
	additionClient.SetSerializeType(serialize.Msgpack)

	cs, t := 100, 100
	var callt, callOldt time.Duration
	{
		start := time.Now()
		call(additionClient, cs, t)
		callt = time.Since(start)
	}
	// {
	// 	start := time.Now()
	// 	callOld(additionClient, cs, t)
	// 	callOldt = time.Since(start)
	// }
	fmt.Println("callt:", callt, "callOldt:", callOldt)
	// time.Sleep(time.Hour)
}

func call(additionClient client.IClient, cs, t int) {
	waitGroup := sync.WaitGroup{}
	for i := 0; i < cs; i++ {
		waitGroup.Add(1)
		go func(i int) {
			// fmt.Println("启动：", i)
			for j := 0; j < t; j++ {
				// fmt.Println("启动：", i, "-", j)
				rand.Seed(time.Now().UnixNano())
				// req := &model.AdditionReq{NumberA: 100, NumberB: 200}
				req := &model.AdditionReq{NumberA: rand.Intn(100000), NumberB: rand.Intn(100000)}

				// 接收返回值的对象
				res := &model.AdditionRes{}

				// 调用服务
				err := additionClient.Call("addition", req, res)
				if err != nil {
					log.Println(err)
				}
				// log.Println(fmt.Sprintf("call %v-%v : %d+%d=%d  right:%v", i, j, req.NumberA, req.NumberB, res.Result, res.Result == req.NumberA+req.NumberB))
				if j%10 == 0 && i%10 == 0 {
					log.Println(fmt.Sprintf("i:%v j:%v right:%v", i, j, res.Result == req.NumberA+req.NumberB))
				}
			}
			waitGroup.Done()
		}(i)
	}
	waitGroup.Wait()

}

// func callOld(additionClient client.IClient, cs, t int) {
//
// 	waitGroup := sync.WaitGroup{}
// 	for i := 0; i < cs; i++ {
// 		waitGroup.Add(1)
// 		go func(i int) {
// 			for j := 0; j < t; j++ {
// 				rand.Seed(time.Now().UnixNano())
// 				// req := &model.AdditionReq{NumberA: 100, NumberB: 200}
// 				req := &model.AdditionReq{NumberA: rand.Intn(100), NumberB: rand.Intn(200)}
//
// 				// 接收返回值的对象
// 				res := &model.AdditionRes{}
//
// 				// 调用服务
// 				err := additionClient.CallOld("addition", req, res)
// 				if err != nil {
// 					log.Println(err)
// 				}
// 				log.Println(fmt.Sprintf("callOld:%v-%v : %d+%d=%d", i, j, req.NumberA, req.NumberB, res.Result))
// 			}
// 			waitGroup.Done()
// 		}(i)
// 	}
// 	waitGroup.Wait()
// }
