package main

import (
	"encoding/json"

	"gopkg.in/yaml.v3"
)

// pg_flags: STRICT IMMUTABLE
func MapToJsonText(yamlSrc string) string {
	if yamlSrc == "" {
		return ""
	}
	var m = make(map[string]interface{})
	yaml.Unmarshal([]byte(yamlSrc), &m)
	b, _ := json.Marshal(m)
	return string(b)
}


// pg_flags: STRICT IMMUTABLE
func SliceToJsonText(yamlSrc string) string {
	if yamlSrc == "" {
		return ""
	}
	var m []interface{}
	yaml.Unmarshal([]byte(yamlSrc), &m)
	b, _ := json.Marshal(m)
	return string(b)
}