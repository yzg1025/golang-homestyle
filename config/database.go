package config

type JWT struct{
	Signingkey   string
	Expirestime  int64
	Buffertime   int64
}

type MySql struct{
	Name     string
	Host     string
	Username string
	Password string
}

type JUHE struct {
	Juheurl    string
	Apikey     string
	Tplid      int
	Tplvalue   string
}
type ZAP struct {
	Level        string
	format       string
	Director     string
	Linkname     string
	Showline     bool
	Encodelevel  string
	Stacktracekey string
	Loginconsole bool
}

type Oss struct {
	Accesskeyid string
	Accesskeysecret string
	Endpoint    string
	Bucket string
	Uploaddir string
	Ossserve string
}

type Ymal struct {
	Addr     int `mapstructure:"addr"`
	JWT      JWT  `mapstructure:"jwt"`
	MySql    MySql  `mapstructure:"mysql"`
	JUHE     JUHE  `mapstructure:"juhe"`
	ZAP      ZAP  `mapstructure:"zap"`
	Oss      Oss  `mapstructure:"oss"`
}