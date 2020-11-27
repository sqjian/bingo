package cmdserver

import (
	"expvar"
	"time"
)

var up = expvar.NewString("up")

func init() {
	up.Set(time.Now().String())
}
