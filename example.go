package main

import (
	"fmt"
	"github.com/rfyiamcool/go-fcntl-pid"
	"time"
)

func pp() {
	for {
		time.Sleep(1 * time.Second)
		fmt.Println("trigger keepalive")
	}
}

func main() {
	filename := "mm.pid"
	pidfile.CheckExit(filename)
	pp()
}
