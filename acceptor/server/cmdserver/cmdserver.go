package cmdserver

import (
	"context"
	"fmt"
	"github.com/gin-contrib/expvar"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/sqjian/bingo/acceptor/schema"
	"github.com/sqjian/bingo/pkg/errors"
	"github.com/sqjian/bingo/pkg/log"
	"net/http"
)

type CmdServer struct {
	viper *viper.Viper

	engine *gin.Engine
	server *http.Server

	addr string
}

func (h *CmdServer) Name() string {
	return "cmdserver"
}

func (h *CmdServer) Init() errors.Error {
	log.GetDefLogger().Infof("%v Init...\n", h.Name())
	h.addr = h.viper.GetString(fmt.Sprintf("%v.addr", h.Name()))
	if len(h.addr) == 0 {
		return errors.NewWithErr(fmt.Errorf("empty addr for %v.addr", h.Name()))
	}

	gin.SetMode(gin.ReleaseMode)
	h.engine = gin.New()
	h.engine.GET("/", rootHandler)
	h.engine.GET("/debug/vars", expvar.Handler())
	pprof.Register(h.engine, "debug/pprof")
	h.server = &http.Server{
		Addr:    h.addr,
		Handler: h.engine,
	}

	return nil
}

func (h *CmdServer) FInit() errors.Error {
	log.GetDefLogger().Infof("%v FInit...\n", h.Name())
	return nil
}

func (h *CmdServer) Run(ctx context.Context) errors.Error {

	go func() {
		log.GetDefLogger().Infof("about bind %v to %v\n", h.Name(), h.addr)
		if err := h.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			panic(fmt.Sprintf("listen: %s\n", err))
		}
	}()

	select {
	case <-ctx.Done():
		{
			log.GetDefLogger().Infof("about to shutdown server.")
			if err := h.server.Shutdown(ctx); err != nil {
				return errors.NewWithErr(fmt.Errorf("%v forced to shutdown->err:%w", h.Name(), err))
			}
		}
	}
	return nil
}
func NewServer(viper *viper.Viper) (schema.Server, errors.Error) {
	return &CmdServer{viper: viper}, nil
}
