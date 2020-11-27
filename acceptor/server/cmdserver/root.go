package cmdserver

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
)

var m map[string]string

func init() {
	m = map[string]string{
		"expvar_path": "/debug/vars",
		"pprof_path":  "/debug/pprof",
	}
}

func rootHandler(c *gin.Context) {
	d, _ := json.Marshal(m)
	_, _ = fmt.Fprintln(c.Writer, string(d))
}
