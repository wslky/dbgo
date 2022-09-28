package test

import (
	"github.com/wslky/dbgo"
	"github.com/wslky/dbgo/config"
	"testing"
)

func TestOutFile(t *testing.T) {
	conf := config.Config{
		PackageName:      "bean",
		DBName:           "test",
		DBUser:           "root",
		DBPassword:       "123456",
		DBUrl:            "127.0.0.1:3306",
		Path:             "",
		FileNameType:     config.UnderScoreCase,
		FilePropertyType: config.CameCase,
		JSONType:         config.UnderScoreCase,
	}

	e := dbgo.New(conf)
	e.OutGoFile()
}
