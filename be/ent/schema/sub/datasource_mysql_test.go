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

	result, err := config.Evaluates(context.Background(), "select * from api_key")
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	for _, i := range result {
		fmt.Printf("%s\n", i)
	}
}
