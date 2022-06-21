/**
 * @date: 2022/6/21
 * @desc:
 */

package swapType

import (
	"errors"
	"fmt"
	"strings"
)

//
// GetGoTypeWithSqlType
// @Description: 通过数据库的数据类型转为go内置数据类型
// @param dbType: 数据库类型(mysql/pg/oracle/vertica)
// @param sqlType: 数据库的数据类型
// @return string: go内置数据类型
// @return error:
//
func GetGoTypeWithSqlType(dbType, sqlType string, value []uint8) (goTypeValue interface{}, err error) {
	dbType = strings.ToLower(dbType)
	sqlType = strings.ToLower(sqlType)
	var goTypeStr string
	switch dbType {
	// todo pg/oracle
	case "mysql":
		goTypeStr = mysql[sqlType]
	default:
		return nil, errors.New(fmt.Sprintf("数据库类型不匹配:%s", dbType))
	}
	goTypeValue, err = swapType(value, goTypeStr)
	if err != nil {
		return nil, err
	}
	return goTypeValue, err
}
