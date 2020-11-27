package container_test

import (
	"github.com/spf13/viper"
	"github.com/sqjian/bingo/acceptor/container"
	"github.com/sqjian/bingo/acceptor/loader/golib"
	"github.com/sqjian/bingo/acceptor/loader/native"
	"github.com/sqjian/bingo/pkg/log"
	"testing"
)

type Tools struct {
}

func (t Tools) Logger() log.Logger {
	return log.GetDefLogger()
}

func Test_Container(t *testing.T) {

	viperInst := viper.New()
	viperInst.Set("golib_httpserver.addr", ":8081")

	container := container.NewContainer(viperInst)
	if err := container.Init(); err != nil {
		t.Fatal(err)
	}

	{
		{
			loader := golib.NewLoader(viperInst)
			if err := loader.Init(); err != nil {
				t.Fatal(err)
			}
			servers, serversErr := loader.Load("httpserver.so")
			if serversErr != nil {
				t.Fatal(serversErr)
			}

			addErr := container.Add(servers)
			if addErr != nil {
				t.Fatal(addErr)
			}
		}
		{
			loader := native.NewLoader(viperInst)
			if err := loader.Init(); err != nil {
				t.Fatal(err)
			}
			servers, serversErr := loader.Load()
			if serversErr != nil {
				t.Fatal(serversErr)
			}

			addErr := container.Add(servers)
			if addErr != nil {
				t.Fatal(addErr)
			}
		}
	}

	servers, serversErr := container.Get("golib_httpserver", "native_httpserver")
	if serversErr != nil {
		t.Fatal(serversErr)
	}
	for ix, server := range servers {
		t.Logf("test->server_%v:%v\n", ix, server.Name())
	}

	if err := container.FInit(); err != nil {
		t.Fatal(err)
	}
}
