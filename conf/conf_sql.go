package conf

import "strconv"

type Mysql struct {
	//带debug内容设置，设置debug全输出
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Db       string `yaml:"db"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Loglevel string `yaml:"log_level"`
	Config   string `yaml:"config"` //高级cf
}

func (m Mysql) Dsn() string {
	return m.User + ":" + m.Password + "@tcp(" + m.Host + ":" + strconv.Itoa(m.Port) + ")/" + m.Db + "?" + m.Config
}
