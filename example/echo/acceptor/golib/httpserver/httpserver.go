package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/sqjian/bingo/acceptor/schema"
	"github.com/sqjian/bingo/pkg/errors"
	"github.com/sqjian/bingo/plugin/proto"
	"github.com/sqjian/bingo/plugin/scheduler/minimal"
	"github.com/sqjian/toolkit/uuid"
	"net/http"
	"strings"
)

type HttpServer struct {
	addr  string
	viper *viper.Viper

	engine *gin.Engine
	server *http.Server

	dag *minimal.Minimal
}

func (h *HttpServer) Name() string {
	return "httpserver"
}

func (h *HttpServer) Init() errors.Error {
	fmt.Printf("%v Init...\n", h.Name())
	h.addr = h.viper.GetString(fmt.Sprintf("%v.addr", h.Name()))
	if len(h.addr) == 0 {
		return errors.NewWithErr(fmt.Errorf("empty addr for %v.addr,AllKeys:%v", h.Name(), h.viper.AllKeys()))
	}

	gin.SetMode(gin.ReleaseMode)
	h.engine = gin.New()
	h.server = &http.Server{
		Addr:    h.addr,
		Handler: h.engine,
	}

	minimalInst, minimalInstErr := minimal.NewMinimal(h.viper)
	if minimalInstErr != nil {
		return errors.NewWithErr(minimalInstErr)
	}
	h.dag = minimalInst
	return nil
}

func (h *HttpServer) FInit() errors.Error {
	fmt.Printf("%v FInit...\n", h.Name())
	return nil
}

func (h *HttpServer) Run(ctx context.Context) errors.Error {

	h.engine.POST("/", func(c *gin.Context) {
		bodyData, bodyDataErr := c.GetRawData()
		if bodyDataErr != nil {
			_, _ = fmt.Fprintln(c.Writer, bodyDataErr.Error())
			return
		}

		data := make(map[string]string)
		_ = json.Unmarshal(bodyData, &data)

		_uuid, _uuidErr := uuid.GenerateUuidV1()
		if _uuidErr != nil {
			_, _ = fmt.Fprintln(c.Writer, _uuidErr.Error())
			return
		}

		msg := &proto.Msg{
			Id: _uuid,
			DataList: []*proto.Data{
				{
					Id:   _uuid,
					Data: []byte(data["data"]),
				},
			},
		}

		processRst, processErr := h.dag.Process(minimal.Dag{Id: _uuid, Steps: strings.Split(data["action"], ",")}, msg)

		if processErr != nil {
			_, _ = fmt.Fprintln(c.Writer, processRst)
			return
		}

		_, _ = fmt.Fprintln(c.Writer, string(processRst))
	})

	go func() {
		fmt.Printf("about bind %v to %v\n", h.Name(), h.addr)
		if err := h.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			panic(fmt.Sprintf("listen: %s\n", err))
		}
	}()

	select {
	case <-ctx.Done():
		{
			fmt.Printf("about to shutdown server.")
			if err := h.server.Shutdown(ctx); err != nil {
				return errors.NewWithErr(fmt.Errorf("%v forced to shutdown->err:%w", h.Name(), err))
			}
		}
	}
	return nil
}

func NewServer(viper *viper.Viper) (schema.Server, errors.Error) {
	return &HttpServer{viper: viper}, nil
}
