/**
 * @date: 2022/6/20
 * @desc:
 */

package dbDriver

import "go-db-utils/types"

type Entry struct {
}

func (e Entry) Mysql(option types.DataBaseOption) (dbLink types.DBLinker, err error) {
	dbLink, err = NewMysqlLink(option)
	if err != nil {
		return nil, err
	}
	return dbLink, err
}

func (e Entry) Postgres(option types.DataBaseOption) (int, int) {
	return 1, 2
}
