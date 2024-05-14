package main

import (
	"fmt"
	"reflect"
	"sort"
	"strings"
)

const SEPARATOR = " "

func Serialize(v any) string {
	typeOf := reflect.TypeOf(v)

	if typeOf.Kind() != reflect.Struct {
		return ""
	}

	var params []string
	for i := 0; i < typeOf.NumField(); i++ {
		reflectType := typeOf.Field(i)
		tag := reflectType.Tag.Get("param")
		if tag == "-" || tag == "" {
			continue
		}

		valueOf := reflect.ValueOf(v)
		reflectValue := valueOf.Field(i)
		params = append(params, fmt.Sprintf("%s=%v", tag, reflectValue.Interface()))
	}

	sort.Strings(params)

	return strings.Join(params, SEPARATOR)
}
