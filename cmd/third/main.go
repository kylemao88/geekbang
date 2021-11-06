package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"net/http"
	"os"
	"os/signal"
	"time"
)

var myserver http.Server

func main() {
	http.HandleFunc("/", hello)
	myserver.Addr = ":8080"

	ctx, cancel := context.WithCancel(context.Background())
	go signalListen(cancel)

	g, ctx := errgroup.WithContext(ctx)
	g.Go(func() error {
		// ListenAndServe always returns a non-nil error. After Shutdown or Close,
		err := myserver.ListenAndServe()
		if err == nil {
			fmt.Println("server start succ")
			return nil
		}

		fmt.Println("exit ListenAndServe:", err)
		return nil
	})

	select {
	case <-ctx.Done():
		fmt.Println("ctx.Done: ", ctx.Err())
		sCtx, _ := context.WithTimeout(context.Background(), time.Second*5)
		if err := myserver.Shutdown(sCtx); err != nil {
			fmt.Println("server shut down fail:", err)
		} else {
			fmt.Println("server shut down succ")
		}
	}

	if err := g.Wait(); err != nil {
		fmt.Println("g.Wait:", err)
		return
	}

	fmt.Println("exit main")
	return
}

func signalListen(cancel context.CancelFunc) {
	defer fmt.Println("exit signalListen")

	c := make(chan os.Signal, 1)
	//signal.Notify(c, os.Interrupt)  // 接受指定 ctrl+C 信号
	signal.Notify(c) // 接受所有所有信号
	for {
		select {
		case s := <-c:
			fmt.Printf("接收信号：%s\n", s)
			if s == os.Interrupt {
				cancel()
				return
			}
		}
	}

}

func hello(w http.ResponseWriter, r *http.Request) {
	time.Sleep(time.Second * 5)
	fmt.Fprintln(w, "Hello, Go!")

}
