package native

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/sqjian/bingo/acceptor/schema"
	"github.com/sqjian/bingo/acceptor/server/cmdserver"
	"github.com/sqjian/bingo/pkg/errors"
	"github.com/sqjian/bingo/pkg/log"
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

func (l *Loader) Load(serverNames ...string) ([]schema.Server, errors.Error) {
	if len(serverNames) == 0 {
		return nil, errors.NewWithErr(fmt.Errorf("can't load server, please specify serverNames"))
	}

	l.Lock()
	defer l.Unlock()

	var servers []schema.Server

	{
		server, serverErr := cmdserver.NewServer(l.viper)
		if serverErr != nil {
			return servers, serverErr
		}
		servers = append(servers, server)
	}
	{
		serverMap := func() map[string]schema.Server {
			m := make(map[string]schema.Server)
			for _, server := range servers {
				m[server.Name()] = server
			}
			return m
		}()
		for _, serverName := range serverNames {
			server, serverOk := serverMap[serverName]
			if !serverOk {
				return nil, errors.NewWithErr(fmt.Errorf("can not found %v", serverName))
			}
			initErr := server.Init()
			if initErr != nil {
				log.GetDefLogger().Infof("%v->Init failed:%v\n", server.Name(), initErr)
				return nil, initErr
			}
		}
	}
	return servers, nil
}
