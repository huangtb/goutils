package config

import (
	"github.com/huangtb/goutils/aws"
	"github.com/huangtb/goutils/mq"
	"github.com/huangtb/goutils/mysql"
	"gopkg.in/yaml.v2"
)

var Common *CommonConfig

type CommonConfig struct {
	Country      string       `json:"country" yaml:"country"`
	EbkDataMysql mysql.Mysql  `json:"ebk_data_mysql" yaml:"ebk_data_mysql"`
	EbkCoreMysql mysql.Mysql  `json:"ebk_core_mysql" yaml:"ebk_core_mysql"`
	Stat         StatConfig   `json:"stat" yaml:"stat"`
	Redis        Redis        `json:"redis" yaml:"redis"`
	Aws          aws.Aws      `json:"aws" yaml:"aws"`
	Nsq          mq.NsqConfig `json:"nsq" yaml:"nsq"`
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
	return yaml.Unmarshal([]byte(content), &Common)
}
