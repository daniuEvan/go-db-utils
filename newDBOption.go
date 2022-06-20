/**
 * @date: 2022/6/20
 * @desc:
 */

package go_db_utils

import "go-db-utils/types"

//
// NewDBOption
// @Description: 创建数据库连接对象
// @return types.DataBaseOption:
//
func NewDBOption() types.DataBaseOption {
	return types.DataBaseOption{
		DBType:     "",
		DBHost:     "",
		DBPort:     0,
		DBName:     "",
		DBUsername: "",
		DBPassword: "",
		Options:    "",
	}
}
