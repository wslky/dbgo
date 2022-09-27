package dbgo

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/wslky/dbgo/config"
	"github.com/wslky/dbgo/define"
	"github.com/wslky/dbgo/utils"
	"log"
)

type Engine struct {
	Conf      config.Config
	TableDefs []*define.TableDefine
	db        *sql.DB
}

func New(config config.Config) *Engine {
	return &Engine{
		TableDefs: make([]*define.TableDefine, 0),
		Conf:      config,
	}
}

func (e *Engine) OutGoFile() {
	e.initDB()
	e.generateGoModel()
	OutGoFile(e)
}

func (e *Engine) generateGoModel() {
	e.loadTablesFromDB()
}

func (e *Engine) generateDBModel() {

}

func (e *Engine) initDB() {
	db, err := sql.Open("mysql", e.Conf.DBUser+":"+e.Conf.DBPassword+"@tcp("+e.Conf.DBUrl+")/"+e.Conf.DBName)
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
	res, err := e.db.Query(utils.GetSql4TableDef(e.Conf.DBName, table.TableName))
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

	e.TableDefs = append(e.TableDefs, table)
}
