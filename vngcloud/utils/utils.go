package utils

import "encoding/json"

func ToJSON(s interface{}) string {
	b, err := json.Marshal(s)
	if err != nil {
		return ""
	}
	return string(b)
}
