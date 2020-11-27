package schema

import (
	"github.com/sqjian/bingo/pkg/log"
)

type Buffer interface {
	SendBack([]byte)
	GetBackData() []byte
}
type Logger interface {
	Logger() log.Logger
}
type Carpool interface {
	Set(interface{}, interface{})
	Get(interface{}) (value interface{}, ok bool)
}
type PluginTools interface {
	Buffer
	Logger
	Carpool
}
