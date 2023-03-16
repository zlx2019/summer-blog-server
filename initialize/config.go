package initialize

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	. "summer/constant"
	"summer/properties"
)

// InitLoadConfigure 解析配置文件,加载到实体中
func InitLoadConfigure() {
	// 解析后要映射到的对象
	conf := &properties.Config{}
	// 读取文件流
	file, err := ioutil.ReadFile(ConfigFile)
	if err != nil {
		panic(fmt.Errorf("load %s error: %s", ConfigFile, err))
	}
	// 将配置文件内容,解析到对象中
	err = yaml.Unmarshal(file, conf)
	if err != nil {
		log.Fatalf("parse %s error: %s", ConfigFile, err)
	}
	// 放到全局变量中
	Config = conf
}
