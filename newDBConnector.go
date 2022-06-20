/**
 * @date: 2022/6/20
 * @desc:
 */

package go_db_utils

import "go-db-utils/types"

//
// NewDBConnector
// @Description: 获取数据库连接
// @param dbOptions: 数据库配置信息
// @return dbLinker: 数据库连接
// @return err:
//
func NewDBConnector(dbOption types.DataBaseOption) (dbLinker *types.DBLinker, err error) {
	dbType := dbOption.DBType
	dbHost := dbOption.DBHost
	dbPort := dbOption.DBPort
	dbName := dbOption.DBName
	dbUsername := dbOption.DBUsername
	dbPassword := dbOption.DBPassword
	options := dbOption.Options

	return nil, err
}
