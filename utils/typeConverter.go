package utils

import "encoding/json"

func TypeConverter[R any](data any) (R, error) {
	var result R
	b, err := json.Marshal(&data)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(b, &result)
	if err != nil {
		return result, err
	}
	return result, err

}
