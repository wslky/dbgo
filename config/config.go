package config

import (
	"github.com/wslky/dbgo/utils"
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
	FileNameType     NameType
	FilePropertyType NameType
	JSONType         NameType
}

type NameType func(old string) string

var (
	CameCase       NameType = utils.CamelCaseHelper
	UnderScoreCase NameType = utils.UnderScoreCaseHelper
)
