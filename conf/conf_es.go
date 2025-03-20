package conf

import (
	"fmt"
)

type ES struct {
	//带debug内容设置，设置debug全输出
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

func (es ES) URL() string {
	return fmt.Sprintf("%s:%d", es.Host, es.Port)
}
