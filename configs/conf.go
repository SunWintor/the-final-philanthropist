package configs

import (
	"log"
	"time"

	"github.com/BurntSushi/toml"
)

var config Conf

type Conf struct {
	TFPServerPort string
	LoginSalt     string
	JWTSalt       string
	Mysql         *MysqlConf
}

type MysqlConf struct {
	Url          string
	MaxLifetime  time.Duration
	MaxOpenConns int
	MaxIdleConns int
}

func LoadConf() {
	if _, err := toml.DecodeFile("./configs/conf.toml", &config); err != nil {
		log.Fatalln(err)
	}
}

// todo
func GetConf() Conf {
	return config
}
