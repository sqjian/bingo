package main

import (
	"github.com/spf13/viper"
	"github.com/sqjian/bingo/pkg/log"
	"github.com/sqjian/bingo/plugin/proto"
	"testing"
)

type PluginTools struct {
}

func (t PluginTools) SendBack(bytes []byte) {
	panic("implement me")
}

func (t PluginTools) GetBackData() []byte {
	panic("implement me")
}

func (t PluginTools) Set(i interface{}, i2 interface{}) {
	panic("implement me")
}

func (t PluginTools) Get(i interface{}) (value interface{}, ok bool) {
	panic("implement me")
}

func (t PluginTools) Logger() log.Logger {
	return log.GetDefLogger()
}


func Test_Print(t *testing.T) {
	plugin, _ := NewPlugin(viper.New())

	msg := &proto.Msg{
		DataList: []*proto.Data{
			{
				Id:   "uuid",
				Desc: map[string][]byte{"key": []byte("val")},
				Data: []byte("data"),
			},
		},
	}

	_, _ = plugin.Process(msg, PluginTools{})
}
