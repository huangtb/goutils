package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"time"
)

type Dao struct {
	Client *gorm.DB
}

type Mysql struct {
	User      string `json:"user" yaml:"user"`
	Pwd       string `json:"pwd" yaml:"pwd"`
	Host      string `json:"host" yaml:"host"`
	Port      int    `json:"port" yaml:"port"`
	DefaultDB string `json:"default_db" yaml:"default_db"`
}

func (m *Mysql) NewDao() (d Dao, err error) {
	url := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=True&loc=Local",
				m.User,m.Pwd,m.Host, m.Port, m.DefaultDB)
	fmt.Printf("d1:%+v,d1地址:%p\n", d, &d)

	if d.Client, err = gorm.Open("db", url); err != nil {
		return d, err
	}
	d.Client.SingularTable(true)       //表名采用单数形式
	d.Client.DB().SetMaxOpenConns(100) 	//SetMaxOpenConns用于设置最大打开的连接数
	d.Client.DB().SetMaxIdleConns(10)  	//SetMaxIdleConns用于设置闲置的连接数
	d.Client.DB().SetConnMaxLifetime(30 * time.Minute)

	d.Client.LogMode(true) 	//启用Logger，显示详细日志

	if err = d.Client.Set("gorm:table_options", "ENGINE=InnoDB").Error; err != nil {
		_ = d.Client.Close()
		return d, err
	}

	return d, nil
}

func (d *Dao) Ping() error {
	return d.Client.DB().Ping()
}

func (d *Dao) Disconnect() error {
	return d.Client.DB().Close()
}
