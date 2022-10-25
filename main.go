package main

import (
	"flag"
	"github.com/wslky/dbgo/internal"
	"github.com/wslky/dbgo/internal/config"
	"log"
)

var username = flag.String("u", "root", "database username")
var password = flag.String("p", "123456", "database password")
var dbName = flag.String("d", "bean", "database password")
var url = flag.String("url", "127.0.0.1:3306", "database url")
var packageName = flag.String("pn", "bean", "go file package name")
var fileNameType = flag.String("fn", "gostyle", "file name type")
var jSONType = flag.String("jt", "go_style", "json type")

func main() {

	log.SetPrefix("[dbgo]")

	flag.Parse()
	conf := config.Config{
		PackageName:      *packageName,
		DBName:           *dbName,
		DBUser:           *username,
		DBPassword:       *password,
		DBUrl:            *url,
		Path:             "",
		FileNameType:     config.NameTypeMapping(*fileNameType),
		FilePropertyType: config.CameCase,
		JSONType:         config.NameTypeMapping(*jSONType),
	}

	e := internal.New(conf)
	e.OutGoFile()
}
