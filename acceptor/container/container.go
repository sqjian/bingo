package container

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/sqjian/bingo/acceptor/schema"
	"github.com/sqjian/bingo/pkg/errors"
	"github.com/sqjian/bingo/pkg/log"
	"sync"
)

type Container struct {
	sync.RWMutex
	viper   *viper.Viper
	servers []schema.Server
}

func NewContainer(viper *viper.Viper) Container {
	return Container{ viper: viper}
}
func (p *Container) Init() errors.Error {
	log.GetDefLogger().Infof("container Init...\n")
	return nil
}
func (p *Container) FInit() errors.Error {
	log.GetDefLogger().Infof("container FInit...\n")
	return nil
}
func (p *Container) Add(servers []schema.Server) errors.Error {
	p.Lock()
	defer p.Unlock()

	checkRepeatAdd := func(serverName string) bool {
		for _, server := range p.servers {
			if server.Name() == serverName {
				return true
			}
		}
		return false
	}

	for _, server := range servers {
		if checkRepeatAdd(server.Name()) {
			return errors.NewWithErr(fmt.Errorf("%v already added", server.Name()))
		}
		p.servers = append(p.servers, server)
	}
	return nil
}

func (p *Container) Remove(pluginName string) errors.Error {
	p.Lock()
	defer p.Unlock()
	for ix, val := range p.servers {
		if val.Name() == pluginName {
			if ix+1 >= len(p.servers) {
				p.servers = p.servers[:ix]
			} else {
				p.servers = append(p.servers[:ix], p.servers[ix+1:]...)
			}
			return p.servers[ix].FInit()
		}
	}
	return nil
}
func (p *Container) Get(serverNames ...string) ([]schema.Server, errors.Error) {
	if len(serverNames) == 0 {
		return nil, errors.NewWithErr(fmt.Errorf("can't get server, please specify serverNames"))
	}

	p.RLock()
	defer p.RUnlock()

	serverMap := func() map[string]schema.Server {
		serverMap := make(map[string]schema.Server)
		for _, server := range p.servers {
			serverMap[server.Name()] = server
		}
		return serverMap
	}()

	servers, serversErr := func() ([]schema.Server, errors.Error) {
		var servers []schema.Server
		for _, serverName := range serverNames {
			server, serverOk := serverMap[serverName]
			if !serverOk {
				return nil, errors.NewWithErr(fmt.Errorf("can not found %v", serverName))
			}
			servers = append(servers, server)
		}
		return servers, nil
	}()

	return servers, serversErr
}
