package xtime

import (
	"os"
	"time"
)

// Xtime methods to get atime, ctime, mtime
type Xtime interface {
	Atime() time.Time
	Ctime() time.Time
	Mtime() time.Time
}

type atime struct {
	t time.Time
}

func (t atime) Atime() time.Time { return t.t }

type ctime struct {
	t time.Time
}

func (t ctime) Ctime() time.Time { return t.t }

type mtime struct {
	t time.Time
}

func (t ctime) Mtime() time.Time { return t.t }

// Get returns Xtime interface Atime(), Ctime(), Mtime()
func Get(f os.FileInfo) Xtime {
	return getTimes(f)
}
