package config

import (
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Nacos struct {
	NConfig NConfig `json:"nacos" yaml:"nacos"`
	NClient
}
type NConfig struct {
	ServerIpAddr []string `json:"server_ip_addr" yaml:"server_ip_addr"`
	Port         uint64   `json:"port" yaml:"port"`
	NamespaceId  string   `json:"namespace_id" yaml:"namespace_id"`
	CommDataId   string   `json:"comm_data_id" yaml:"comm_data_id"`
	ExpDataId    string   `json:"exp_data_id" yaml:"exp_data_id"`
	Group        string   `json:"group" yaml:"group"`
	LogDir       string   `json:"log_dir" yaml:"log_dir"`
	CacheDir     string   `json:"cache_dir" yaml:"cache_dir"`
	LogLevel     string   `json:"log_level" yaml:"log_level"`
}

type NClient struct {
	Client config_client.IConfigClient
}

type unmarshalFunc func(arg string) error


func (n *Nacos) LoadNConfig(path string) error {
	body, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf("Read config file error:%s", err.Error())
		panic(err)
	}
	fmt.Printf("body>>%s", string(body))
	return yaml.Unmarshal(body, &n)
}

func (n *Nacos) NewNClient() {
	var scs []constant.ServerConfig
	for i := 0; i < len(n.NConfig.ServerIpAddr); i++ {
		sc := constant.ServerConfig{
			IpAddr: n.NConfig.ServerIpAddr[i],
			Port:   n.NConfig.Port,
		}
		scs = append(scs, sc)
	}

	cc := constant.ClientConfig{
		NamespaceId:         n.NConfig.NamespaceId, //namespace id
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              n.NConfig.LogDir,
		CacheDir:            n.NConfig.CacheDir,
		RotateTime:          "1h",
		MaxAge:              3,
		LogLevel:            n.NConfig.LogLevel,
	}

	client, err := clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": scs,
		"clientConfig":  cc,
	})

	if err != nil {
		panic(err)
	}
	n.Client = client
}

func (n *Nacos) GetCenterConfig(dataId, group string) (string, error) {
	content, err := n.Client.GetConfig(vo.ConfigParam{
		DataId: dataId,
		Group:  group,
	})
	if err != nil {
		return "", err
	}
	return content, nil
}

func (n *Nacos) ListenDefaultCenterConfig() error {
	err := n.Client.ListenConfig(vo.ConfigParam{
		DataId: n.NConfig.CommDataId,
		Group:  n.NConfig.Group,
		OnChange: func(namespace, group, dataId, data string) {
			fmt.Println("center config changed\n group:" + group + ", dataId:" + dataId + ", content:" + data)
			UnmarshalCommConfig(data)
		},
	})
	return err
}

//func (n *Nacos) NewDefaultListenConfigParam() vo.ConfigParam {
//	return vo.ConfigParam{
//		DataId: n.NConfig.CommDataId,
//		Group:  n.NConfig.Group,
//		OnChange: func(namespace, group, dataId, data string) {
//			fmt.Println("center config changed\n group:" + group + ", dataId:" + dataId + ", content:" + data)
//			UnmarshalCommConfig(data)
//		},
//	}
//}


func (n *Nacos) NewListenConfigParam(dataId,group string,uf unmarshalFunc) vo.ConfigParam {
	return vo.ConfigParam{
		DataId: dataId,
		Group:  group,
		OnChange: func(namespace, group, dataId, data string) {
			fmt.Println("center config changed\n group:" + group + ", dataId:" + dataId + ", content:" + data)
			uf(data)
		},
	}
}

func (n *Nacos) ListenByConfigParam(param vo.ConfigParam) error {
	err := n.Client.ListenConfig(param)
	return err
}


func (n *Nacos) ListenCenterConfig(dataId,group string,uf unmarshalFunc) error {
	err := n.Client.ListenConfig(vo.ConfigParam{
		DataId: dataId,
		Group:  group,
		OnChange: func(namespace, group, dataId, data string) {
			fmt.Println("center config changed\n group:" + group + ", dataId:" + dataId + ", content:" + data)
			uf(data)
		},
	})
	return err
}
