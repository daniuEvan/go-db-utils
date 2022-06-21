/**
 * @date: 2022/6/20
 * @desc: struct
 */

package types

import "time"

//
// DataBaseOption
// @Description: 数据库连接信息
//
type DataBaseOption struct {
	DBType         string
	DBHost         string
	DBPort         uint
	DBName         string
	DBUsername     string
	DBPassword     string
	ConnPollConfig struct {
		ConnMaxLifetime time.Duration // 连接存活时间default 5 minutes
		MaxOpenConns    int           // 最大连接数default 20
		MaxIdleConns    int           // 最大空闲连接数default 20
	}
	Options string // 格式 arg1=1&arg2=2&arg3=3
}
