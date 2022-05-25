package error

import "strings"

func SplitErrorsToMap(msg string) (string, map[string]string) {
	// Remove last char ";"
	if msg[len(msg)-1:] == ";" {
		msg = msg[:len(msg)-1]
	}

	// Msg to map
	msgMap := strings.Split(msg, ";")

	// Splitting ...
	var jsonMap = make(map[string]string)
	for _, invalidValueStr := range msgMap {
		invalidValue := strings.Split(invalidValueStr, ":")

		// if not like json string
		if len(invalidValue) != 2 {
			return msg, nil
		}
		tag := strings.TrimSpace(invalidValue[0])
		msg := strings.TrimSpace(invalidValue[1])

		jsonMap[tag] = msg
	}

	return "", jsonMap
}
