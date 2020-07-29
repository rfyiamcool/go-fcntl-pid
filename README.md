# go-fcntl-pid

use syscall.flock() to write pid file, avoid concurrently start to cause bug concurrently .

## example:

```
package main

import (
	"fmt"
	"time"

	"github.com/rfyiamcool/go-fcntl-pid"
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
