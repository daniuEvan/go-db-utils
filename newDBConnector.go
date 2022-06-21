/**
 * @date: 2022/6/20
 * @desc:
 */

package go_db_utils

import (
	"errors"
	"go-db-utils/dbDriver"
	"go-db-utils/types"
	"reflect"
	"strings"
)

//
// NewDBConnector
// @Description: 获取数据库连接
// @param dbOptions: 数据库配置信息
// @return dbLinker: 数据库连接
// @return err:
//
func NewDBConnector(dbOption types.DataBaseOption) (dbLinker types.DBLinker, err error) {
	dbType := strings.ToLower(dbOption.DBType)
	refEntry := reflect.ValueOf(dbDriver.Entry{})
	refMethod := refEntry.MethodByName(leftToUpper(dbType))
	args := []reflect.Value{
		reflect.ValueOf(dbOption),
	}
	valueSlice := refMethod.Call(args)
	// 判断错误
	errRes := valueSlice[1].Interface()
	if errRes != nil {
		return nil, errRes.(error)
	}
	var ok bool
	dbLinker, ok = valueSlice[0].Interface().(types.DBLinker)
	if !ok {
		return nil, errors.New("断言类型错误")
	}
	return dbLinker, err
}
