/**
 * @date: 2022/6/20
 * @desc:
 */

package dbDriver

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"go-db-utils/dbDriver/common/swapType"
	"go-db-utils/types"
)

//
// MysqlLink
// @Description: Mysql 连接
//
type MysqlLink struct {
	DBOption types.DataBaseOption
	Conn     *sql.DB
}

func NewMysqlLink(dbOption types.DataBaseOption) (*MysqlLink, error) {
	// "用户名:密码@[连接方式](主机名:端口号)/数据库名"
	dataSourceName := fmt.Sprintf("%s:%s@(%s:%d)/%s", dbOption.DBUsername, dbOption.DBPassword, dbOption.DBHost, dbOption.DBPort, dbOption.DBName)
	if dbOption.Options != "" {
		dataSourceName += fmt.Sprintf("?%s", dbOption.Options)
	}
	conn, err := sql.Open(
		dbOption.DBType,
		dataSourceName,
	)
	if err != nil {
		return nil, err
	}

	// 数据库连接池设置
	conn.SetConnMaxLifetime(dbOption.ConnPollConfig.ConnMaxLifetime) // 连接存活时间
	conn.SetMaxOpenConns(dbOption.ConnPollConfig.MaxOpenConns)       // 最大连接数
	conn.SetMaxIdleConns(dbOption.ConnPollConfig.MaxIdleConns)       // 最大空闲连接数

	err = conn.Ping()
	if err != nil {
		return nil, err
	}
	return &MysqlLink{
		DBOption: dbOption,
		Conn:     conn,
	}, err
}

//
// GetConn
// @Description: 获取conn
//
func (l *MysqlLink) GetConn() *sql.DB {
	return l.Conn
}

//
// Query
// @Description: 原生查询,对结果不做任何转换
// @return map[string][]byte: value类型是[]byte([]uint8)
//
func (l *MysqlLink) Query(sqlStr string) ([]map[string][]byte, error) {
	rows, err := l.Conn.Query(sqlStr)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}
	values := make([][]byte, len(columns))        //values是每个列的值，这里获取到byte里
	scanArgs := make([]interface{}, len(columns)) //query.Scan的参数，因为每次查询出来的列是不定长的，用len(cols)定住当次查询的长度
	for i := range values {
		scanArgs[i] = &values[i]
	}
	results := make([]map[string][]byte, 0) // res
	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			return nil, err
		}
		row := make(map[string][]byte) //每行数据
		for k, v := range values {
			key := columns[k]
			row[key] = v
		}
		results = append(results, row)
	}
	return results, nil

}

//
// QueryWithString
// @Description: 字符串查询, 返回结果全部转为string
// @return map[string]string: value类型是string
//
func (l *MysqlLink) QueryWithString(sqlStr string) ([]map[string]string, error) {
	rows, err := l.Conn.Query(sqlStr)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}
	values := make([][]byte, len(columns))        //values是每个列的值，这里获取到byte里
	scanArgs := make([]interface{}, len(columns)) //query.Scan的参数，因为每次查询出来的列是不定长的，用len(cols)定住当次查询的长度
	for i := range values {
		scanArgs[i] = &values[i]
	}
	results := make([]map[string]string, 0) // res
	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			return nil, err
		}
		row := make(map[string]string) //每行数据
		for k, v := range values {
			key := columns[k]
			row[key] = string(v)
		}
		results = append(results, row)
	}
	return results, nil
}

//
// QueryWithRealType
// @Description: 真实字段类型查询, 返回结果全部转为interface, 通过断言转为对应类型
// @return map[string]string: value类型是interface, 通过断言转为对应类型
//
func (l *MysqlLink) QueryWithRealType(sqlStr string) ([]map[string]interface{}, error) {
	rows, err := l.Conn.Query(sqlStr)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}
	columnsType, err := rows.ColumnTypes()
	if err != nil {
		return nil, err
	}

	values := make([][]byte, len(columns))        //values是每个列的值，这里获取到byte里
	scanArgs := make([]interface{}, len(columns)) //query.Scan的参数，因为每次查询出来的列是不定长的，用len(cols)定住当次查询的长度
	for i := range values {
		scanArgs[i] = &values[i]
	}
	results := make([]map[string]interface{}, 0) // res
	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			return nil, err
		}
		row := make(map[string]interface{}) //每行数据
		for k, v := range values {
			key := columns[k]
			valueType := columnsType[k]
			realTypeValue, err := swapType.GetGoTypeWithSqlType(l.DBOption.DBType, valueType.DatabaseTypeName(), v)
			if err != nil {
				return nil, err
			}
			row[key] = realTypeValue
		}
		results = append(results, row)
	}
	return results, nil
}

//
// Exec
// @Description: 简单执行sql
// @param sqlStr:
// @return sql.Result:  result.RowsAffected()
// @return error:
//
func (l *MysqlLink) Exec(sqlStr string) (sql.Result, error) {
	result, err := l.Conn.Exec(sqlStr)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Close 关闭数据库连接
func (l *MysqlLink) Close() error {
	err := l.Conn.Close()
	if err != nil {
		return err
	}
	return nil
}
