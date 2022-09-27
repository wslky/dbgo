package dbgo

import (
	"fmt"
	"github.com/wslky/dbgo/config"
	"github.com/wslky/dbgo/define"
	"github.com/wslky/dbgo/utils"
	"log"
	"os"
	"strings"
)

type nameData struct {
	packageName  string
	structName   string
	jsonName     []string
	fileAttrName []string
	datatype     []string
	maxNameLen   int
	maxTypeLen   int
}

func OutGoFile(e *Engine) {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	dir += "/"
	dir += e.Conf.Path
	dir = strings.ReplaceAll(dir, "\\", "/")
	var data string
	for _, v := range e.TableDefs {
		file, err := os.Create(fmt.Sprintf("%v%v%v", dir, e.Conf.FileNameType(v.TableName), ".go"))
		if err != nil {
			log.Fatal(err)
		}
		tmp := formatName(v, e.Conf)
		data = generateGoFileData(tmp)
		_, err = file.WriteString(data)
		if err != nil {
			log.Fatal(err)
		}
		err = file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}
	log.Println("完成")
}

func formatName(table *define.TableDefine, conf config.Config) *nameData {
	res := &nameData{
		jsonName:     make([]string, len(table.Columns)),
		fileAttrName: make([]string, len(table.Columns)),
		datatype:     make([]string, len(table.Columns)),
		packageName:  conf.PackageName,
		structName:   utils.CamelCaseHelper(table.TableName),
		maxNameLen:   0,
		maxTypeLen:   0,
	}
	for i, v := range table.Columns {
		res.jsonName[i] = conf.JSONType(v.ColumnName)
		res.fileAttrName[i] = conf.FilePropertyType(v.ColumnName)
		res.datatype[i] = v.GOType()
		if len(res.datatype[i]) > res.maxTypeLen {
			res.maxTypeLen = len(res.datatype[i])
		}
		if len(res.fileAttrName[i]) > res.maxNameLen {
			res.maxNameLen = len(res.fileAttrName[i])
		}
	}
	return res
}

func generateGoFileData(data *nameData) string {
	res := packageLine(data.packageName)
	if data.maxTypeLen >= 9 {
		res += TIME
	}
	res += defLine(data.structName)
	for i := 0; i < len(data.jsonName); i++ {
		res += attrLine(data.fileAttrName[i],
			data.datatype[i], getPlaceholder(data.maxNameLen, data.fileAttrName[i]),
			getPlaceholder(data.maxTypeLen, data.datatype[i]))
	}
	res += endLine()
	return res
}

func getPlaceholder(length int, str string) string {
	s := ""
	for i := len(str); i < length; i++ {
		s += " "
	}
	return s
}

func packageLine(name string) string {
	return fmt.Sprintf("package %v%v%v", name, ENTER, ENTER)
}

func defLine(name string) string {
	return fmt.Sprintf("%v %v %v {%v", TYPE, name, STRUCT, ENTER)
}

func attrLine(name, dataType, p1, p2 string) string {
	return fmt.Sprintf("%v %v%v %v%v %v%v", TAB, name, p1, dataType, p2, jsonTag(name), ENTER)
}

func jsonTag(name string) string {
	return fmt.Sprintf("`json:\"%v\"`", name)
}

func endLine() string {
	return "}"
}

var (
	TYPE   = "type"
	STRUCT = "struct"
	TAB    = "\t"
	ENTER  = "\n"
	TIME   = "import (\n" + "\t\"time\"" + "\n)\n\n"
)
