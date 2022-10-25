package config

import (
	"github.com/wslky/dbgo/internal/utils"
)

type Config struct {
	// 包名
	PackageName string

	// 数据库连接地址
	DBName     string
	DBUser     string
	DBPassword string
	DBUrl      string

	// 文件 生成/读取 路径
	Path string

	// 命名规则
	FileNameType     NameType // 文件名命名规则
	FilePropertyType NameType // 文件属性命名规则
	JSONType         NameType // json命名规则
}

type NameType func(old string) string

var (
	CameCase       NameType = utils.CamelCaseHelper
	UnderScoreCase NameType = utils.UnderScoreCaseHelper
	LowerCase      NameType = utils.LowerCaseHelper
)

func NameTypeMapping(name string) NameType {
	if name == "goStyle" {
		return CameCase
	} else if name == "go_style" {
		return UnderScoreCase
	} else if name == "gostyle" {
		return LowerCase
	}
	panic("name type error")
	return nil
}
