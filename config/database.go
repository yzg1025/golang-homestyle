package config

type JWT struct{
	Signingkey string
	Expirestime int64
	Buffertime int64
}

type MySql struct{
	Name string
	Host string
	Username string
	Password string
}

type Ymal struct {
	Addr int `mapstructure:"addr"`
	JWT  JWT  `mapstructure:"jwt"`
	MySql MySql  `mapstructure:"mysql"`
}