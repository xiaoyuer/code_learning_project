package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

//defer wg.wait()

func CtxWaitGroup() {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("老财做账")
		wg.Done()
	}()

	go func() {
		time.Sleep(1 * time.Second)
		fmt.Println("老财审单")
		wg.Done()
	}()
	wg.Wait()
	fmt.Println("这就是老财们的日常工作")
}

func main() {
	// CtxWaitGroup()
	// CtxStopInitiative()
	// CtxContext()
	CtxContextManyGoroutine()
}

func CtxStopInitiative() {
	stop := make(chan bool)
	go func() {
		for {
			select {
			case <-stop:
				fmt.Println("You are fired!")
				return
			default:
				fmt.Println("working")
				time.Sleep(1 * time.Second)
			}
		}
	}()
	time.Sleep(5 * time.Second)
	fmt.Println("那个老财动作太慢了！开除！")
	stop <- true
	time.Sleep(5 * time.Second)
	fmt.Println("老财滚蛋了")
}

func CtxContext() {
	ctx, cancle := context.WithCancel(context.Background())
	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("You are fired!")
				return
			default:
				fmt.Println("working")
				time.Sleep(1 * time.Second)
			}
		}
	}()
	time.Sleep(5 * time.Second)
	fmt.Println("那个老财动作太慢了！开除！")
	cancle()
	time.Sleep(1 * time.Second)
	fmt.Println("老财滚蛋了")
}

func CtxContextManyGoroutine() {
	ctx, cancle := context.WithCancel(context.Background())
	go worker(ctx, "老财1")
	go worker(ctx, "老财2")
	go worker(ctx, "老财3")
	time.Sleep(1 * time.Second)
	fmt.Println("建立财务共享中心，老财全部优化！")
	cancle()
	time.Sleep(1 * time.Second)
	fmt.Println("老财们都滚蛋了")
}

func worker(ctx context.Context, str string) {
	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println(str, "你被优化了！")
				return
			default:
				fmt.Println(str, "working")
				time.Sleep(1 * time.Second)
			}
		}
	}()
}
