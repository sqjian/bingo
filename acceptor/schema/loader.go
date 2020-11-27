package schema

import (
	"github.com/spf13/viper"
	"github.com/sqjian/bingo/pkg/errors"
)

var NewServerObj NewServer

type NewServer = func(*viper.Viper) (Server, errors.Error)

type Loader interface {
	Init() errors.Error
	FInit() errors.Error
	Load(...string) ([]Server, errors.Error)
}
