package config

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"

	"gorm.io/gorm"
)

var DB *gorm.DB

// DBConfig is a type of DBConfiguration info
// Host     string `yaml:"host"`
// Port     int    `yaml:"port"`
// User     string `yaml:"user"`
// Password string `yaml:"password"`
// DBName   string `yaml:"db_name"`
type DBConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBName   string `yaml:"db_name"`
}

// BuildDBConfig config database
func BuildDBConfig() *DBConfig {
	var dbConfig DBConfig
	file, err := ioutil.ReadFile("./config/database.yml")

	if err != nil {
		log.Panic(err)
	}

	yaml.Unmarshal(file, &dbConfig)

	return &dbConfig
}

// DbURL generate DB URL
func DbURL(dbConfig *DBConfig) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
	)
}
