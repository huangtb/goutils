package mysql

import (
	"testing"
)

var dmpDao *Dao

func GetDmpDao() *Dao {
	return dmpDao
}

const (
	User = ""
	Pwd  = ""
	Host = ""
)

func TestNewDao(t *testing.T) {
	dmpMysql := Mysql{
		User:      User,
		Pwd:       Pwd,
		Host:      Host,
		Port:      3306,
		DefaultDB: "",
	}

	_, err := dmpMysql.NewDao()
	if err != nil {
		panic(err)
	}

}
