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

	result, err := config.Evaluates(context.Background(), "select '2020-01-01' as name,count(1) as eng,100 as 中文 from api_key union all select '2020-02-02' as name ,100 as eng,999 as 中文 from api_key")
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	for _, i := range result {
		fmt.Printf("%#v\n", i)
	}
}
