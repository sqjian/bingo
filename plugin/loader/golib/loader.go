package golib

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/sqjian/bingo/pkg/errors"
	"github.com/sqjian/bingo/pkg/log"
	"github.com/sqjian/bingo/plugin/schema"
	"plugin"
	"strings"
	"sync"
)

type Loader struct {
	sync.RWMutex
	viper *viper.Viper
}

func NewLoader(viper *viper.Viper) schema.Loader {
	return &Loader{ viper: viper}
}

func (l *Loader) Init() errors.Error {
	return nil
}

func (l *Loader) FInit() errors.Error {
	return nil
}
func (l *Loader) reWritePluginName(pluginName string) string {
	if strings.HasSuffix(pluginName, ".so") {
		return pluginName
	}
	return pluginName + ".so"
}
func (l *Loader) Load(pluginNames ...string) ([]schema.Plugin, errors.Error) {
	if len(pluginNames) == 0 || func() bool {
		for _, pluginName := range pluginNames {
			if len(pluginName) == 0 {
				return true
			}
		}
		return false
	}() {
		return nil, errors.NewWithErr(fmt.Errorf("can't load plugin, please specify pluginNames"))
	}

	l.Lock()
	defer l.Unlock()

	var plugins []schema.Plugin
	for _, pluginName := range pluginNames {
		plugin, pluginErr := l.verify(l.reWritePluginName(pluginName))
		if pluginErr != nil {
			return nil, pluginErr
		}
		plugins = append(plugins, plugin)
	}
	return plugins, nil
}

func (l *Loader) verify(pluginName string) (schema.Plugin, errors.Error) {
	log.GetDefLogger().Infof("verifying plugin: %v\n", pluginName)

	pg, err := plugin.Open(pluginName)
	if err != nil {
		log.GetDefLogger().Infof("can not open plugin: %v\n", pluginName)
		return nil, errors.NewWithErr(err)
	}

	f, err := pg.Lookup("NewPlugin")
	if err != nil {
		log.GetDefLogger().Infof("Lookup plugin: %v failed\n", pluginName)
		return nil, errors.NewWithErr(err)
	}
	inst, instErr := func() (schema.Plugin, errors.Error) {
		NewPlugin, NewPluginOk := f.(schema.NewPlugin)
		if !NewPluginOk {
			log.GetDefLogger().Infof("%v->convert %T to %T failed\n", pluginName, f, schema.NewPluginObj)
			return nil, errors.PluginMethodNotFound
		}
		plugin, pluginErr := NewPlugin(l.viper)

		if pluginErr != nil {
			log.GetDefLogger().Infof("%v->NewPlugin failed:%v\n", pluginName, pluginErr)
			return nil, pluginErr
		}
		initErr := plugin.Init()
		if initErr != nil {
			log.GetDefLogger().Infof("%v->Init failed:%v\n", pluginName, initErr)
			return nil, initErr
		}
		return plugin, nil
	}()
	return inst, instErr
}
