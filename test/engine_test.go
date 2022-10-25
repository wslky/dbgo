package test

import (
	"github.com/wslky/dbgo/internal"
	"github.com/wslky/dbgo/internal/config"
	"github.com/wslky/dbgo/internal/utils"
	"log"
	"testing"
)

func TestCamelCaseHelper(t *testing.T) {
	str1 := "user_name"
	str2 := "UserName"
	str3 := "userName"
	str4 := "_user_name"
	log.Println(utils.CamelCaseHelper(str1) == "UserName")
	log.Println(utils.CamelCaseHelper(str2) == "UserName")
	log.Println(utils.CamelCaseHelper(str3) == "UserName")
	log.Println(utils.CamelCaseHelper(str4) == "UserName")
}

func TestUnderScoreCaseHelper(t *testing.T) {
	str1 := "user_name"
	str2 := "UserName"
	str3 := "userName"
	str4 := "_user_name"
	log.Println(utils.UnderScoreCaseHelper(str1) == "user_name")
	log.Println(utils.UnderScoreCaseHelper(str2) == "user_name")
	log.Println(utils.UnderScoreCaseHelper(str3) == "user_name")
	log.Println(utils.UnderScoreCaseHelper(str4) == "user_name")
}

func TestOutFile(t *testing.T) {
	conf := config.Config{
		PackageName:      "test",
		DBName:           "test",
		DBUser:           "root",
		DBPassword:       "132456",
		DBUrl:            "127.0.0.1:3306",
		Path:             "",
		FileNameType:     config.LowerCase,
		FilePropertyType: config.LowerCase,
		JSONType:         config.UnderScoreCase,
	}

	e := internal.New(conf)
	e.OutGoFile()
}
