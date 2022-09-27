package define

type ColumnDefine struct {
	ColumnName string
	DataType   string
	Size       int
}

func (c *ColumnDefine) GOType() string {
	switch c.DataType {
	case "bigint":
		return "int64"
	case "binary":
		return ""
	case "bit":
		return ""
	case "blob":
		return "string"
	case "char":
		return "byte"
	case "date":
		return "time.Time"
	case "datetime":
		return "time.Time"
	case "decimal":
		return "float64"
	case "double":
		return "float64"
	case "enum":
		return ""
	case "float":
		return "float32"
	case "geometry":
		return ""
	case "geometrycollection":
		return ""
	case "int":
		return "int"
	case "integer":
		return "int"
	case "json":
		return "string"
	case "linestring":
		return ""
	case "longblob":
		return "string"
	case "longtext":
		return "string"
	case "mediumblob":
		return "string"
	case "mediumint":
		return "int"
	case "mediumtext":
		return "string"
	case "multilinestring":
		return ""
	case "multipoint":
		return ""
	case "multipolygon":
		return ""
	case "numeric":
		return ""
	case "point":
		return ""
	case "polygon":
		return ""
	case "real":
		return ""
	case "set":
		return ""
	case "smallint":
		return "int"
	case "text":
		return "string"
	case "time":
		return "time.Time"
	case "timestamp":
		return "time.Time"
	case "tinyblob":
		return "string"
	case "tinyint":
		return "int"
	case "tinytext":
		return "string"
	case "varbinary":
		return "string"
	case "varchar":
		return "string"
	case "year":
		return "time.Time"
	default:
		return "string"
	}
}
func (c *ColumnDefine) DBType() string {
	switch c.DataType {
	case "string":
		return "varchar"
	case "int":
		return "int"
	case "int8":
		return "tinyint"
	case "int16":
		return "int"
	case "bool":
		return "tinyint"
	case "float32":
		return "float32"
	case "float64":
		return "float64"
	case "int32":
		return "int"
	case "int64":
		return "bigint"
	case "uint":
		return "int"
	case "uint8":
		return "tinyint"
	case "uint16":
		return "int"
	case "uint32":
		return "int"
	case "uint64":
		return "bigint"
	default:
		return "varchar"
	}
}
