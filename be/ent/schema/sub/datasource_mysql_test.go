package sub

import (
	"context"
	"fmt"
	"testing"
)

func TestMySQLConfig_EvalScript(t *testing.T) {
	config := MySQLConfig{
		User:     "root",
		Password: "123456",
		Host:     "127.0.0.1",
		Port:     3306,
		DBName:   "alert",
	}

	result, err := config.Evaluates(context.Background(), "SELECT CONVERT(12.9,DECIMAL(10,2))")
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	for _, i := range result {
		fmt.Printf("%s\n", i)
	}
}
