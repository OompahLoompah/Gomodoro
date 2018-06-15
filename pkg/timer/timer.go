package timer

import (
	"time"
)

type fn func() 

func Timer(t int, notifier fn) {
	time.Sleep(time.Duration(t) * time.Second)
	if notifier != nil {
		notifier()
	}
}
