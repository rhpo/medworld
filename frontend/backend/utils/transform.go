package utils

import (
	"encoding/json"
	"regexp"
	"strings"
)

var (
	matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
	matchAllCap   = regexp.MustCompile("([a-z0-9])([A-Z])")
)

// ToSnakeCase converts a string from camelCase to snake_case
func ToSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

// ToCamelCase converts a string from snake_case to camelCase
func ToCamelCase(str string) string {
	parts := strings.Split(str, "_")
	for i := 1; i < len(parts); i++ {
		if len(parts[i]) > 0 {
			parts[i] = strings.ToUpper(parts[i][:1]) + parts[i][1:]
		}
	}
	return strings.Join(parts, "")
}

// TransformKeys recursively transforms keys in a map
func TransformKeys(data interface{}, transformer func(string) string) interface{} {
	switch v := data.(type) {
	case map[string]interface{}:
		result := make(map[string]interface{})
		for key, value := range v {
			result[transformer(key)] = TransformKeys(value, transformer)
		}
		return result
	case []interface{}:
		result := make([]interface{}, len(v))
		for i, item := range v {
			result[i] = TransformKeys(item, transformer)
		}
		return result
	default:
		return data
	}
}

// SnakeToCamel converts JSON with snake_case keys to camelCase
func SnakeToCamel(data interface{}) interface{} {
	return TransformKeys(data, ToCamelCase)
}

// CamelToSnake converts JSON with camelCase keys to snake_case
func CamelToSnake(data interface{}) interface{} {
	return TransformKeys(data, ToSnakeCase)
}

// MapToStruct converts a map to a struct using JSON marshaling
func MapToStruct(m map[string]interface{}, s interface{}) error {
	data, err := json.Marshal(m)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, s)
}
