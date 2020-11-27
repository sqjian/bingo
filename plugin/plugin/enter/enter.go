package enter

import (
	"github.com/spf13/viper"
	"github.com/sqjian/bingo/pkg/errors"
	"github.com/sqjian/bingo/pkg/log"
	"github.com/sqjian/bingo/plugin/proto"
	"github.com/sqjian/bingo/plugin/schema"
)

type enter struct {
	viper *viper.Viper
}

func (p *enter) Init() errors.Error {
	log.GetDefLogger().Infof("%v Init...\n", p.Name())
	return nil
}
func (p *enter) FInit() errors.Error {
	log.GetDefLogger().Infof("%v FInit...\n", p.Name())
	return nil
}
func (p *enter) Interest(msg *proto.Msg, tools schema.PluginTools) (schema.Action, errors.Error) {
	return schema.CONTINUE, nil
}

func (p *enter) PreProcess(msg *proto.Msg, tools schema.PluginTools) (schema.Action, errors.Error) {
	log.GetDefLogger().Infof("enter PreProcess msg:%v\n", msg.String())
	return schema.CONTINUE, nil
}

func (p *enter) Process(msg *proto.Msg, tools schema.PluginTools) (schema.Action, errors.Error) {
	log.GetDefLogger().Infof("enter Process msg:%v\n", msg.String())
	return schema.CONTINUE, nil
}

func (p *enter) PostProcess(msg *proto.Msg, tools schema.PluginTools) (schema.Action, errors.Error) {
	log.GetDefLogger().Infof("enter PostProcess msg:%v\n", msg.String())
	return schema.CONTINUE, nil
}

func (p *enter) Name() string {
	return "enter"
}
func NewPlugin(viper *viper.Viper) (schema.Plugin, errors.Error) {
	return &enter{viper: viper}, nil
}
