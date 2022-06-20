/**
 * @date: 2022/6/20
 * @desc: interface
 */

package types

import "database/sql"

//
// DBLinker
// @Description: 数据库连接抽象层
//
type DBLinker interface {
	// GetConn 获取conn
	GetConn() *sql.DB
	//
	// Query
	// @Description: 原生查询,对结果不做任何转换
	// @return map[string][]byte: value类型是[]byte([]uint8)
	//
	Query(sqlStr string) ([]map[string][]byte, error)

	//
	// QueryWithString
	// @Description: 字符串查询, 返回结果全部转为string
	// @return map[string]string: value类型是string
	//
	QueryWithString(sqlStr string) ([]map[string]string, error)

	//
	// QueryWithRealType
	// @Description: 真实字段类型查询, 返回结果全部转为interface, 通过断言转为对应类型
	// @return map[string]string: value类型是interface, 通过断言转为对应类型
	//
	QueryWithRealType(sqlStr string) ([]map[string]interface{}, error)

	//
	// Exec
	// @Description: 执行update/delete/insert语句
	// @return sql.Result
	//
	Exec(sqlStr string) (sql.Result, error)
	// Close 关闭数据库连接
	Close() error
}
