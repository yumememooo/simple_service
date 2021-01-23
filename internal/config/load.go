package config

import (
	"flag"
	"io/ioutil"

	"github.com/BurntSushi/toml"
)

var (
	Configuration   *ConfigurationStruct
	ConfigDirectory = "./configs"
	ConfigFileName  = "configuration.toml"
)

func LoadConfig() error {
	var confEnv string
	var confDir string

	flag.StringVar(&confDir, "confdir", "", "Specify config directory ex:configs")
	flag.StringVar(&confEnv, "confenv", "", "Specify a dev/prod.")
	flag.Parse()

	path := determinePath(confDir, confEnv)

	Configuration = &ConfigurationStruct{}
	err := loadFromFile(path, confEnv, Configuration)
	if err != nil {
		return err
	}

	return nil
}

func loadFromFile(path string, env string, configuration interface{}) error {

	contents, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	//將內容綁到configuration上
	err = toml.Unmarshal(contents, configuration)
	if err != nil {
		return err
	}

	return nil
}

func determinePath(confDir string, confEnv string) string {
	path := confDir
	if len(path) == 0 {
		path = ConfigDirectory
	}
	filePathaName := path + "/" + ConfigFileName
	if len(confEnv) > 0 {
		filePathaName = path + "/" + confEnv + "/" + ConfigFileName
	}
	return filePathaName
}
