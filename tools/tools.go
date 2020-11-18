package tools

import "encoding/json"

func Response(code int, msg string, data interface{}) string {
	out := make(map[string]interface{})
	out["code"] = 1
	out["msg"] = "input pi error!"
	if data == nil {
		out["data"] = []byte{}
	} else {
		out["data"] = data
	}

	json, _ := json.Marshal(out)

	return string(json)
}
