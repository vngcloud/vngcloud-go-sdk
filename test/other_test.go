package test

import (
	ltesting "testing"
	ltime "time"
)

func TestUnixNano(t *ltesting.T) {
	now := ltime.Now()

	t.Log(now.UnixNano())

	// convert unix nano to time
	me := ltime.Unix(0, now.UnixNano())
	ltime.Sleep(5 * ltime.Second)
	t.Log(me.UnixNano())

	now = ltime.Now()
	t.Log(now.UnixNano())

	t.Log(now.Sub(me) >= 5*ltime.Second)
}
