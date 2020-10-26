package config

import (
	"fmt"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	data := LoadConfig("config.yaml")
	fmt.Println(data.Mysql.User)
	fmt.Println(data.Mysql.DBName)
}
