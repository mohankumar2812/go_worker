package utils

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func GetParams(input interface{}, name string) string {
	inp := convertOBJtoString(input)
	val := ""

	if name == "attributes" {
		val = "atr"
	} else if name == "traits" {
		val = "uatr"
	}

	output := make(map[string]map[string]string)
	for i := 1; i <= len(inp)/3; i++ {
		key := inp[fmt.Sprintf("%sk%d", val, i)]
		if key == "" {
			break
		}
		value := inp[fmt.Sprintf("%sv%d", val, i)]
		attrType := inp[fmt.Sprintf("%st%d", val, i)]

		if _, ok := output[key]; !ok {
			output[key] = make(map[string]string)
		}

		if _, ok := output[name][key]; !ok {
			output[key] = make(map[string]string)
		}

		output[key]["value"] = value
		output[key]["type"] = attrType
	}

	jsonData, err := json.MarshalIndent(output, "", "   ")
	if err != nil {
		return ""
	}

	return string(jsonData)
}

func convertOBJtoString(obj interface{}) map[string]string {
	values := make(map[string]string)

	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)

		if field.PkgPath != "" {
			continue
		}

		strValue := fmt.Sprintf("%v", value.Interface())
		tag := field.Tag.Get("json")
		if tag != "" && tag != "-" {
			values[tag] = strValue
		} else {
			values[field.Name] = strValue
		}
	}

	return values
}
