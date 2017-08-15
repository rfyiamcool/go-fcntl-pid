# go-fcntl-pid

使用fcntl规避了在并发启动时有几率发生的多次启动问题.

## example:

```
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
```
