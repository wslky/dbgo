package config

import (
	"github.com/wslky/dbgo/internal/utils"
	"log"
)

type Config struct {
	PackageName string

	DBName     string
	DBUser     string
	DBPassword string
	DBUrl      string

	FileNameType     NameType
	FilePropertyType NameType
	JSONType         NameType
}

type NameType func(old string) string

var (
	CameCase       NameType = utils.CamelCaseHelper
	UnderScoreCase NameType = utils.UnderScoreCaseHelper
	LowerCase      NameType = utils.LowerCaseHelper
)

// NameTypeMapping Returns a naming convention based on a string
func NameTypeMapping(name string) NameType {
	if name == "goStyle" {
		return CameCase
	} else if name == "go_style" {
		return UnderScoreCase
	} else if name == "gostyle" {
		return LowerCase
	}
	log.Fatal("name type error")
	return nil
}
