package schema

import (
	"github.com/spf13/viper"
	"github.com/sqjian/bingo/pkg/errors"
)

var NewPluginObj NewPlugin

type NewPlugin = func(*viper.Viper) (Plugin, errors.Error)

type Loader interface {
	Init() errors.Error
	FInit() errors.Error
	Load(...string) ([]Plugin, errors.Error)
}
