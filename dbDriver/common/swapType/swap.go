/**
 * @date: 2022/6/21
 * @desc:
 */

package swapType

import (
	"encoding/binary"
	"errors"
	"fmt"
	"math"
	"strconv"
)

func swapType(value []uint8, goType string) (interface{}, error) {
	switch goType {
	case "string":
		return string(value), nil
	case "float64":
		value = toEight(value)
		bits := binary.LittleEndian.Uint64(value)
		resFloat64 := math.Float64frombits(bits)
		return resFloat64, nil
	case "int64", "int":
		value = toEight(value)
		resInt64 := int64(binary.BigEndian.Uint64(value))
		return resInt64, nil
	case "uint64", "uint":
		value = toEight(value)
		resUint64 := binary.BigEndian.Uint64(value)
		return resUint64, nil
	case "[]byte":
		return value, nil
	case "bool":
		s := string(value)
		boolRes, err := strconv.ParseBool(s)
		if err != nil {
			return nil, err
		}
		return boolRes, err
		//return bool(value), nil
	case "time.Time": // 日期转为字符串
		return string(value), nil
	default:
		return nil, errors.New(fmt.Sprintf("数据类型不匹配:%s", goType))
	}
}

// []int8 补齐八位
func toEight(value []uint8) []uint8 {
	if len(value) == 8 {
		return value
	}
	eightSlice := make([]uint8, 8)
	for k, v := range value {
		eightSlice[8-len(value)+k] = v
	}
	return eightSlice
}
