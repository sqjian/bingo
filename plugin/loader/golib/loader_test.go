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
	loader := NewLoader(viper.New())
	if err := loader.Init(); err != nil {
		t.Fatal(err)
	}
	plugins, pluginsErr := loader.Load("print.so")
	if pluginsErr != nil {
		t.Fatal(pluginsErr)
	}
	for ix, plugin := range plugins {
		t.Logf("test->plugin_%v:%v\n", ix, plugin.Name())
	}
	if err := loader.FInit(); err != nil {
		t.Fatal(err)
	}
}
