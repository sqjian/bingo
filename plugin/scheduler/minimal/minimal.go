package minimal

import (
	"github.com/spf13/viper"
	"github.com/sqjian/bingo/pkg/log"
	"github.com/sqjian/bingo/plugin/container"
	"github.com/sqjian/bingo/plugin/loader/golib"
	"github.com/sqjian/bingo/plugin/loader/native"
	"github.com/sqjian/bingo/plugin/proto"
	"github.com/sqjian/bingo/plugin/schema"
	"github.com/sqjian/bingo/plugin/tool"
)

func NewMinimal(viper *viper.Viper) (*Minimal, error) {
	minimal := &Minimal{viper: viper}
	if err := minimal.Init(); err != nil {
		return nil, err
	}
	return minimal, nil
}

type Minimal struct {
	viper     *viper.Viper
	container container.Container
}

func (m *Minimal) FInit() error {
	return m.container.FInit()
}

func (m *Minimal) Process(dag Dag, msg *proto.Msg, schedulerOpts ...schedulerOpt) ([]byte, error) {
	var pluginTools schema.PluginTools = &tool.PluginToolsImpl{}

	schedulerOptTmp := NewSchedulerOpts()
	for _, schedulerOpt := range schedulerOpts {
		schedulerOpt(schedulerOptTmp)
	}
	for key, val := range schedulerOptTmp.kvs {
		pluginTools.Set(key, val)
	}
	pluginTools.Logger().Infof("dag:%v", dag)
	steps := func() []string {
		steps := append([]string{"enter"}, dag.Steps...)
		steps = append(steps, "leave")
		return steps
	}()

	plugins, pluginsErr := m.container.Get(steps...)
	if pluginsErr != nil {
		pluginTools.Logger().Errorf("container.get failed,err:%v", pluginsErr)
		return nil, pluginsErr
	}
	for _, plugin := range plugins {
		pluginTools.Logger().Infof("checking if plugin:%v of dag:%v is interested in", plugin.Name(), dag)
		action, actionErr := plugin.Interest(msg, pluginTools)
		if actionErr != nil {
			pluginTools.Logger().Infof("checking if plugin:%v of dag:%v is interested in failed:%v", plugin.Name(), dag, actionErr)
			return nil, actionErr
		}
		if action == schema.SKIP {
			pluginTools.Logger().Warnf("plugin:%v not interested in,skip", plugin.Name())
			continue
		}
		pluginTools.Logger().Infof("plugin:%v,about to PreProcessing...", plugin.Name())
		_, preProcessErr := plugin.PreProcess(msg, pluginTools)
		if preProcessErr != nil {
			pluginTools.Logger().Errorf("plugin:%v,preProcess failed:%v", plugin.Name(), preProcessErr)
			return nil, preProcessErr
		}
		pluginTools.Logger().Infof("plugin:%v,about to Process...", plugin.Name())
		_, processErr := plugin.Process(msg, pluginTools)
		if processErr != nil {
			pluginTools.Logger().Errorf("plugin:%v,Process failed:%v", plugin.Name(), processErr)
			return nil, processErr
		}
		pluginTools.Logger().Infof("plugin:%v,about to PostProcess...", plugin.Name())
		_, postProcessErr := plugin.PostProcess(msg, pluginTools)
		if postProcessErr != nil {
			pluginTools.Logger().Errorf("plugin:%v,PostProcess failed:%v", plugin.Name(), postProcessErr)
			return nil, postProcessErr
		}
	}
	return pluginTools.GetBackData(), nil
}
func (m *Minimal) Init() error {

	m.container = container.NewContainer(m.viper)
	if err := m.container.Init(); err != nil {
		return err
	}

	// Init golib plugin
	{
		log.GetDefLogger().Infof("golib.NewLoader...")
		loader := golib.NewLoader(m.viper)
		if err := loader.Init(); err != nil {
			log.GetDefLogger().Errorf("golib.loader.Init failed,err:%w", err)
			return err
		}

		golibPlugins := m.viper.GetStringSlice("Plugin.golib")
		log.GetDefLogger().Infof("about to deal golibPlugin:%v", golibPlugins)
		if len(golibPlugins) == 0 {
			log.GetDefLogger().Infof("empty golibPlugins.")
		} else {
			log.GetDefLogger().Infof("about to load %v", golibPlugins)
			plugins, pluginsErr := loader.Load(golibPlugins...)
			if pluginsErr != nil {
				log.GetDefLogger().Errorf("golib.loader.Load failed,pluginsErr:%w", pluginsErr)
				return pluginsErr
			}

			addErr := m.container.Add(plugins)
			if addErr != nil {
				log.GetDefLogger().Errorf("golib.loader.Add failed,addErr:%w", addErr)
				return addErr
			}
		}
	}
	// Init native plugin
	{
		log.GetDefLogger().Infof("native.NewLoader...")
		loader := native.NewLoader(m.viper)
		if err := loader.Init(); err != nil {
			log.GetDefLogger().Errorf("native.loader.Init failed,err:%w", err)
			return err
		}

		log.GetDefLogger().Infof("about to load all native plugin...")
		plugins, pluginsErr := loader.Load()
		if pluginsErr != nil {
			log.GetDefLogger().Errorf("native.loader.Load failed,pluginsErr:%w", pluginsErr)
			return pluginsErr
		}

		addErr := m.container.Add(plugins)
		if addErr != nil {
			log.GetDefLogger().Errorf("native.loader.Add failed,addErr:%w", addErr)
			return addErr
		}
	}
	return nil
}
