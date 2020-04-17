package main

import "log"

type ConnLimiter struct {
	concurrentConn int // 容量(可同时保持连接的最大连接数)
	bucket chan int
}

// 构造函数
func NewConnLimiter(cc int)  *ConnLimiter {
	return &ConnLimiter{
		concurrentConn: cc,
		bucket:         make(chan int, cc), // channel缓冲区
	}
}

func (cl *ConnLimiter) GetConn() bool {
	if len(cl.bucket) >= cl.concurrentConn {
		log.Printf("Reached the rate limitation.")
		return false
	}

	cl.bucket <- 1 // 写入channel
	return true
}

func (cl *ConnLimiter) ReleaseConn() {
	c :=<- cl.bucket
	log.Printf("New connection coming: %d", c)
}

/*
1. 在streamserver中Streaming和Upload files都需要保持长连接，
和之前的api短连接是不一样的。所以在多路长连接同时保持的时候就会出现一个问题，
如果再不断的发起连接，打开视频，总有一个时刻会把我们的server crash掉，因此我们需要
一个流控部分来控制connection。

2. 用tokenbucket来实现流控模块。
bucket：token1，token2...token3
request, 获取token
response，释放token
go语言的理念：shared channel instead of shared memory.
*/