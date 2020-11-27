package leave

import (
	"github.com/spf13/viper"
	"github.com/sqjian/bingo/pkg/errors"
	"github.com/sqjian/bingo/pkg/log"
	"github.com/sqjian/bingo/plugin/proto"
	"github.com/sqjian/bingo/plugin/schema"
)

type leave struct {
	viper *viper.Viper
}

func (p *leave) Init() errors.Error {
	log.GetDefLogger().Infof("%v Init...\n", p.Name())
	return nil
}
func (p *leave) FInit() errors.Error {
	log.GetDefLogger().Infof("%v FInit...\n", p.Name())
	return nil
}
func (p *leave) Interest(msg *proto.Msg, tools schema.PluginTools) (schema.Action, errors.Error) {
	return schema.CONTINUE, nil
}

func (p *leave) PreProcess(msg *proto.Msg, tools schema.PluginTools) (schema.Action, errors.Error) {
	log.GetDefLogger().Infof("leave PreProcess msg:%v\n", msg.String())
	return schema.CONTINUE, nil
}

func (p *leave) Process(msg *proto.Msg, tools schema.PluginTools) (schema.Action, errors.Error) {
	log.GetDefLogger().Infof("leave Process msg:%v\n", msg.String())
	return schema.CONTINUE, nil
}

func (p *leave) PostProcess(msg *proto.Msg, tools schema.PluginTools) (schema.Action, errors.Error) {
	log.GetDefLogger().Infof("leave PostProcess msg:%v\n", msg.String())
	return schema.CONTINUE, nil
}

func (p *leave) Name() string {
	return "leave"
}
func NewPlugin(viper *viper.Viper) (schema.Plugin, errors.Error) {
	return &leave{viper: viper}, nil
}
