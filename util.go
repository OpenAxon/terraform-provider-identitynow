package main

import (
	"fmt"
	"strings"
)

func toArrayInterface(in []string) []interface{} {
	out := make([]interface{}, len(in))
	for i, v := range in {
		out[i] = v
	}
	return out
}

func toArrayString(in []interface{}) []string {
	out := make([]string, len(in))
	for i, v := range in {
		if v == nil {
			out[i] = ""
			continue
		}
		out[i] = v.(string)
	}
	return out
}

func splitAccountSchemaAttributeID(id string) (sourceId string, name string, err error) {
	separator := "-"

	result := strings.Split(id, separator)
	if len(result) == 2 {
		return result[0], result[1], nil
	}
	return "", "", fmt.Errorf("[ERROR Getting source id and name. id: %s", id)
}
