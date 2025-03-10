package conf

type Mysql struct {
	//带debug内容设置，设置debug全输出
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Db       string `yaml:"db"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Loglevel string `yaml:"loglevel"`
}
