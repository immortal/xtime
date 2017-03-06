package xtime

import (
	"os"
	"syscall"
	"time"
)

// timespec implements Xtime
type timespec struct {
	atime
	ctime
	mtime
}

func timespecToTime(ts syscall.Timespec) time.Time {
	return time.Unix(int64(ts.Sec), int64(ts.Nsec))
}

func getTimes(f os.FileInfo) (x timespec) {
	stat := f.Sys().(*syscall.Stat_t)
	x.atime.t = timespecToTime(stat.Atimespec)
	x.ctime.t = timespecToTime(stat.Ctimespec)
	x.mtime.t = timespecToTime(stat.Mtimespec)
	return x
}
