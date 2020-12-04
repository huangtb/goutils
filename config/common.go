package config

import (
	"github.com/huangtb/go-utils/aws"
	"github.com/huangtb/go-utils/mq"
	"github.com/huangtb/go-utils/mysql"
	"gopkg.in/yaml.v2"
)

var CommConfig *CommonConfig

type CommonConfig struct {
	EbkDataMysql mysql.Mysql  `json:"ebk_data_mysql" yaml:"ebk_data_mysql"`
	EbkCoreMysql mysql.Mysql  `json:"ebk_core_mysql" yaml:"ebk_core_mysql"`
	StatConfig   StatConfig   `json:"stat_config" yaml:"stat_config"`
	Redis        Redis        `json:"redis" yaml:"redis"`
	AwsConfig    aws.Aws      `json:"aws_config" yaml:"aws_config"`
	NsqConfig    mq.NsqConfig `json:"nsq_config" yaml:"nsq_config"`
}

type Redis struct {
	RedisAddr    string `json:"redis_addr" yaml:"redis_addr"`
	RedisDB      int    `json:"redis_db" yaml:"redis_db"`
	TokenExpired int    `json:"token_expired" yaml:"token_expired"`
}

type StatConfig struct {
	StatEnv     string `json:"StatEnv" yaml:"stat_env"`
	StatUdpHost string `json:"StatUdpHost" yaml:"stat_udp_host"`
	StatLogPath string `json:"StatLogPath" yaml:"stat_log_path"`
}

func UnmarshalCommConfig(content string) error {
	return yaml.Unmarshal([]byte(content), &CommConfig)
}
