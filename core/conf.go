package core

import (
	"blog_server/conf"
	"blog_server/global"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

// 初始化，读取配置
func IninCf() {
	const CfFile = "settings.yaml"
	yamlcf, err := ioutil.ReadFile(CfFile)
	c := &conf.Config{}
	if err != nil {
		panic(fmt.Errorf("yam_cf_error: %s", err))
	}
	err = yaml.Unmarshal(yamlcf, c)
	if err != nil {
		log.Fatalf("Cf init unmarshal: %v", err)
	}
	log.Println("Cf init success")
	global.Config = c

}
