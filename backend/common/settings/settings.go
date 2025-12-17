package settings

import (
	"github.com/dulisoft/spirit/core/log/zapx"
	"github.com/dulisoft/spirit/core/store/database"
	"sync"

	"github.com/dulisoft/spirit/core/common"
	"github.com/dulisoft/spirit/core/mq"
)

var (
	lock   = new(sync.RWMutex)
	once   = new(sync.Once)
	config *Config
)

func init() {
	once.Do(func() {
		config = new(Config)
	})
}

func GetConfig() *Config {
	lock.RLock()
	defer lock.RUnlock()
	return config
}

func ResetConfig(conf *Config) {
	lock.Lock()
	defer lock.Unlock()
	config = conf
}

// Config 全局配置，这里的配置，应该都公共的
type Config struct {
	Server   common.ServerConf `json:"server"`
	Database database.Options  `json:"database"`
	MQ       mq.MQConf         `json:"mq"`
	zapx.LogConfigs
	DepServices DepServices `json:"depServices"`
}

type DepServices struct {
	UserMgmPrivate   string `json:"userMgmPrivate"`
	HydraAdmin       string `json:"hydraAdmin"`
	CCHost           string `json:"config-center-host"`
	BGHost           string `json:"business-grooming-host"`
	CSHost           string `json:"catalog-service-host"`
	DVHost           string `json:"data-view-host"`
	WorkflowRestHost string `json:"workflowRestHost"`
	DocAuditRestHost string `json:"docAuditRestHost"`
	MQ               *MQ    `json:"mq"`
}

type MQ struct {
	Auth          Auth   `json:"auth"`
	ConnectorType string `json:"connectorType"`
	MqHost        string `json:"mqHost"`
	MqLookupdHost string `json:"mqLookupdHost"`
	MqLookupdPort string `json:"mqLookupdPort"`
	MqPort        string `json:"mqPort"`
	NsqdPortTCP   string `json:"nsqdPortTCP"`
}

type Auth struct {
	Mechanism string `json:"mechanism"`
	Password  string `json:"password"`
	Username  string `json:"username"`
}
