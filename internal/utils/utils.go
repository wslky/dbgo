package utils

import (
	"fmt"
	"strings"
)

// GetSql4TableDef return sql for get table info
func GetSql4TableDef(dbName, tableName string) string {
	return fmt.Sprintf("SELECT COLUMN_NAME AS 'ColumnName', DATA_TYPE AS 'DataType' "+
		"FROM information_schema.COLUMNS WHERE TABLE_SCHEMA = '%v' AND TABLE_NAME ='%v'", dbName, tableName)
}

// CamelCaseHelper GoStyle
func CamelCaseHelper(old string) string {
	now := ""
	tmp := strings.Split(old, "_")
	for _, v := range tmp {
		if len(v) == 0 {
			continue
		}
		c := string(v[0])
		if v[0] >= 'a' && v[0] <= 'z' {
			c = string(v[0] - 32)
		}
		now += c
		now += v[1:]
	}
	return now
}

// UnderScoreCaseHelper go_style
func UnderScoreCaseHelper(old string) string {
	now := ""
	for _, v := range old {
		if v >= 65 && v <= 90 {
			v += 32
			now += "_"
		}
		now += string(v)
	}
	if now[0] == '_' {
		return now[1:]
	}
	return now
}

// LowerCaseHelper  gostyle
func LowerCaseHelper(old string) string {
	return strings.ToLower(CamelCaseHelper(old))
}
