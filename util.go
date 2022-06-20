/**
 * @date: 2022/6/20
 * @desc:
 */

package go_db_utils

import "strings"

//
// leftToUpper
// @Description: 字符串首字母转大写
// @param str:
// @return string:
//
func leftToUpper(str string) string {
	if len(str) > 0 {
		return strings.ToUpper(string(str[0])) + str[1:]
	}
	return str
}
