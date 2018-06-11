package dbconf

import (
	"io/ioutil"
	"log"

	yaml "gopkg.in/yaml.v2"
)

// DbConf struct defines how db_config.yaml should be built
// dbChoice is only used in interactive applications where one wants to choose the database
type DbConf struct {
	DbChoice string `yaml:"dbChoice"`
	ScUser   string `yaml:"scUser"`
	ScPw     string `yaml:"scPw"`
	ScHost   string `yaml:"scHost"`
	ScDb     string `yaml:"scDb"`
	OmsUser  string `yaml:"omsUser"`
	OmsPw    string `yaml:"omsPw"`
	OmsHost  string `yaml:"omsHost"`
	OmsDb    string `yaml:"omsDb"`
	BobUser  string `yaml:"bobUser"`
	BobPw    string `yaml:"bobPw"`
	BobHost  string `yaml:"bobHost"`
	BobDb    string `yaml:"bobDb"`
	BiUser   string `yaml:"biUser"`
	BiPw     string `yaml:"biPw"`
	BiHost   string `yaml:"biHost"`
	BiDb     string `yaml:"biDb"`
	BaaUser  string `yaml:"baaUser"`
	BaaPw    string `yaml:"baaPw"`
	BaaHost  string `yaml:"baaHost"`
	BaaDb    string `yaml:"baaDb"`
}

// ReadYamlDbConf reads the configuration of the databases
// config.yaml should be in the same folder as the executable
func (c *DbConf) ReadYamlDbConf() *DbConf {

	yamlFile, err := ioutil.ReadFile("db_config.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return c
}
