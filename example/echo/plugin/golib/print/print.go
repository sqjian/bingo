package main

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/sqjian/bingo/pkg/errors"
	"github.com/sqjian/bingo/plugin/proto"
	"github.com/sqjian/bingo/plugin/schema"
)

type Print struct {
	viper *viper.Viper
}

func (p *Print) Init() errors.Error {
	fmt.Printf("%v Init...\n", p.Name())
	return nil
}
func (p *Print) FInit() errors.Error {
	fmt.Printf("%v FInit...\n", p.Name())
	return nil
}
func (p *Print) Interest(msg *proto.Msg, tools schema.PluginTools) (schema.Action, errors.Error) {
	return schema.CONTINUE, nil
}

func (p *Print) PreProcess(msg *proto.Msg, tools schema.PluginTools) (schema.Action, errors.Error) {
	return schema.CONTINUE, nil
}

func (p *Print) Process(msg *proto.Msg, tools schema.PluginTools) (schema.Action, errors.Error) {
	data := []byte(fmt.Sprintf("reply msg:%v\n", msg.String()))
	tools.Logger().Infof("print process:%v", string(data))
	tools.SendBack(data)
	return schema.CONTINUE, nil
}

func (p *Print) PostProcess(msg *proto.Msg, tools schema.PluginTools) (schema.Action, errors.Error) {
	return schema.CONTINUE, nil
}

func (p *Print) Name() string {
	return "print"
}
func NewPlugin(viper *viper.Viper) (schema.Plugin, errors.Error) {
	return &Print{viper: viper}, nil
}
