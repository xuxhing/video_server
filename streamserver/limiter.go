package main

import (
	"log"
)

type ConnLimiter struct {
	concurrentConn int
	bucket chan int
}

func NewConnLimiter(cc int) *ConnLimiter {
	return &ConnLimiter {
		concurrentConn: cc, 
		bucket: make(chan int, cc)
	}
}

func (cl *ConnLimiter) GetConn() bool {
	if len(c1.bucket) >= c1.concurrentConn {
		log.Printf("Reached the rate limitation.")
	}
	c1.bucket <= 1
	return true
}

func (cl *ConnLimiter) ReleaseConn() {
	c:=<-c1.bucket
	log.Printf("New connection ready: %d", c)
}