package native

import (
	"github.com/spf13/viper"
	"github.com/sqjian/bingo/pkg/errors"
	"github.com/sqjian/bingo/pkg/log"
	"github.com/sqjian/bingo/plugin/plugin/enter"
	"github.com/sqjian/bingo/plugin/plugin/leave"
	"github.com/sqjian/bingo/plugin/schema"
	"sync"
)

type Loader struct {
	sync.RWMutex
	viper *viper.Viper
}

func NewLoader(viper *viper.Viper) schema.Loader {
	return &Loader{viper: viper}
}

func (l *Loader) Init() errors.Error {
	return nil
}

func (l *Loader) FInit() errors.Error {
	return nil
}

func (l *Loader) Load(_ ...string) ([]schema.Plugin, errors.Error) {
	l.Lock()
	defer l.Unlock()

	var plugins []schema.Plugin

	{
		{
			plugin, pluginErr := enter.NewPlugin(l.viper)
			if pluginErr != nil {
				return plugins, pluginErr
			}
			plugins = append(plugins, plugin)
		}
		{
			plugin, pluginErr := leave.NewPlugin(l.viper)
			if pluginErr != nil {
				return plugins, pluginErr
			}
			plugins = append(plugins, plugin)
		}
	}
	{
		for _, plugin := range plugins {
			initErr := plugin.Init()
			if initErr != nil {
				log.GetDefLogger().Infof("%v->Init failed:%v\n", plugin.Name(), initErr)
				return nil, initErr
			}
		}
	}
	return plugins, nil
}
