package config

import (
	"msg/common/dbx"
)

type Config struct {
	Mysql dbx.DbConfig
}
