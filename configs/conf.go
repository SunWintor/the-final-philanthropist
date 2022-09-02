package configs

import "time"

var config Conf

type Conf struct {
	TFPServerPort string
	LoginSalt     string
	Mysql         *MysqlConf
}

type MysqlConf struct {
	Url          string
	MaxLifetime  time.Duration
	MaxOpenConns int
	MaxIdleConns int
}

func LoadConf() {
	config = Conf{
		TFPServerPort: "9000",
		LoginSalt:     "盐",
	}
}

// todo
func GetConf() Conf {
	return config
}
