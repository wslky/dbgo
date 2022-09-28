package dbgo

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/wslky/dbgo/config"
	"github.com/wslky/dbgo/define"
	"github.com/wslky/dbgo/utils"
	"log"
	"os"
	"strings"
)

type Engine struct {
	conf      config.Config
	tableDefs []*define.TableDefine
	db        *sql.DB
}

func New(config config.Config) *Engine {
	return &Engine{
		tableDefs: make([]*define.TableDefine, 0),
		conf:      config,
	}
}

func (e *Engine) OutGoFile() {
	e.initDB()
	e.generateGoModel()
	outGoFile(e)
}

func outGoFile(e *Engine) {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	dir += "/"
	dir += e.conf.Path
	dir = strings.ReplaceAll(dir, "\\", "/")
	var data string
	for _, v := range e.tableDefs {
		file, err := os.Create(fmt.Sprintf("%v%v%v", dir, e.conf.FileNameType(v.TableName), ".go"))
		if err != nil {
			log.Fatal(err)
		}
		tmp := formatName(v, e.conf)
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

func (e *Engine) generateGoModel() {
	e.loadTablesFromDB()
}

func (e *Engine) generateDBModel() {

}

func (e *Engine) initDB() {
	db, err := sql.Open("mysql", e.conf.DBUser+":"+e.conf.DBPassword+"@tcp("+e.conf.DBUrl+")/"+e.conf.DBName)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	e.db = db
}

func (e *Engine) loadTablesFromDB() {
	res, err := e.db.Query("SHOW TABLES")
	if err != nil {
		log.Fatal(err)
	}

	var tableName string
	for res.Next() {
		err := res.Scan(&tableName)
		if err != nil {
			log.Fatal(err)
		}
		table := &define.TableDefine{
			TableName: tableName,
			Columns:   make([]*define.ColumnDefine, 0),
		}
		e.loadColumnsFromDB(table)
	}
}

func (e *Engine) loadColumnsFromDB(table *define.TableDefine) {
	res, err := e.db.Query(utils.GetSql4TableDef(e.conf.DBName, table.TableName))
	if err != nil {
		log.Fatal(err)
	}

	for res.Next() {
		var column define.ColumnDefine
		err := res.Scan(&column.ColumnName, &column.DataType)
		if err != nil {
			log.Fatal(err)
		}
		table.Columns = append(table.Columns, &column)
	}

	e.tableDefs = append(e.tableDefs, table)
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
		res += _TIME
	}
	res += defLine(data.structName)
	for i := 0; i < len(data.jsonName); i++ {
		res += attrLine(data.fileAttrName[i], data.jsonName[i],
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
	return fmt.Sprintf("package %v%v%v", name, _ENTER, _ENTER)
}

func defLine(name string) string {
	return fmt.Sprintf("%v %v %v {%v", _TYPE, name, _STRUCT, _ENTER)
}

func attrLine(name1, name2, dataType, p1, p2 string) string {
	return fmt.Sprintf("%v %v%v %v%v %v%v", _TAB, name1, p1, dataType, p2, jsonTag(name2), _ENTER)
}

func jsonTag(name string) string {
	return fmt.Sprintf("`json:\"%v\"`", name)
}

func endLine() string {
	return "}"
}

var (
	_TYPE   = "type"
	_STRUCT = "struct"
	_TAB    = "\t"
	_ENTER  = "\n"
	_TIME   = "import (\n" + "\t\"time\"" + "\n)\n\n"
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
