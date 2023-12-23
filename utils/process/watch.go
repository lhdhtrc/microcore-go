package process

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func Watch(handle func()) {
	fmt.Println("----- watch process start -----")

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT)
	// 接收信号
	_ = <-ch

	handle()

	fmt.Println("----- watch process end -----")
}
