package configs

type Conf struct {
	TFPServerPort int64
	LoginSalt     string
}

// todo
func GetConf() Conf {
	return Conf{}
}
