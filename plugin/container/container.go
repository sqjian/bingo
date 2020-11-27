package container

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/sqjian/bingo/pkg/errors"
	"github.com/sqjian/bingo/plugin/schema"
	"sync"
)

type Container struct {
	sync.RWMutex
	viper   *viper.Viper
	plugins []schema.Plugin
}

func NewContainer(viper *viper.Viper) Container {
	return Container{viper: viper}
}
func (p *Container) Init() errors.Error {
	return nil
}
func (p *Container) FInit() errors.Error {
	return nil
}
func (p *Container) Add(plugins []schema.Plugin) errors.Error {
	p.Lock()
	defer p.Unlock()

	checkRepeatAdd := func(pluginName string) bool {
		for _, plugin := range p.plugins {
			if plugin.Name() == pluginName {
				return true
			}
		}
		return false
	}
	for _, plugin := range plugins {
		if checkRepeatAdd(plugin.Name()) {
			return errors.NewWithErr(fmt.Errorf("%v already added", plugin.Name()))
		}
		p.plugins = append(p.plugins, plugin)
	}

	return nil
}

func (p *Container) Remove(pluginName string) errors.Error {
	p.Lock()
	defer p.Unlock()
	for ix, val := range p.plugins {
		if val.Name() == pluginName {
			if ix+1 >= len(p.plugins) {
				p.plugins = p.plugins[:ix]
			} else {
				p.plugins = append(p.plugins[:ix], p.plugins[ix+1:]...)
			}
			return p.plugins[ix].FInit()
		}
	}
	return nil
}
func (p *Container) Get(pluginNames ...string) ([]schema.Plugin, errors.Error) {
	if len(pluginNames) == 0 {
		return nil, errors.NewWithErr(fmt.Errorf("please specify pluginNames"))
	}

	p.RLock()
	defer p.RUnlock()

	pluginMap := func() map[string]schema.Plugin {
		pluginMap := make(map[string]schema.Plugin)
		for _, plugin := range p.plugins {
			pluginMap[plugin.Name()] = plugin
		}
		return pluginMap
	}()

	plugins, pluginsErr := func() ([]schema.Plugin, errors.Error) {
		var plugins []schema.Plugin
		for _, pluginName := range pluginNames {
			plugin, pluginOk := pluginMap[pluginName]
			if !pluginOk {
				return nil, errors.NewWithErr(fmt.Errorf("can not found %v", pluginName))
			}
			plugins = append(plugins, plugin)
		}
		return plugins, nil
	}()

	return plugins, pluginsErr
}
