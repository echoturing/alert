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

	result, err := config.Evaluates(context.Background(), "select count(1) as eng,100 as 中文 from api_key")
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	fmt.Printf("%#v\n", result[0])
	fmt.Printf("%#v\n", result[1])
}
