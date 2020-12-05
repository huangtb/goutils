package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"testing"
)

//服务独有的配置文件
var ExpConfig *ExpandConfig

type ExpandConfig struct {
	Switch string `json:"switch" yaml:"switch"`
}

func unmarshalExpConfig(content string) error {
	return yaml.Unmarshal([]byte(content), &ExpConfig)
}


func TestNacos(t *testing.T) {
	localFilePath := "E:\\code\\goproject\\ken-dmp\\goutils\\config\\config-test.yml"
	nacos := &Nacos{}
	nacos.LoadNConfig(localFilePath)
	fmt.Printf("nacos.NConfig.ServerIpAddr[0]>>%v \n", nacos.NConfig.ServerIpAddr[0])
	nacos.NewNClient()
	content, _ := nacos.GetCenterConfig(nacos.NConfig.CommDataId, nacos.NConfig.Group)
	fmt.Printf("公共配置内容content :%s \n", content)

	if err := UnmarshalCommConfig(content); err != nil {
		panic(err)
	}
	fmt.Printf("公共配置 CommConfig :%+v \n", CommConfig.EbkDataMysql.User)
	//cp := nacos.NewDefaultListenConfigParam()
	//commCP := nacos.NewListenConfigParam(nacos.NConfig.CommDataId,nacos.NConfig.Group,UnmarshalCommConfig)
	//nacos.ListenCenterConfig(commCP)
	nacos.ListenCenterConfig(nacos.NConfig.CommDataId,nacos.NConfig.Group,UnmarshalCommConfig)
	nacos.NConfig.ExpDataId = "ken-kyc-config"
	if len(nacos.NConfig.ExpDataId) > 0 {
		//cp.DataId = nacos.NConfig.ExpDataId
		//cp.OnChange = func(namespace, group, dataId, data string) {
		//	unmarshalExpConfig(data)
		//}
		c, _ := nacos.GetCenterConfig(nacos.NConfig.ExpDataId, nacos.NConfig.Group)
		if err := unmarshalExpConfig(c); err != nil {
			panic(err)
		}
		nacos.ListenCenterConfig(nacos.NConfig.ExpDataId,nacos.NConfig.Group,unmarshalExpConfig)
		//expCP := nacos.NewListenConfigParam(nacos.NConfig.ExpDataId,nacos.NConfig.Group,unmarshalExpConfig)
		//nacos.ListenCenterConfig(expCP)
		fmt.Printf("扩展配置 ExpConfig :%+v \n", ExpConfig)
	}

	<- make(chan bool)
}
