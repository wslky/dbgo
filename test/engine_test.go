package test

import (
	"github.com/wslky/dbgo"
	"github.com/wslky/dbgo/config"
	"testing"
)

var c = config.Config{
	DBName:           "test",
	DBUser:           "root",
	DBPassword:       "123456",
	DBUrl:            "127.0.0.1:3306",
	Path:             "",
	PackageName:      "entity",
	FileNameType:     config.CameCase,
	FilePropertyType: config.CameCase,
	JSONType:         config.UnderScoreCase,
}

func TestOutGoFile(t *testing.T) {
	e := dbgo.New(c)
	e.OutGoFile()
}
