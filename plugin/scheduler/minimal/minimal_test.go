package minimal_test

import (
	"github.com/spf13/viper"
	"github.com/sqjian/bingo/pkg/log"
	"github.com/sqjian/bingo/plugin/proto"
	"github.com/sqjian/bingo/plugin/scheduler/minimal"
	"testing"
)

type PluginTools struct {
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

func (t PluginTools) SendBack(_ []byte) {
	panic("implement me")
}

func (t PluginTools) GetBackData() []byte {
	panic("implement me")
}

func Test_minimal(t *testing.T) {
	viperInst := viper.New()
	viperInst.Set("plugin.golib", []string{"print"})

	minimalInst, minimalInstErr := minimal.NewMinimal(viperInst)
	if minimalInstErr != nil {
		t.Fatal(minimalInstErr)
	}

	msg := &proto.Msg{
		DataList: []*proto.Data{
			{
				Id:   "uuid",
				Desc: map[string][]byte{"key": []byte("val")},
				Data: []byte("data"),
			},
		},
	}

	_, processErr := minimalInst.Process(minimal.Dag{Id: "xxx", Steps: []string{"print"}}, msg)
	if processErr != nil {
		t.Fatal(processErr)
	}
}
