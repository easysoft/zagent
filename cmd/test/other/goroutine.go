package other

import (
	"context"
	"log"
	"sync"
	"time"
)

func main() {
	ctx := context.TODO()

	Run(ctx)
}

func Run(ctx context.Context) {
	withTimeout, cancelFunc := context.WithTimeout(ctx, time.Second*10)
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(2)
	go func() { //协程1
		time.Sleep(time.Second * 3)
		log.Printf("g1 end")
		waitGroup.Done()
	}()

	go func() { //协程2
		time.Sleep(time.Second * 6)
		log.Printf("g2 end")
		waitGroup.Done()
	}()

	go func() { //协程3  监听协程1、协程2是否完成
		select {
		case <-withTimeout.Done(): //part1
			return //结束监听协程
		default: //part2 等待协程1、协程2执行完毕，执行完毕后就手动取消上下文，停止阻塞
			waitGroup.Wait()
			cancelFunc()
			return //结束监听协程
		}
	}()

	<-withTimeout.Done()

	log.Printf("all end")
}
