package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
	"runtime"
)

var (
	_, b, _, _ = runtime.Caller(0)
	// Root folder of this project
	Root = filepath.Join(filepath.Dir(b), "../../..")
)

var Config config

type config struct {
	Env struct {
		AK string `yaml:"ak"`
		SK string `yaml:"sk"`
	} `yaml:"env"`
	WPS struct {
		Addr       string `yaml:"addr"`
		CompanyId  string `yaml:"companyId"`
		PlatformId string `yaml:"platformId"`
	} `yaml:"wps"`
	Log struct {
		StorageLocation       string   `yaml:"storageLocation"`
		RotationTime          int      `yaml:"rotationTime"`
		RemainRotationCount   uint     `yaml:"remainRotationCount"`
		RemainLogLevel        uint     `yaml:"remainLogLevel"`
		ElasticSearchSwitch   bool     `yaml:"elasticSearchSwitch"`
		ElasticSearchAddr     []string `yaml:"elasticSearchAddr"`
		ElasticSearchUser     string   `yaml:"elasticSearchUser"`
		ElasticSearchPassword string   `yaml:"elasticSearchPassword"`
	} `yaml:"log"`
	DecSalt        string `yaml:"decSalt"`
	DecCsvFilePath string `yaml:"decCsvFilePath"`
	Mysql          struct {
		DBAddress      []string `yaml:"dbMysqlAddress"`
		DBUserName     string   `yaml:"dbMysqlUserName"`
		DBPassword     string   `yaml:"dbMysqlPassword"`
		DBDatabaseName string   `yaml:"dbMysqlDatabaseName"`
		DBMsgTableNum  int      `yaml:"dbMsgTableNum"`
		DBMaxOpenConns int      `yaml:"dbMaxOpenConns"`
		DBMaxIdleConns int      `yaml:"dbMaxIdleConns"`
		DBMaxLifeTime  int      `yaml:"dbMaxLifeTime"`
		LogLevel       int      `yaml:"logLevel"`
		SlowThreshold  int      `yaml:"slowThreshold"`
	}
}

func unmarshalConfig(config interface{}, configName string) {
	var env string
	if configName == "config.yaml" {
		env = "CONFIG_NAME"
	} else {
		panic("configName must be config.yaml")
	}
	cfgName := os.Getenv(env)
	if len(cfgName) != 0 {
		bytes, err := os.ReadFile(filepath.Join(cfgName, "config", configName))
		if err != nil {
			bytes, err = os.ReadFile(filepath.Join(Root, "config", configName))
			if err != nil {
				panic(err.Error() + " config: " + filepath.Join(cfgName, "config", configName))
			}
		} else {
			Root = cfgName
		}
		if err = yaml.Unmarshal(bytes, config); err != nil {
			panic(err.Error())
		}
	} else {
		if Config.Env.AK != "" {
			//容器
			bytes, err := os.ReadFile(fmt.Sprintf("%s", configName))
			if err != nil {
				panic(err.Error() + configName)
			}
			if err = yaml.Unmarshal(bytes, config); err != nil {
				panic(err.Error())
			}
		} else {
			//本地
			//bytes, err := os.ReadFile(fmt.Sprintf("./pkg/config/%s", configName))
			bytes, err := os.ReadFile(fmt.Sprintf("../config/%s", configName))
			if err != nil {
				panic(err.Error() + configName)
			}
			if err = yaml.Unmarshal(bytes, config); err != nil {
				panic(err.Error())
			}
		}
	}
}
func init() {
	Config.Env.AK = os.Getenv("AK")
	Config.Env.SK = os.Getenv("SK")
	unmarshalConfig(&Config, "config.yaml")
}
