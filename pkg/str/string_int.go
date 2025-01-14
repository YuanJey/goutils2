package str

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type StringInt string
type IntString int64

func (is *StringInt) UnmarshalJSON(data []byte) error {
	// 试着查看是否是一个字符串
	var strValue string
	if err := json.Unmarshal(data, &strValue); err == nil {
		// 如果是字符串，尝试转换为 int64
		value, err := strconv.ParseInt(strValue, 10, 64)
		if err != nil {
			return err
		}
		*is = StringInt(strconv.FormatInt(value, 10))
		return nil
	}

	// 试着查看是否是一个整数
	var intValue int64
	if err := json.Unmarshal(data, &intValue); err == nil {
		*is = StringInt(strconv.FormatInt(intValue, 10))
		return nil
	}

	return fmt.Errorf("invalid value for Int64OrString: %s", data)
}
func (is *IntString) UnmarshalJSON(data []byte) error {
	var strValue string
	if err := json.Unmarshal(data, &strValue); err == nil {
		value, err := strconv.ParseInt(strValue, 10, 64)
		if err != nil {
			return err
		}
		*is = IntString((value))
		return nil
	}
	var Value int
	if err := json.Unmarshal(data, &Value); err == nil {
		*is = IntString(Value)
		return nil
	}

	return fmt.Errorf("invalid value for IntString: %s", data)
}
