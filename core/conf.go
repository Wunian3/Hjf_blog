package core

import (
	"blog_server/conf"
	"blog_server/global"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/fs"
	"io/ioutil"
	"log"
)

const CfFile = "settings.yaml"

// 初始化，读取配置
func IninCf() {

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

func SetYaml() error {
	byteData, err := yaml.Marshal(global.Config)
	if err != nil {
		global.Log.Error(err)
		return err
	}
	err = ioutil.WriteFile(CfFile, byteData, fs.ModePerm)
	if err != nil {
		global.Log.Error(err)
		return err
	}
	global.Log.Info("Cf init success")
	return nil
}
