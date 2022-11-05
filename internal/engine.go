package internal

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/wslky/dbgo/internal/config"
	define2 "github.com/wslky/dbgo/internal/define"
	"github.com/wslky/dbgo/internal/utils"
	"log"
	"os"
	"strings"
)

type Engine struct {
	conf      config.Config
	tableDefs []*define2.TableDefine
	db        *sql.DB
}

func New(config config.Config) *Engine {
	return &Engine{
		tableDefs: make([]*define2.TableDefine, 0),
		conf:      config,
	}
}

func (e *Engine) OutGoFile() {
	log.Println("start connect db...")
	e.initDB()

	log.Println("start generate go model...")
	e.generateGoModel()

	log.Println("start output go file...")
	e.doOutGoFile()

	log.Println("over")
}

// doOutGoFile output go file
func (e *Engine) doOutGoFile() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	dir += "/"
	dir = strings.ReplaceAll(dir, "\\", "/")

	log.Println("current dir:", dir)

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
}

// generateGoModel generate go model
func (e *Engine) generateGoModel() {
	e.loadTablesFromDB()
}

// initDB init db
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

// loadTablesFromDB load table info from db
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

		table := &define2.TableDefine{
			TableName: tableName,
			Columns:   make([]*define2.ColumnDefine, 0),
		}

		e.loadColumnsFromDB(table)
	}
}

// loadColumnsFromDB load columns info from db
func (e *Engine) loadColumnsFromDB(table *define2.TableDefine) {
	res, err := e.db.Query(utils.GetSql4TableDef(e.conf.DBName, table.TableName))
	if err != nil {
		log.Fatal(err)
	}

	for res.Next() {
		var column define2.ColumnDefine
		err := res.Scan(&column.ColumnName, &column.DataType)
		if err != nil {
			log.Fatal(err)
		}

		table.Columns = append(table.Columns, &column)
	}

	e.tableDefs = append(e.tableDefs, table)
}

// formatName format name
func formatName(table *define2.TableDefine, conf config.Config) *nameData {
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

// generateGoFileData generate go file data
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

// getPlaceholder get placeholder for format
func getPlaceholder(length int, str string) string {
	s := ""
	for i := len(str); i < length; i++ {
		s += " "
	}
	return s
}

// packageLine generate package name line
func packageLine(name string) string {
	return fmt.Sprintf("package %v%v%v", name, _ENTER, _ENTER)
}

// defLine generate struct define line
func defLine(name string) string {
	return fmt.Sprintf("%v %v %v {%v", _TYPE, name, _STRUCT, _ENTER)
}

// attrLine generate struct attribute line
func attrLine(name1, name2, dataType, p1, p2 string) string {
	return fmt.Sprintf("%v %v%v %v%v %v%v", _TAB, name1, p1, dataType, p2, jsonTag(name2), _ENTER)
}

// jsonTag generate struct json tag
func jsonTag(name string) string {
	return fmt.Sprintf("`json:\"%v\"`", name)
}

// endLine generate struct end line
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

// nameData table name data
type nameData struct {
	packageName string

	structName string

	jsonName []string

	fileAttrName []string

	datatype []string

	// maxNameLen max name length for format
	maxNameLen int

	// maxTypeLen is the max length of data type for format
	maxTypeLen int
}
