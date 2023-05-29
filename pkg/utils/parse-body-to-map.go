package utils

import "encoding/json"

func ParseBodyToMap(body []byte) (map[string]interface{}, error) {
	out := make(map[string]interface{})
	err := json.Unmarshal(body, &out)
	if err != nil {
		return nil, err
	}
	return out, err
}
