/**
 * @date: 2022/6/20
 * @desc:
 */

package go_db_utils

import (
	"go-db-utils/types"
	"time"
)

const (
	ConnMaxLifetime = time.Minute * 5
	MaxOpenConns    = 20
	MaxIdleConns    = 20
)

//
// NewDBOption
// @Description: 创建数据库连接对象
// @return types.DataBaseOption:
//
func NewDBOption(dbType, dbHost string, dbPort uint, dbName, dbUsername, dbPassword, options string) types.DataBaseOption {
	dbOption := types.DataBaseOption{
		DBType:     dbType,
		DBHost:     dbHost,
		DBPort:     dbPort,
		DBName:     dbName,
		DBUsername: dbUsername,
		DBPassword: dbPassword,
		Options:    options, // 格式 arg1=1&arg2=2&arg3=3
	}

	dbOption.ConnPollConfig.ConnMaxLifetime = ConnMaxLifetime
	dbOption.ConnPollConfig.MaxOpenConns = MaxOpenConns
	dbOption.ConnPollConfig.MaxIdleConns = MaxIdleConns
	return dbOption
}
