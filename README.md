# xtime
return atime, mtime, ctime

Example:


package main

```go
import (
	"fmt"
	"io/ioutil"
	"time"

	"github.com/immortal/xtime"
)

func main() {
	files, _ := ioutil.ReadDir("./")
	for _, f := range files {
		a := xtime.Get(f).Atime()
		c := xtime.Get(f).Ctime()
		m := xtime.Get(f).Mtime()
		x := time.Now().Unix() - xtime.Get(f).Ctime().Unix()
		fmt.Printf("atime: %s\nctime: %s\nmtime: %s\ndiff from now: %d\n",
			a.Format(time.RFC3339),
			c.Format(time.RFC3339),
			m.Format(time.RFC3339),
			x,
		)
	}
}
```
