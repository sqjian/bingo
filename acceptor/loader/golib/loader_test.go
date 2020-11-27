package golib

import (
	"github.com/spf13/viper"
	"github.com/sqjian/bingo/pkg/log"
	"testing"
)

type Tools struct {
}

func (t Tools) Logger() log.Logger {
	return log.GetDefLogger()
}

func Test_Loader(t *testing.T) {

	viperInst := viper.New()
	viperInst.Set("golib_httpserver.addr", "0.0.0.0")

	loader := NewLoader(viperInst)
	if err := loader.Init(); err != nil {
		t.Fatal(err)
	}

	servers, serversErr := loader.Load("httpserver.so")
	if serversErr != nil {
		t.Fatal(serversErr)
	}

	for ix, server := range servers {
		t.Logf("test->server_%v:%v\n", ix, server.Name())
	}
	if err := loader.FInit(); err != nil {
		t.Fatal(err)
	}
}
