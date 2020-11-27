package schema

import (
	"github.com/sqjian/bingo/pkg/errors"
	"github.com/sqjian/bingo/plugin/proto"
)

type Action int

const (
	STOP Action = iota
	CONTINUE
	SKIP
)

type Plugin interface {
	Name() string
	Init() errors.Error
	FInit() errors.Error
	Interest(msg *proto.Msg, tools PluginTools) (Action, errors.Error)
	PreProcess(msg *proto.Msg, tools PluginTools) (Action, errors.Error)
	Process(msg *proto.Msg, tools PluginTools) (Action, errors.Error)
	PostProcess(msg *proto.Msg, tools PluginTools) (Action, errors.Error)
}
