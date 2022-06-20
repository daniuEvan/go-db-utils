/**
 * @date: 2022/6/20
 * @desc: struct
 */

package types

//
// DataBaseOption
// @Description: 数据库连接信息
//
type DataBaseOption struct {
	DBType     string // todo 字符串校验(mysql/oracle/pg/vertica)
	DBHost     string
	DBPort     uint
	DBName     string
	DBUsername string
	DBPassword string
	Options    string // 格式 arg1=1&arg2=2&arg3=3
}
